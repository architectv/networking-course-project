package services

import (
	"yak/backend/pkg/models"
	"yak/backend/pkg/repositories"
)

type User interface {
	GetAll() ([]models.User, error)
	Create(user models.User) (string, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (string, error)
}

type Project interface {
}

type Board interface {
	GetAll(userId, projectId string) ([]models.Board, error)
	Create(userId, projectId string, board models.Board) (string, error)
}

type TaskList interface {
}

type Task interface {
}

type Service struct {
	User
	Project
	Board
	TaskList
	Task
}

func NewService(repos *repositories.Repository) *Service {
	return &Service{
		User:     NewUserService(repos.User),
		Project:  NewProjectService(repos.Project),
		Board:    NewBoardService(repos.Board, repos.Project),
		TaskList: NewTaskListService(repos.TaskList),
		Task:     NewTaskService(repos.Task),
	}
}
