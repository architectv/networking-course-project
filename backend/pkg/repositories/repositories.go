package repositories

import (
	"yak/backend/pkg/models"
	"yak/backend/pkg/repositories/postgres"

	"github.com/jmoiron/sqlx"
)

type User interface {
	GetAll() ([]*models.User, error)
	Create(user *models.User) (int, error)
	Get(nickname, password string) (*models.User, error)
	GetByNickname(nickname string) (*models.User, error)
	// SignOut(token string) error
	// FindToken(token string) error
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

// func NewRepository(db *mongo.Database) *Repository {
// 	return &Repository{
// 		User:     repoMongo.NewUserMongo(db),
// 		Project:  repoMongo.NewProjectMongo(db),
// 		Board:    repoMongo.NewBoardMongo(db),
// 		TaskList: repoMongo.NewTaskListMongo(db),
// 		Task:     repoMongo.NewTaskMongo(db),
// 	}
// }

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: postgres.NewUserPg(db),
	}
}
