package services

import (
	"yak/backend/pkg/models"
	"yak/backend/pkg/repositories"
)

type UserService struct {
	repo repositories.User
}

func NewUserService(repo repositories.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAll() ([]models.User, error) {
	return s.repo.GetAll()
}
