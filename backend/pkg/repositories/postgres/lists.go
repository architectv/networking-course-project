package postgres

import (
	"fmt"
	"strings"
	"yak/backend/pkg/models"

	"github.com/jmoiron/sqlx"
)

type TaskListPg struct {
	db *sqlx.DB
}

func NewTaskListPg(db *sqlx.DB) *TaskListPg {
	return &TaskListPg{db: db}
}

func (r *TaskListPg) GetAll(boardId int) ([]*models.TaskList, error) {
	var lists []*models.TaskList
	// TODO: order by pos
	query := fmt.Sprintf(
		`SELECT tl.id, tl.board_id, tl.title, tl.position
		FROM %s AS tl
			INNER JOIN %s AS b ON tl.board_id = b.id
		WHERE b.id = $1`,
		taskListsTable, boardsTable)
	if err := r.db.Select(&lists, query, boardId); err != nil {
		return nil, err
	}

	return lists, nil
}

func (r *TaskListPg) GetById(listId int) (*models.TaskList, error) {
	list := &models.TaskList{}
	query := fmt.Sprintf(
		`SELECT * FROM %s WHERE id = $1`, taskListsTable)
	err := r.db.Get(list, query, listId)

	return list, err
}

func (r *TaskListPg) Create(list *models.TaskList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	var position int

	query := fmt.Sprintf(
		`SELECT MAX(tl.position)
		FROM %s AS tl
		INNER JOIN %s AS b ON b.id = tl.board_id
		WHERE b.id = $1;`, taskListsTable, boardsTable)

	row := r.db.QueryRow(query, list.BoardId)
	if err := row.Scan(&position); err != nil {
		return 0, err
	}
	position++

	query = fmt.Sprintf(
		`INSERT INTO %s (board_id, title, position)
		VALUES ($1, $2, $3) RETURNING id`, taskListsTable)

	row = r.db.QueryRow(query, list.BoardId, list.Title, position)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	tx.Commit()

	return id, nil
}

func (r *TaskListPg) Delete(listId int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	var boardId, position int
	query := fmt.Sprintf(`SELECT board_id, position FROM %s WHERE id = $1`, taskListsTable)
	row := tx.QueryRow(query, listId)
	err = row.Scan(&boardId, &position)
	if err != nil {
		tx.Rollback()
		return err
	}

	query = fmt.Sprintf(
		`UPDATE %s SET position = position - 1
	WHERE board_id = $1 AND position > $2`,
		taskListsTable)
	_, err = tx.Exec(query, boardId, position)
	if err != nil {
		tx.Rollback()
		return err
	}

	query = fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, taskListsTable)
	_, err = tx.Exec(query, listId)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return err
}

func (r *TaskListPg) Update(listId int, input *models.UpdateTaskList) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Position != nil {
		newPos := *input.Position

		var boardId, oldPos int
		query := fmt.Sprintf(`SELECT board_id, position FROM %s WHERE id = $1`, taskListsTable)
		row := tx.QueryRow(query, listId)
		err = row.Scan(&boardId, &oldPos)
		if err != nil {
			tx.Rollback()
			return err
		}

		var operation string
		var start, end int
		if oldPos < newPos {
			operation = "-"
			start, end = oldPos+1, newPos
		} else if oldPos > newPos {
			operation = "+"
			start, end = newPos, oldPos-1
		}

		if operation != "" {
			setValues = append(setValues, fmt.Sprintf("position=$%d", argId))
			args = append(args, newPos)
			argId++

			query = fmt.Sprintf(
				`UPDATE %s SET position = position %s 1
			WHERE board_id = $1 AND position >= $2 AND position <= $3`,
				taskListsTable, operation)
			_, err = tx.Exec(query, boardId, start, end)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf(`UPDATE %s SET %s where id=$%d`,
		taskListsTable, setQuery, argId)
	args = append(args, listId)
	_, err = tx.Exec(query, args...)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return err
}
