package services

import (
	"yak/backend/pkg/models"
	"yak/backend/pkg/repositories"
)

type User interface {
	GetAll() ([]models.User, error)
}

type Project interface {
	Create(userId string, project models.Project) (string, error)
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
		Board:    NewBoardService(repos.Board),
		TaskList: NewTaskListService(repos.TaskList),
		Task:     NewTaskService(repos.Task),
	}
}
