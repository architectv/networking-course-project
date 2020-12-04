package postgres

import (
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

type ProjectPermsPg struct {
	db *sqlx.DB
}

type ObjectParams struct {
	Title   string
	IdTitle string
	Table   string
}

func NewProjectPermsPg(db *sqlx.DB) *ProjectPermsPg {
	return &ProjectPermsPg{db: db}
}

func (r *ProjectPermsPg) Get(objectId, memberId, objectType int) (*models.Permission, error) {
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

func (r *ProjectPermsPg) Create(objectId, memberId, objectType int, permissions *models.Permission) (int, error) {
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

func (r *ProjectPermsPg) Delete(objectId, memberId, ownerProjectId, objectType int) error {
	objParams, err := getObjectParams(objectType)
	if err != nil {
		return err
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	if objectType == IsBoard && ownerProjectId != 0 {
		err := updateOwnerId(tx, objectId, memberId, ownerProjectId)
		if err != nil {
			tx.Rollback()
			return err
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

func (r *ProjectPermsPg) Update(projectId, memberId int, permissions *models.UpdatePermission) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	var projectPermsId int
	query := fmt.Sprintf(
		`SELECT pu.permissions_id 
		FROM %s AS pu
		WHERE pu.project_id = $1 AND pu.user_id = $2`,
		projectUsersTable)

	row := tx.QueryRow(query, projectId, memberId)
	err = row.Scan(&projectPermsId)

	if err = updatePermissions(tx, projectPermsId, permissions); err != nil {
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
