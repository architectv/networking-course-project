package postgres

import (
	"fmt"
	"yak/backend/pkg/models"

	"github.com/jmoiron/sqlx"
)

type ProjectPg struct {
	db *sqlx.DB
}

func NewProjectPg(db *sqlx.DB) *ProjectPg {
	return &ProjectPg{db: db}
}

func (r *ProjectPg) Create(project *models.Project) (int, error) {
	var projectId int
	var defPermissionId, datetimesId, permissionId int

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	defPermissions := project.DefaultPermissions
	query := fmt.Sprintf(
		`INSERT INTO %s (read, write, admin)
		VALUES ($1, $2, $3) RETURNING id`, permissionsTable)

	row := tx.QueryRow(query, defPermissions.Read, defPermissions.Write,
		defPermissions.Admin)
	if err := row.Scan(&defPermissionId); err != nil {
		tx.Rollback()
		return 0, err
	}

	datetimes := project.Datetimes
	query = fmt.Sprintf(
		`INSERT INTO %s (created, updated, accessed)
		VALUES ($1, $2, $3) RETURNING id`, datetimesTable)

	row = tx.QueryRow(query, datetimes.Created, datetimes.Updated,
		datetimes.Accessed)
	if err := row.Scan(&datetimesId); err != nil {
		tx.Rollback()
		return 0, err
	}

	query = fmt.Sprintf(
		`INSERT INTO %s
		(owner_id, default_permissions_id, datetimes_id, title, description)
		VALUES ($1, $2, $3, $4, $5) RETURNING id`, projectsTable)

	row = tx.QueryRow(query, project.OwnerId, defPermissionId,
		datetimesId, project.Title, project.Description)
	if err := row.Scan(&projectId); err != nil {
		tx.Rollback()
		return 0, err
	}

	query = fmt.Sprintf(
		`INSERT INTO %s (read, write, admin)
		VALUES (true, true, true) RETURNING id`, permissionsTable)

	row = tx.QueryRow(query)
	if err := row.Scan(&permissionId); err != nil {
		tx.Rollback()
		return 0, err
	}

	query = fmt.Sprintf(
		`INSERT INTO %s (user_id, project_id, permissions_id)
		VALUES ($1, $2, $3)`, projectUsersTable)

	_, err = tx.Exec(query, project.OwnerId, projectId, permissionId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()
	return projectId, nil
}

func (r *ProjectPg) GetById(projectId int) (*models.Project, error) {
	project := &models.Project{}
	defaultPermissions := &models.Permission{}
	datetimes := &models.Datetimes{}

	query := fmt.Sprintf(
		`SELECT p.id, p.owner_id, dper.read, dper.write, dper.admin, 
		d.created, d.updated, d.accessed, p.title, p.description
		FROM %s AS p
			INNER JOIN %s AS dper ON p.default_permissions_id = dper.id
			INNER JOIN %s AS d ON p.datetimes_id = d.id
		WHERE p.id = $1`,
		projectsTable, permissionsTable, datetimesTable)

	row := r.db.QueryRow(query, projectId)
	err := row.Scan(&project.Id, &project.OwnerId, &defaultPermissions.Read,
		&defaultPermissions.Write, &defaultPermissions.Admin,
		&datetimes.Created, &datetimes.Updated, &datetimes.Accessed,
		&project.Title, &project.Description)

	if err != nil {
		return nil, err
	}
	project.DefaultPermissions = defaultPermissions
	project.Datetimes = datetimes
	return project, nil
}

func (r *ProjectPg) GetAll(userId int) ([]*models.Project, error) {
	var projects []*models.Project

	query := fmt.Sprintf(
		`SELECT p.id, p.owner_id, dper.read, dper.write, dper.admin, 
		d.created, d.updated, d.accessed, p.title, p.description
		FROM %s AS pu
			INNER JOIN %s AS per ON pu.permissions_id = per.id
			INNER JOIN %s AS p ON pu.project_id = p.id
			INNER JOIN %s AS dper ON p.default_permissions_id = dper.id
			INNER JOIN %s AS d ON p.datetimes_id = d.id
		WHERE pu.user_id = $1 AND per.read = true`,
		projectUsersTable, permissionsTable, projectsTable, permissionsTable,
		datetimesTable)

	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		project := &models.Project{}
		defaultPermissions := &models.Permission{}
		datetimes := &models.Datetimes{}

		err := rows.Scan(&project.Id, &project.OwnerId, &defaultPermissions.Read,
			&defaultPermissions.Write, &defaultPermissions.Admin,
			&datetimes.Created, &datetimes.Updated, &datetimes.Accessed,
			&project.Title, &project.Description)

		if err != nil {
			return nil, err
		}

		project.DefaultPermissions = defaultPermissions
		project.Datetimes = datetimes
		projects = append(projects, project)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}

// func (r *ProjectPg) Update(projectId string, project models.Project) error {

// 	return nil
// }

func (r *ProjectPg) GetPermissions(userId, projectId int) (*models.Permission, error) {
	permissions := &models.Permission{}

	query := fmt.Sprintf(
		`SELECT per.read, per.write, per.admin
		FROM %s AS pu
			INNER JOIN %s AS per ON pu.permissions_id = per.id
		WHERE pu.project_id = $1 AND pu.user_id = $2`,
		projectUsersTable, permissionsTable)

	row := r.db.QueryRow(query, projectId, userId)
	err := row.Scan(&permissions.Read, &permissions.Write, &permissions.Admin)

	if err != nil {
		return nil, err
	}

	return permissions, nil
}
