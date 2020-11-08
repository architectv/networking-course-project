package services

import "yak/backend/pkg/repositories"

type ProjectService struct {
	repo repositories.Project
}

func NewProjectService(repo repositories.Project) *ProjectService {
	return &ProjectService{repo: repo}
}
