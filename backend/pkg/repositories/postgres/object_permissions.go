package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"yak/backend/pkg/models"

	"github.com/jmoiron/sqlx"
)

const (
	DbResultNotFound = "sql: no rows in result set"
	IsProject        = 1
	IsBoard          = 2
)

type ObjectPermsPg struct {
	db *sqlx.DB
}

type ObjectParams struct {
	Title   string
	IdTitle string
	Table   string
}

func NewObjectPermsPg(db *sqlx.DB) *ObjectPermsPg {
	return &ObjectPermsPg{db: db}
}

func (r *ObjectPermsPg) Get(objectId, memberId, objectType int) (*models.Permission, error) {
	permissions := &models.Permission{}
	objParams, err := getObjectParams(objectType)
	if err != nil {
		return permissions, err
	}

	query := fmt.Sprintf(
		`SELECT per.read, per.write, per.admin 
		FROM %s AS obj
			INNER JOIN %s AS per ON obj.permissions_id = per.id
		WHERE obj.%s = $1 AND obj.user_id = $2`,
		objParams.Table, permissionsTable, objParams.IdTitle)

	row := r.db.QueryRow(query, objectId, memberId)
	err = row.Scan(&permissions.Read, &permissions.Write, &permissions.Admin)
	fmt.Println(objectId, memberId, permissions)
	return permissions, err
}

func (r *ObjectPermsPg) Create(objectId, memberId, objectType int, permissions *models.Permission) (int, error) {
	fmt.Println(objectId, memberId, objectType)
	objParams, err := getObjectParams(objectType)
	if err != nil {
		return 0, err
	}

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	_, err = r.Get(objectId, memberId, objectType)
	if err != nil && err.Error() != DbResultNotFound {
		return 0, err
	} else if err == nil {
		errText := fmt.Sprintf("Member already has permissions in the %s", objParams.Title)
		return 0, errors.New(errText)
	}

	var objectPermsId int
	permissionsId, err := createPermissions(tx, permissions)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	query := fmt.Sprintf(
		`INSERT INTO %s (user_id, %s, permissions_id)
		VALUES ($1, $2, $3) RETURNING id`, objParams.Table, objParams.IdTitle)

	row := tx.QueryRow(query, memberId, objectId, permissionsId)
	if err := row.Scan(&objectPermsId); err != nil {
		tx.Rollback()
		return 0, err
	}
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()
	return objectPermsId, err
}

func (r *ObjectPermsPg) Delete(objectId, memberId, ownerProjectId, objectType int) error {
	objParams, err := getObjectParams(objectType)
	if err != nil {
		return err
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	if ownerProjectId != 0 {
		if objectType == IsProject {
			err = deleteMemberFromAllBoardsInProject(tx, objectId, memberId)
			if err != nil {
				tx.Rollback()
				return err
			}

			err := updateOwnerIdByProjectId(tx, objectId, memberId, ownerProjectId)
			if err != nil {
				tx.Rollback()
				return err
			}
		} else if objectType == IsBoard {
			err := updateOwnerIdByBoardId(tx, objectId, memberId, ownerProjectId)
			if err != nil {
				tx.Rollback()
				return err
			}
		} else {
			return errors.New("Object type is not defined")
		}
	}

	query := fmt.Sprintf(
		`DELETE FROM %s AS per USING %s AS obj
		WHERE per.id = obj.permissions_id AND obj.%s=$1 AND obj.user_id=$2`,
		permissionsTable, objParams.Table, objParams.IdTitle)
	_, err = tx.Exec(query, objectId, memberId)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return err
}

func (r *ObjectPermsPg) Update(objectId, memberId, ownerProjectId, objectType int, permissions *models.UpdatePermission) error {
	objParams, err := getObjectParams(objectType)
	if err != nil {
		return err
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	if ownerProjectId != 0 {
		if objectType == IsProject {
			err := updateOwnerIdByProjectId(tx, objectId, memberId, ownerProjectId)
			if err != nil {
				tx.Rollback()
				return err
			}
		} else if objectType == IsBoard {
			err := updateOwnerIdByBoardId(tx, objectId, memberId, ownerProjectId)
			if err != nil {
				tx.Rollback()
				return err
			}
		} else {
			return errors.New("Object type is not defined")
		}
	}

	var objectPermsId int
	query := fmt.Sprintf(
		`SELECT obj.permissions_id 
		FROM %s AS obj
		WHERE obj.%s = $1 AND obj.user_id = $2`,
		objParams.Table, objParams.IdTitle)

	row := tx.QueryRow(query, objectId, memberId)
	err = row.Scan(&objectPermsId)
	fmt.Println(objectPermsId)

	if err = updatePermissions(tx, objectPermsId, permissions); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func getObjectParams(objectType int) (*ObjectParams, error) {
	var objParams ObjectParams
	switch objectType {
	case IsProject:
		objParams = ObjectParams{
			Title:   "project",
			IdTitle: "project_id",
			Table:   projectUsersTable,
		}
	case IsBoard:
		objParams = ObjectParams{
			Title:   "board",
			IdTitle: "board_id",
			Table:   boardUsersTable,
		}
	default:
		fmt.Println(objectType)
		return &objParams, errors.New("Object type is not defined")
	}
	return &objParams, nil
}

func deleteMemberFromAllBoardsInProject(tx *sql.Tx, projectId, memberId int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id IN (
			SELECT bu.permissions_id
			FROM %s AS bu
				INNER JOIN %s AS b ON bu.board_id = b.id
			WHERE b.project_id = $1 AND b.owner_id = $2 AND bu.user_id = b.owner_id)`,
		permissionsTable, boardUsersTable, boardsTable)
	fmt.Println(query, projectId, memberId)
	_, err := tx.Exec(query, projectId, memberId)
	return err
}