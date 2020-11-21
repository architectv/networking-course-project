package services

import (
	"time"
	"yak/backend/pkg/models"
	"yak/backend/pkg/repositories"
)

type BoardService struct {
	repo        repositories.Board
	projectRepo repositories.Project
}

func NewBoardService(repo repositories.Board, projectRepo repositories.Project) *BoardService {
	return &BoardService{repo: repo, projectRepo: projectRepo}
}

// func (s *BoardService) GetAll(userId, projectId string) *models.ApiResponse {
// 	r := &models.ApiResponse{}
// 	projects, err := s.repo.GetAll(userId, projectId)
// 	if err != nil {
// 		r.Error(StatusInternalServerError, err.Error())
// 		return r
// 	}
// 	r.Set(StatusOK, "OK", Map{"projects": projects})
// 	return r
// }

func (s *BoardService) Create(userId, projectId int, board *models.Board) *models.ApiResponse {
	r := &models.ApiResponse{}
	permissions, err := s.projectRepo.GetPermissions(userId, projectId)
	if err != nil || permissions.Read == false {
		r.Error(StatusForbidden, "Forbidden")
		return r
	}

	board.ProjectId = projectId
	curTime := time.Now().Unix()
	datetimes := &models.Datetimes{
		Created:  curTime,
		Updated:  curTime,
		Accessed: curTime,
	}
	board.Datetimes = datetimes

	if board.DefaultPermissions == nil {
		board.DefaultPermissions = &models.Permission{
			Read:  true,
			Write: true,
			Admin: false,
		}
	}

	// TODO: добавлять админа проекта?
	boardId, err := s.repo.Create(userId, board)
	if err != nil {
		r.Error(StatusInternalServerError, err.Error())
		return r
	}

	r.Set(StatusOK, "OK", Map{"boardId": boardId})
	return r
}

func (s *BoardService) GetById(userId, projectId, boardId int) *models.ApiResponse {
	r := &models.ApiResponse{}
	// TODO: права для проекта?
	// projectPermissions, err := s.projectRepo.GetPermissions(userId, projectId)
	// if err != nil || projectPermissions.Read == false {
	// 	r.Error(StatusForbidden, "Forbidden")
	// 	return r
	// }

	boardPermissions, err := s.repo.GetPermissions(userId, boardId)
	if err != nil || boardPermissions.Read == false {
		r.Error(StatusForbidden, "Forbidden")
		return r
	}

	board, err := s.repo.GetById(boardId)
	if err != nil {
		r.Error(StatusInternalServerError, err.Error())
		return r
	}

	r.Set(StatusOK, "OK", Map{"board": board})
	return r
}
