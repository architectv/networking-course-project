package services

import (
	"yak/backend/pkg/models"
	"yak/backend/pkg/repositories"
)

type User interface {
	GetAll() ([]*models.User, error)
	Create(user *models.User) *models.ApiResponse
	GenerateToken(username, password string) *models.ApiResponse
	ParseToken(token string) (int, error)
	// SignOut(token string) *models.ApiResponse
}

type Project interface {
	Create(userId int, project *models.Project) *models.ApiResponse
	GetAll(userId int) *models.ApiResponse
	GetById(userId, projectId int) *models.ApiResponse
	// Delete(userId, projectId int)  *models.ApiResponse
	// Update(userId, projectId int, project models.Project)  *models.ApiResponse
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
