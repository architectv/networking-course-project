package services

import "yak/backend/pkg/repositories"

type TaskService struct {
	repo repositories.Task
}

func NewTaskService(repo repositories.Task) *TaskService {
	return &TaskService{repo: repo}
}
