package services

import (
	"yak/backend/pkg/models"
	"yak/backend/pkg/repositories"
)

type UrlValidatorService struct {
	projectRepo repositories.Project
	boardRepo   repositories.Board
	listRepo    repositories.TaskList
	taskRepo    repositories.Task
}

func NewUrlValidator(boardRepo repositories.Board, listRepo repositories.TaskList,
	taskRepo repositories.Task) *UrlValidatorService {
	return &UrlValidatorService{
		boardRepo: boardRepo,
		listRepo:  listRepo,
		taskRepo:  taskRepo,
	}
}

func (s *UrlValidatorService) Validation(urlIds *models.UrlIds) *models.ApiResponse {
	r := &models.ApiResponse{}

	if urlIds.ProjectId != 0 && urlIds.BoardId != 0 {
		var projectId int
		board, err := s.boardRepo.GetById(urlIds.BoardId)
		if err == nil {
			projectId = board.ProjectId
		}

		if urlIds.ProjectId != projectId {
			r.Error(StatusBadRequest, "There is no requested board inside the project")
			return r
		}
	}

	if urlIds.ListId != 0 {
		var boardId int
		list, err := s.listRepo.GetById(urlIds.ListId)
		if err == nil {
			boardId = list.BoardId
		}

		if urlIds.BoardId != boardId {
			r.Error(StatusBadRequest, "There is no requested list inside the board")
			return r
		}
	}

	if urlIds.TaskId != 0 {
		var listId int
		task, err := s.taskRepo.GetById(urlIds.TaskId)
		if err == nil {
			listId = task.ListId
		}

		if urlIds.ListId != listId {
			r.Error(StatusBadRequest, "There is no requested task inside the list")
			return r
		}
	}

	r.Set(StatusOK, "OK", Map{})
	return r
}
