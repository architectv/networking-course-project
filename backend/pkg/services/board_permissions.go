package services

import (
	"fmt"
	"yak/backend/pkg/models"
	"yak/backend/pkg/repositories"
)

type BoardPermsService struct {
	repo        repositories.ProjectPerms
	boardRepo   repositories.Board
	projectRepo repositories.Project
}

func NewBoardPermsService(repo repositories.ProjectPerms, boardRepo repositories.Board,
	projectRepo repositories.Project) *BoardPermsService {
	return &BoardPermsService{repo: repo, boardRepo: boardRepo, projectRepo: projectRepo}
}

func (s *BoardPermsService) Get(userId, projectId, boardId, memberId int) *models.ApiResponse {
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

	_, err = s.repo.Get(boardId, userId, IsBoard)
	if err != nil {
		if err.Error() == DbResultNotFound {
			r.Error(StatusNotFound, "Request author is not board member")
			return r
		}
		r.Error(StatusInternalServerError, err.Error())
		return r
	}

	_, err = s.repo.Get(projectId, memberId, IsProject)
	if err != nil {
		if err.Error() == DbResultNotFound {
			r.Error(StatusNotFound, "Board member is not project member")
			return r
		}
		r.Error(StatusInternalServerError, err.Error())
		return r
	}

	permissions, err := s.repo.Get(boardId, memberId, IsBoard)
	if err != nil {
		if err.Error() == DbResultNotFound {
			r.Error(StatusNotFound, "Board member not found")
			return r
		}
		r.Error(StatusInternalServerError, err.Error())
		return r
	}

	r.Set(StatusOK, "OK", Map{"permissions": permissions})
	return r
}

func (s *BoardPermsService) Create(userId, projectId, boardId, memberId int, boardPerms *models.Permission) *models.ApiResponse {
	r := &models.ApiResponse{}

	if err := permsValidation(boardPerms); err != nil {
		if err.Error() == ErrPermsIsNotDefined {
			board, err := s.boardRepo.GetById(boardId)
			if err != nil {
				r.Error(StatusInternalServerError, err.Error())
				return r
			}
			if board.DefaultPermissions == nil { // TODO права по умолчанию должны обязательно указываться при создании проекта или доски
				r.Error(StatusInternalServerError, "Default permissions is not defined")
				return r
			}
			boardPerms = board.DefaultPermissions
		} else {
			r.Error(StatusBadRequest, err.Error())
			return r
		}
	}

	_, err := s.repo.Get(projectId, userId, IsProject)
	if err != nil {
		if err.Error() == DbResultNotFound {
			r.Error(StatusNotFound, "Request author is not project member")
			return r
		}
		r.Error(StatusInternalServerError, err.Error())
		return r
	}
	permissions, err := s.repo.Get(boardId, userId, IsBoard)
	if err != nil {
		if err.Error() == DbResultNotFound {
			r.Error(StatusNotFound, "Request author is not board member")
			return r
		}
		r.Error(StatusInternalServerError, err.Error())
		return r
	}

	if permissions.Admin != true {
		r.Error(StatusForbidden, "Request author is not board admin")
		return r
	}

	_, err = s.repo.Get(projectId, memberId, IsProject)
	if err != nil {
		if err.Error() == DbResultNotFound {
			r.Error(StatusNotFound, "New board member is not project member")
			return r
		}
		r.Error(StatusInternalServerError, err.Error())
		return r
	}

	permissionsId, err := s.repo.Create(boardId, memberId, IsBoard, boardPerms)
	if err != nil {
		r.Error(StatusInternalServerError, err.Error())
		return r
	}

	r.Set(StatusOK, "OK", Map{"Board permissions id": permissionsId})
	return r
}

func (s *BoardPermsService) Delete(userId, projectId, boardId, memberId int) *models.ApiResponse {
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
	permissions, err := s.repo.Get(boardId, userId, IsBoard)
	if err != nil {
		if err.Error() == DbResultNotFound {
			r.Error(StatusNotFound, "Request author is not board member")
			return r
		}
		r.Error(StatusInternalServerError, err.Error())
		return r
	}

	if permissions.Admin != true {
		r.Error(StatusForbidden, "Request author is not board admin")
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

	_, err = s.repo.Get(boardId, memberId, IsBoard)
	if err != nil {
		if err.Error() == DbResultNotFound {
			r.Error(StatusNotFound, "Excluding user is not board member")
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

	var projectOwnerId int
	boardsIds, err := s.boardRepo.GetCountByOwnerId(projectId, memberId)
	fmt.Println("boardsIds ", boardsIds, projectId, memberId)
	if boardsIds != 0 {
		if project.OwnerId != userId {
			r.Error(StatusBadRequest, "Exclude owner boards can only be project owner")
			return r
		} else {
			projectOwnerId = userId
		}
	}
	fmt.Println(projectId, memberId, projectOwnerId, IsBoard)
	err = s.repo.Delete(boardId, memberId, projectOwnerId, IsBoard)
	if err != nil {
		r.Error(StatusInternalServerError, err.Error())
		return r
	}

	r.Set(StatusOK, "OK", Map{})
	return r
}

// func (s *BoardPermsService) Update(userId, projectId, memberId int, ProjectPerms *models.UpdatePermission) *models.ApiResponse {
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
