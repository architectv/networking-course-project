package repositories

import (
	"yak/backend/pkg/models"
	repoMongo "yak/backend/pkg/repositories/mongo"

	"go.mongodb.org/mongo-driver/mongo"
)

type User interface {
	GetAll() ([]models.User, error)
}

type Project interface {
	Create(project models.Project) (string, error)
	GetAll(userId string) ([]models.Project, error)
	// GetById(userId, projectId string) (models.Project, error)
	// Delete(userId, projectId string) error
	// Update(userId, projectId string, project models.Project) error
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
