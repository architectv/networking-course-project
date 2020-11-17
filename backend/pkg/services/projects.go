package services

import (
	"fmt"
	"time"
	"errors"
	"yak/backend/pkg/models"
	"yak/backend/pkg/repositories"
)

type ProjectService struct {
	repo repositories.Project
}

func (s *ProjectService) Create(userId string, project models.Project) (string, error) {
	project.OwnerId = userId
	curTime := time.Now().Unix() 
	project.Datetimes = &models.Datetimes {
		Created:  curTime,
		Updated:  curTime,
		Accessed: curTime,
	}

	if project.DefaultPermissions == nil {
        project.DefaultPermissions = &models.Permission{
            Read:  true,
            Write: true,
            Grant: false,
        }
    }

	projectId, err := s.repo.Create(project)
	if err != nil {
		return "", err
	}

	ownerPermission := &models.Permission{
		Read:  true,
		Write: true,
		Grant: true,
	}
	projectUser := models.ProjectUser {
		UserId: userId,
		ProjectId: projectId,
		Permissions: ownerPermission,
	}

	if s.repo.ProjectUserCreate(projectUser) != nil {
		return "", err
	}

	return projectId, nil
}


func NewProjectService(repo repositories.Project) *ProjectService {
	return &ProjectService{repo: repo}
}

func (s *ProjectService) GetAll(userId string) ([]models.Project, error) {
	projectsId, err := s.repo.ProjectIdGetByPermissions(userId)
	fmt.Println(projectsId)
	if err != nil {
		return make([]models.Project, 0), err
	}
	return s.repo.GetAll(projectsId)
}

func (s *ProjectService) GetById(userId, projectId string) (models.Project, error) {
	permissions, err := s.repo.GetPermission(userId, projectId)

	var project models.Project
	if err != nil {
		return project, err
	} else if permissions == nil || permissions.Read == false {  // TODO
		return project, errors.New("Forbidden")
	} 
	return s.repo.GetById(projectId)
}

func (s *ProjectService) Update(userId, projectId string, project models.Project) error {
	permissions, err := s.repo.GetPermission(userId, projectId)

	if err != nil {
		return err
	} else if permissions == nil || permissions.Write == false {  // TODO
		return errors.New("Forbidden")
	} 
	
	project, err = s.repo.GetById(userId)
	if err != nil {
		return err
	}

	curTime := time.Now().Unix() 
	project.Datetimes = &models.Datetimes {
		Created:  project.Datetime.Created,
		Updated:  curTime,
		Accessed: curTime,
	}

	if err := s.repo.Update(projectId, project); err != nil {
		return err
	}

	return nil
}