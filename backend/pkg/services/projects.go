package services

import (
	"time"
	"yak/backend/pkg/models"
	"yak/backend/pkg/repositories"
)

type ProjectService struct {
	repo repositories.Project
}

func NewProjectService(repo repositories.Project) *ProjectService {
	return &ProjectService{repo: repo}
}

func (s *ProjectService) GetAll(userId string) ([]models.Project, error) {
	return s.repo.GetAll(userId)
}

func (s *ProjectService) Create(userId string, project models.Project) (string, error) {
	project.OwnerId = userId
	curTime := time.Now().Unix() 
	datetimes := models.Datetimes {curTime, curTime, curTime}
	project.Datetimes = &datetimes
	return s.repo.Create(project)
}