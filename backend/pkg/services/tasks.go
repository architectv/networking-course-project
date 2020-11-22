package services

import (
	"yak/backend/pkg/models"
	"yak/backend/pkg/repositories"
)

type TaskService struct {
	repo        repositories.Task
	boardRepo   repositories.Board
	projectRepo repositories.Project
}

func NewTaskService(repo repositories.Task, boardRepo repositories.Board, projectRepo repositories.Project) *TaskService {
	return &TaskService{repo: repo, boardRepo: boardRepo, projectRepo: projectRepo}
}

func (s *TaskService) GetAll(userId, projectId, boardId, listId int) *models.ApiResponse {
	r := &models.ApiResponse{}
	projectPermissions, err := s.projectRepo.GetPermissions(userId, projectId)
	if err != nil || projectPermissions.Read == false {
		r.Error(StatusForbidden, "Forbidden")
		return r
	}

	boardPermissions, err := s.boardRepo.GetPermissions(userId, boardId)
	if err != nil || boardPermissions.Read == false {
		r.Error(StatusForbidden, "Forbidden")
		return r
	}

	tasks, err := s.repo.GetAll(listId)
	if err != nil {
		r.Error(StatusInternalServerError, err.Error())
		return r
	}

	r.Set(StatusOK, "OK", Map{"tasks": tasks})
	return r
}

func (s *TaskService) GetById(userId, projectId, boardId, listId, taksId int) *models.ApiResponse {
	r := &models.ApiResponse{}
	projectPermissions, err := s.projectRepo.GetPermissions(userId, projectId)
	if err != nil || projectPermissions.Read == false {
		r.Error(StatusForbidden, "Forbidden")
		return r
	}

	boardPermissions, err := s.boardRepo.GetPermissions(userId, boardId)
	if err != nil || boardPermissions.Read == false {
		r.Error(StatusForbidden, "Forbidden")
		return r
	}

	task, err := s.repo.GetById(taksId)
	if err != nil {
		r.Error(StatusInternalServerError, err.Error())
		return r
	}

	r.Set(StatusOK, "OK", Map{"task": task})
	return r
}
