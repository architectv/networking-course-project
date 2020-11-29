package postgres

import (
	"errors"
	"fmt"
	"yak/backend/pkg/models"

	"github.com/jmoiron/sqlx"
)

const (
	DbResultNotFound = "sql: no rows in result set"
	PROJECT          = 1
	BOARD            = 2
)

type ProjectPermsPg struct {
	db *sqlx.DB
}

func NewProjectPermsPg(db *sqlx.DB) *ProjectPermsPg {
	return &ProjectPermsPg{db: db}
}

func (r *ProjectPermsPg) Get(objectId, userId, objectType int) (*models.Permission, error) {
	var objectIdTitle, objectTable string
	switch objectType {
	case PROJECT:
		objectIdTitle = "project_id"
		objectTable = projectUsersTable
	case BOARD:
		objectIdTitle = "board_id"
		objectTable = boardUsersTable
	default:
		return nil, errors.New("Object type is not defined")
	}

	permissions := &models.Permission{}
	query := fmt.Sprintf(
		`SELECT per.read, per.write, per.admin 
		FROM %s AS pu
			INNER JOIN %s AS per ON pu.permissions_id = per.id
		WHERE pu.%s = $1 AND pu.user_id = $2`,
		objectTable, permissionsTable, objectIdTitle)

	row := r.db.QueryRow(query, objectId, userId)
	err := row.Scan(&permissions.Read, &permissions.Write, &permissions.Admin)
	fmt.Println(objectId, userId, permissions)
	return permissions, err
}

func (r *ProjectPermsPg) Create(projectId, memberId int, permissions *models.Permission) (int, error) {

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	_, err = r.Get(projectId, memberId, PROJECT)
	if err != nil && err.Error() != DbResultNotFound {
		return 0, err
	} else if err == nil {
		return 0, errors.New("Member already has permissions in the project")
	}

	var projectPermsId int
	permissionsId, err := createPermissions(tx, permissions)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	query := fmt.Sprintf(
		`INSERT INTO %s (user_id, project_id, permissions_id)
		VALUES ($1, $2, $3) RETURNING id`, projectUsersTable)

	row := tx.QueryRow(query, memberId, projectId, permissionsId)
	if err := row.Scan(&projectPermsId); err != nil {
		tx.Rollback()
		return 0, err
	}
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()
	return projectPermsId, err
}

func (r *ProjectPermsPg) Delete(projectId, memberId int) error {
	query := fmt.Sprintf(
		`DELETE FROM %s AS per USING %s AS pu
		WHERE per.id = pu.permissions_id AND pu.project_id=$1 AND pu.user_id=$2`,
		permissionsTable, projectUsersTable)
	_, err := r.db.Exec(query, projectId, memberId)
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
