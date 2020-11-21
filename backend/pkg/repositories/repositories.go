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
	Create(project *models.Project) (int, error)
	GetAll(userId int) ([]*models.Project, error)
	GetById(projectId int) (*models.Project, error)
	Delete(projectId int) error
	Update(projectId int, project *models.UpdateProject) error
	GetPermissions(userId, projectId int) (*models.Permission, error)
}

type Board interface {
	// GetAll(userId, projectId string) ([]models.Board, error)
	Create(userId int, board *models.Board) (int, error)
	GetById(boardId int) (*models.Board, error)
	GetPermissions(userId, boardId int) (*models.Permission, error)
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
		User:    postgres.NewUserPg(db),
		Project: postgres.NewProjectPg(db),
		Board:   postgres.NewBoardPg(db),
	}
}
