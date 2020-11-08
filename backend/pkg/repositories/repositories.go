package repositories

import (
	"yak/backend/pkg/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type User interface {
	GetAll() ([]models.User, error)
}

type Project interface {
}

type Board interface {
}

type TaskList interface {
}

type Task interface {
}

type Repository struct {
	User
	Project
	Board
	TaskList
	Task
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		User:     NewUserMongo(db),
		Project:  NewProjectMongo(db),
		Board:    NewBoardMongo(db),
		TaskList: NewTaskListMongo(db),
		Task:     NewTaskMongo(db),
	}
}
