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

func (s *UserService) GetById(id string) (models.User, error) {
	return s.repo.GetById(id)
}

func (s *UserService) Create(user models.User) (string, error) {
	return s.repo.Create(user)
}