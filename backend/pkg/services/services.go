package services

import (
	"context"
	"yak/backend/pkg/models"
	"yak/backend/pkg/repositories"
)

type User interface {
	GetAll() ([]models.User, error)
	Create(ctx context.Context, user *models.User) *models.ApiResponse
	GenerateToken(ctx context.Context, username, password string) *models.ApiResponse
	ParseToken(ctx context.Context, token string) (string, error)
	SignOut(ctx context.Context, token string) *models.ApiResponse
}

type Project interface {
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
