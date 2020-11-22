package postgres

import (
	"fmt"
	"yak/backend/pkg/models"

	"github.com/jmoiron/sqlx"
)

type TaskPg struct {
	db *sqlx.DB
}

func NewTaskPg(db *sqlx.DB) *TaskPg {
	return &TaskPg{db: db}
}

func (r *TaskPg) GetAll(listId int) ([]*models.Task, error) {
	var tasks []*models.Task
	// TODO: order by pos
	query := fmt.Sprintf(
		`SELECT t.id, t.list_id, t.title, d.created, d.updated, d.accessed, t.position
		FROM %s AS t
			INNER JOIN %s AS tl ON t.list_id = tl.id
			INNER JOIN %s AS d ON t.datetimes_id = d.id
		WHERE tl.id = $1`,

		tasksTable, taskListsTable, datetimesTable)

	rows, err := r.db.Query(query, listId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		task := &models.Task{}
		datetimes := &models.Datetimes{}

		err := rows.Scan(&task.Id, &task.ListId, &task.Title, &datetimes.Created,
			&datetimes.Updated, &datetimes.Accessed, &task.Position)

		if err != nil {
			return nil, err
		}

		task.Datetimes = datetimes
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *TaskPg) GetById(taskId int) (*models.Task, error) {
	task := &models.Task{}
	datetimes := &models.Datetimes{}

	query := fmt.Sprintf(
		`SELECT t.id, t.list_id, t.title, d.created, d.updated, d.accessed, t.position
		FROM %s AS t
		INNER JOIN %s AS d ON t.datetimes_id = d.id
		WHERE t.id = $1`,
		tasksTable, datetimesTable)

	row := r.db.QueryRow(query, taskId)
	err := row.Scan(&task.Id, &task.ListId, &task.Title, &datetimes.Created,
		&datetimes.Updated, &datetimes.Accessed, &task.Position)
	if err != nil {
		return nil, err
	}

	task.Datetimes = datetimes
	return task, nil
}
