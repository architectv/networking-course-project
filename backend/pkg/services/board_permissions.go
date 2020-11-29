package services

import (
	"yak/backend/pkg/models"
	"yak/backend/pkg/repositories"
)

type BoardPermsService struct {
	repo      repositories.ProjectPerms
	boardRepo repositories.Board
}

func NewBoardPermsService(repo repositories.ProjectPerms, boardRepo repositories.Board) *BoardPermsService {
	return &BoardPermsService{repo: repo, boardRepo: boardRepo}
}

func (s *BoardPermsService) Get(userId, projectId, boardId, memberId int) *models.ApiResponse {
	r := &models.ApiResponse{}

	_, err := s.repo.Get(boardId, userId, BOARD)
	if err != nil {
		if err.Error() == DbResultNotFound {
			r.Error(StatusNotFound, "Request author is not board member")
			return r
		}
		r.Error(StatusInternalServerError, err.Error())
		return r
	}

	permissions, err := s.repo.Get(boardId, memberId, BOARD)
	if err != nil {
		if err.Error() == DbResultNotFound {
			r.Error(StatusNotFound, "Member permissions not found")
			return r
		}
		r.Error(StatusInternalServerError, err.Error())
		return r
	}

	r.Set(StatusOK, "OK", Map{"permissions": permissions})
	return r
}

// func (s *ProjectPermsService) Create(userId, projectId, memberId int, projectPerms *models.Permission) *models.ApiResponse {
// 	r := &models.ApiResponse{}

// 	permissions, err := s.repo.Get(projectId, userId)
// 	if err != nil {
// 		if err.Error() == DbResultNotFound {
// 			r.Error(StatusNotFound, "Request author is not project member")
// 			return r
// 		}
// 		r.Error(StatusInternalServerError, err.Error())
// 		return r
// 	}

// 	if permissions.Admin != true {
// 		r.Error(StatusForbidden, "Request author is not project admin")
// 		return r
// 	}

// 	permissionsId, err := s.repo.Create(projectId, memberId, projectPerms)
// 	if err != nil {
// 		r.Error(StatusInternalServerError, err.Error())
// 		return r
// 	}

// 	r.Set(StatusOK, "OK", Map{"Project permissions id": permissionsId})
// 	return r
// }

// func (s *ProjectPermsService) Delete(userId, projectId, memberId int) *models.ApiResponse {
// 	r := &models.ApiResponse{}

// 	permissions, err := s.repo.Get(projectId, userId)
// 	if err != nil {
// 		if err.Error() == DbResultNotFound {
// 			r.Error(StatusNotFound, "Request author is not project member")
// 			return r
// 		}
// 		r.Error(StatusInternalServerError, err.Error())
// 		return r
// 	}

// 	if permissions.Admin != true {
// 		r.Error(StatusForbidden, "Request author is not project admin")
// 		return r
// 	}

// 	_, err = s.repo.Get(projectId, memberId)
// 	if err != nil {
// 		if err.Error() == DbResultNotFound {
// 			r.Error(StatusNotFound, "Excluding user is not project member")
// 			return r
// 		}
// 		r.Error(StatusInternalServerError, err.Error())
// 		return r
// 	}

// 	project, err := s.projectRepo.GetById(projectId)
// 	if err != nil {
// 		r.Error(StatusInternalServerError, err.Error())
// 		return r
// 	}
// 	if project.OwnerId == memberId {
// 		r.Error(StatusBadRequest, "You can't exclude project owner")
// 		return r
// 	}

// 	err = s.repo.Delete(projectId, memberId)
// 	if err != nil {
// 		r.Error(StatusInternalServerError, err.Error())
// 		return r
// 	}

// 	r.Set(StatusOK, "OK", Map{})
// 	return r
// }

// func (s *ProjectPermsService) Update(userId, projectId, memberId int, ProjectPerms *models.UpdatePermission) *models.ApiResponse {
// 	r := &models.ApiResponse{}

// 	permissions, err := s.repo.Get(projectId, userId)
// 	if err != nil {
// 		if err.Error() == DbResultNotFound {
// 			r.Error(StatusNotFound, "Request author is not project member")
// 			return r
// 		}
// 		r.Error(StatusInternalServerError, err.Error())
// 		return r
// 	}

// 	if permissions.Admin != true {
// 		r.Error(StatusForbidden, "Request author is not project admin")
// 		return r
// 	}

// 	_, err = s.repo.Get(projectId, memberId)
// 	if err != nil {
// 		if err.Error() == DbResultNotFound {
// 			r.Error(StatusNotFound, "Updating user is not project member")
// 			return r
// 		}
// 		r.Error(StatusInternalServerError, err.Error())
// 		return r
// 	}

// 	project, err := s.projectRepo.GetById(projectId) // TODO изменение прав автора
// 	if err != nil {
// 		r.Error(StatusInternalServerError, err.Error())
// 		return r
// 	}
// 	if project.OwnerId == memberId {
// 		r.Error(StatusBadRequest, "You can't update project owner permissions")
// 		return r
// 	}

// 	err = s.repo.Update(projectId, memberId, ProjectPerms)
// 	if err != nil {
// 		r.Error(StatusInternalServerError, err.Error())
// 		return r
// 	}

// 	r.Set(StatusOK, "OK", Map{})
// 	return r
// }
