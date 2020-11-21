package postgres

import (
	"fmt"
	"yak/backend/pkg/models"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type BoardPg struct {
	db *sqlx.DB
}

func NewBoardPg(db *sqlx.DB) *BoardPg {
	return &BoardPg{db: db}
}

func (r *BoardPg) Create(userId int, board *models.Board) (int, error) {
	logrus.Info(board.Title, board.Datetimes, board.DefaultPermissions)
	var boardId int
	var defPermissionId, datetimesId, permissionId int

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	logrus.Info("check 1")

	defPermissions := board.DefaultPermissions
	query := fmt.Sprintf(
		`INSERT INTO %s (read, write, admin)
		VALUES ($1, $2, $3) RETURNING id`, permissionsTable)

	row := tx.QueryRow(query, defPermissions.Read, defPermissions.Write,
		defPermissions.Admin)
	if err := row.Scan(&defPermissionId); err != nil {
		tx.Rollback()
		return 0, err
	}

	logrus.Info("check 2")

	datetimes := board.Datetimes
	query = fmt.Sprintf(
		`INSERT INTO %s (created, updated, accessed)
		VALUES ($1, $2, $3) RETURNING id`, datetimesTable)

	row = tx.QueryRow(query, datetimes.Created, datetimes.Updated,
		datetimes.Accessed)
	if err := row.Scan(&datetimesId); err != nil {
		tx.Rollback()
		return 0, err
	}

	logrus.Info("check 3")

	query = fmt.Sprintf(
		`INSERT INTO %s
		(project_id, default_permissions_id, datetimes_id, title)
		VALUES ($1, $2, $3, $4) RETURNING id`, boardsTable)

	row = tx.QueryRow(query, board.ProjectId, defPermissionId,
		datetimesId, board.Title)
	if err := row.Scan(&boardId); err != nil {
		tx.Rollback()
		return 0, err
	}

	logrus.Info("check 4")

	query = fmt.Sprintf(
		`INSERT INTO %s (read, write, admin)
		VALUES (true, true, true) RETURNING id`, permissionsTable)

	row = tx.QueryRow(query)
	if err := row.Scan(&permissionId); err != nil {
		tx.Rollback()
		return 0, err
	}

	// TODO: права владельца проекта
	query = fmt.Sprintf(
		`INSERT INTO %s (user_id, board_id, permissions_id)
		VALUES ($1, $2, $3)`, boardUsersTable)

	_, err = tx.Exec(query, userId, boardId, permissionId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()
	return boardId, nil
}

func (r *BoardPg) GetById(boardId int) (*models.Board, error) {
	board := &models.Board{}
	defaultPermissions := &models.Permission{}
	datetimes := &models.Datetimes{}

	query := fmt.Sprintf(
		`SELECT b.id, b.project_id, bper.read, bper.write, bper.admin, 
		d.created, d.updated, d.accessed, b.title
		FROM %s AS b
			INNER JOIN %s AS bper ON b.default_permissions_id = bper.id
			INNER JOIN %s AS d ON b.datetimes_id = d.id
		WHERE b.id = $1`,
		boardsTable, permissionsTable, datetimesTable)

	row := r.db.QueryRow(query, boardId)
	err := row.Scan(&board.Id, &board.ProjectId, &defaultPermissions.Read,
		&defaultPermissions.Write, &defaultPermissions.Admin,
		&datetimes.Created, &datetimes.Updated, &datetimes.Accessed,
		&board.Title)
	if err != nil {
		return nil, err
	}

	board.DefaultPermissions = defaultPermissions
	board.Datetimes = datetimes
	return board, nil
}

func (r *BoardPg) GetPermissions(userId, boardId int) (*models.Permission, error) {
	permissions := &models.Permission{}

	query := fmt.Sprintf(
		`SELECT per.read, per.write, per.admin
		FROM %s AS bu
			INNER JOIN %s AS per ON bu.permissions_id = per.id
		WHERE bu.board_id = $1 AND bu.user_id = $2`,
		boardUsersTable, permissionsTable)

	row := r.db.QueryRow(query, boardId, userId)
	err := row.Scan(&permissions.Read, &permissions.Write, &permissions.Admin)
	if err != nil {
		return nil, err
	}

	return permissions, nil
}
