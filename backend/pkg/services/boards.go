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

func (s *BoardService) GetAll(userId, projectId string) ([]models.Board, error) {
	return s.repo.GetAll(userId, projectId)
}

func (s *BoardService) Create(userId, projectId string, board models.Board) (string, error) {
	// _, err := s.projectRepo.GetById(userId, projectId)
	// if err != nil {
	// 	return "", err
	// }

	board.ProjectId = projectId
	curTime := time.Now().Unix()
	datetimes := models.Datetimes{
		Created:  curTime,
		Updated:  curTime,
		Accessed: curTime,
	}
	board.Datetimes = &datetimes

	if board.DefaultPermissions == nil {
		board.DefaultPermissions = &models.Permission{
			Read:  true,
			Write: true,
			Grant: false,
		}
	}

	// TODO: boardUser

	return s.repo.Create(userId, projectId, board)
}
