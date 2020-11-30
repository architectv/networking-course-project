package services

import (
	"errors"
	"yak/backend/pkg/models"
	"yak/backend/pkg/repositories"
)

const (
	IsProject            = 1
	IsBoard              = 2
	ErrPermsIsNotDefined = "Permissions is not defined"
	ErrPermsIncor        = "Permissions are set incorrectly"
)

type ProjectPermsService struct {
	repo        repositories.ProjectPerms
	projectRepo repositories.Project
}

func NewProjectPermsService(repo repositories.ProjectPerms, projectRepo repositories.Project) *ProjectPermsService {
	return &ProjectPermsService{repo: repo, projectRepo: projectRepo}
}

func (s *ProjectPermsService) Get(userId, projectId, memberId int) *models.ApiResponse {
	r := &models.ApiResponse{}

	_, err := s.repo.Get(projectId, userId, IsProject)
	if err != nil {
		if err.Error() == DbResultNotFound {
			r.Error(StatusNotFound, "Request author is not project member")
			return r
		}
		r.Error(StatusInternalServerError, err.Error())
		return r
	}

	permissions, err := s.repo.Get(projectId, memberId, IsProject)
	if err != nil {
		if err.Error() == DbResultNotFound {
			r.Error(StatusNotFound, "Project member not found")
			return r
		}
		r.Error(StatusInternalServerError, err.Error())
		return r
	}

	r.Set(StatusOK, "OK", Map{"permissions": permissions})
	return r
}

func (s *ProjectPermsService) Create(userId, projectId, memberId int, projectPerms *models.Permission) *models.ApiResponse {
	r := &models.ApiResponse{}

	if err := permsValidation(projectPerms); err != nil {
		if err.Error() == ErrPermsIsNotDefined {
			project, err := s.projectRepo.GetById(projectId)
			if err != nil {
				r.Error(StatusInternalServerError, err.Error())
				return r
			}
			if project.DefaultPermissions == nil { // TODO права по умолчанию должны обязательно указываться при создании проекта или доски
				r.Error(StatusInternalServerError, "Default permissions is not defined")
				return r
			}
			projectPerms = project.DefaultPermissions
		} else {
			r.Error(StatusBadRequest, err.Error())
			return r
		}
	}

	permissions, err := s.repo.Get(projectId, userId, IsProject)
	if err != nil {
		if err.Error() == DbResultNotFound {
			r.Error(StatusNotFound, "Request author is not project member")
			return r
		}
		r.Error(StatusInternalServerError, err.Error())
		return r
	}

	if permissions.Admin != true {
		r.Error(StatusForbidden, "Request author is not project admin")
		return r
	}

	permissionsId, err := s.repo.Create(projectId, memberId, IsProject, projectPerms)
	if err != nil {
		r.Error(StatusInternalServerError, err.Error())
		return r
	}

	r.Set(StatusOK, "OK", Map{"Project permissions id": permissionsId})
	return r
}

func (s *ProjectPermsService) Delete(userId, projectId, memberId int) *models.ApiResponse {
	r := &models.ApiResponse{}

	permissions, err := s.repo.Get(projectId, userId, IsProject)
	if err != nil {
		if err.Error() == DbResultNotFound {
			r.Error(StatusNotFound, "Request author is not project member")
			return r
		}
		r.Error(StatusInternalServerError, err.Error())
		return r
	}

	if permissions.Admin != true {
		r.Error(StatusForbidden, "Request author is not project admin")
		return r
	}

	_, err = s.repo.Get(projectId, memberId, IsProject)
	if err != nil {
		if err.Error() == DbResultNotFound {
			r.Error(StatusNotFound, "Excluding user is not project member")
			return r
		}
		r.Error(StatusInternalServerError, err.Error())
		return r
	}

	project, err := s.projectRepo.GetById(projectId)
	if err != nil {
		r.Error(StatusInternalServerError, err.Error())
		return r
	}
	if project.OwnerId == memberId {
		r.Error(StatusBadRequest, "You can't exclude project owner")
		return r
	}

	err = s.repo.Delete(projectId, memberId, 0, IsProject)
	if err != nil {
		r.Error(StatusInternalServerError, err.Error())
		return r
	}

	r.Set(StatusOK, "OK", Map{})
	return r
}

func (s *ProjectPermsService) Update(userId, projectId, memberId int, projectPerms *models.UpdatePermission) *models.ApiResponse {
	r := &models.ApiResponse{}

	if err := updatePermsValidation(projectPerms); err != nil {
		r.Error(StatusBadRequest, err.Error())
		return r
	}

	permissions, err := s.repo.Get(projectId, userId, IsProject)
	if err != nil {
		if err.Error() == DbResultNotFound {
			r.Error(StatusNotFound, "Request author is not project member")
			return r
		}
		r.Error(StatusInternalServerError, err.Error())
		return r
	}

	if permissions.Admin != true {
		r.Error(StatusForbidden, "Request author is not project admin")
		return r
	}

	_, err = s.repo.Get(projectId, memberId, IsProject)
	if err != nil {
		if err.Error() == DbResultNotFound {
			r.Error(StatusNotFound, "Updating user is not project member")
			return r
		}
		r.Error(StatusInternalServerError, err.Error())
		return r
	}

	project, err := s.projectRepo.GetById(projectId) // TODO изменение прав автора
	if err != nil {
		r.Error(StatusInternalServerError, err.Error())
		return r
	}
	if project.OwnerId == memberId {
		r.Error(StatusBadRequest, "You can't update project owner permissions")
		return r
	}

	err = s.repo.Update(projectId, memberId, projectPerms)
	if err != nil {
		r.Error(StatusInternalServerError, err.Error())
		return r
	}

	r.Set(StatusOK, "OK", Map{})
	return r
}

func updatePermsValidation(updPerms *models.UpdatePermission) error {
	if updPerms != nil {
		switch {
		case updPerms.Read != nil && updPerms.Write != nil && updPerms.Admin != nil:
			perms := &models.Permission{
				Read:  *updPerms.Read,
				Write: *updPerms.Write,
				Admin: *updPerms.Admin,
			}
			return permsValidation(perms)
		case updPerms.Read != nil && updPerms.Write != nil:
			perms := &models.Permission{
				Read:  *updPerms.Read,
				Write: *updPerms.Write,
				Admin: false,
			}
			return permsValidation(perms)
		case updPerms.Read != nil:
			perms := &models.Permission{
				Read:  *updPerms.Read,
				Write: false,
				Admin: false,
			}
			return permsValidation(perms)
		case updPerms.Read == nil && updPerms.Write == nil && updPerms.Admin == nil:
			return errors.New(ErrPermsIsNotDefined)
		default:
			return errors.New(ErrPermsIncor)
		}
	} else {
		return errors.New(ErrPermsIsNotDefined)
	}
}

func permsValidation(perms *models.Permission) error {
	if perms != nil {
		if perms.Read == true && perms.Write == false && perms.Admin == false ||
			perms.Read == true && perms.Write == true && perms.Admin == false ||
			perms.Read == true && perms.Write == true && perms.Admin == true {
			return nil
		} else if perms.Read == false && perms.Write == false && perms.Admin == false {
			return errors.New(ErrPermsIsNotDefined)
		} else {
			return errors.New(ErrPermsIncor)
		}
	} else {
		return errors.New(ErrPermsIncor)
	}
}
