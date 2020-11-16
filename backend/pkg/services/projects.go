package services

import (
	"fmt"
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
	projectsId, error := s.repo.ProjectIdGetByPermissions(userId)
	fmt.Println(projectsId)
	if error != nil {
		return make([]models.Project, 0), error
	}
	return s.repo.GetAll(projectsId)
}

func (s *ProjectService) Create(userId string, project models.Project) (string, error) {
	project.OwnerId = userId
	curTime := time.Now().Unix() 
	datetimes := models.Datetimes {curTime, curTime, curTime}
	project.Datetimes = &datetimes

	projectId, error := s.repo.Create(project)
	if error != nil {
		return "", error
	}

	projectUser := models.ProjectUser {
		UserId: userId,
		ProjectId: projectId,
		Permissions: project.DefaultPermissions,
	}

	if s.repo.ProjectUserCreate(projectUser) != nil {
		return "", error
	}

	return projectId, nil
}