package repositories

import (
	"yak/backend/pkg/models"
	repoMongo "yak/backend/pkg/repositories/mongo"

	"go.mongodb.org/mongo-driver/mongo"
)

type User interface {
	GetAll() ([]models.User, error)
	Create(user models.User) (string, error)
	GetUser(username, password string) (models.User, error)
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
		User:     repoMongo.NewUserMongo(db),
		Project:  repoMongo.NewProjectMongo(db),
		Board:    repoMongo.NewBoardMongo(db),
		TaskList: repoMongo.NewTaskListMongo(db),
		Task:     repoMongo.NewTaskMongo(db),
	}
}
