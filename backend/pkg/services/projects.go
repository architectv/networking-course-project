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

func (s *ProjectService) Create(userId int, project models.Project) *models.ApiResponse {
	r := &models.ApiResponse{}

	project.OwnerId = userId
	curTime := time.Now().Unix()
	project.Datetimes = &models.Datetimes{
		Created:  curTime,
		Updated:  curTime,
		Accessed: curTime,
	}

	if project.DefaultPermissions == nil {
		project.DefaultPermissions = &models.Permission{
			Read:  true,
			Write: true,
			Admin: false,
		}
	}

	projectId, err := s.repo.Create(project)
	if err != nil {
		r.Error(StatusInternalServerError, err.Error())
		return r
	}

	r.Set(StatusOK, "OK", Map{"projectId": projectId})
	return r
}

func (s *ProjectService) GetAll(userId int) *models.ApiResponse {
	r := &models.ApiResponse{}
	projects, err := s.repo.GetAll(userId)
	if err != nil {
		r.Error(StatusInternalServerError, err.Error())
		return r
	}
	r.Set(StatusOK, "OK", Map{"projects": projects})
	return r
}

func (s *ProjectService) GetById(userId, projectId int) *models.ApiResponse {
	r := &models.ApiResponse{}
	permissions, err := s.repo.GetPermissions(userId, projectId)

	if err != nil || permissions.Read == false {
		r.Error(StatusForbidden, "Forbidden")
		return r
	}

	project, err := s.repo.GetById(projectId)
	if err != nil {
		r.Error(StatusInternalServerError, err.Error())
		return r
	}
	r.Set(StatusOK, "OK", Map{"project": project})
	return r
}

// func (s *ProjectService) Update(userId, projectId string, project models.Project) error {
// 	permissions, err := s.repo.GetPermission(userId, projectId)
// 	if err != nil {
// 		return errors.New("Forbidden")
// 	}
// 	if permissions.Write == false {
// 		return errors.New("Forbidden")
// 	}

// 	project, err = s.repo.GetById(userId)
// 	if err != nil {
// 		return err
// 	}

// 	curTime := time.Now().Unix()
// 	project.Datetimes = &models.Datetimes{
// 		Created:  project.Datetimes.Created,
// 		Updated:  curTime,
// 		Accessed: curTime,
// 	}

// 	if err := s.repo.Update(projectId, project); err != nil {
// 		return err
// 	}

// 	return nil
// }
