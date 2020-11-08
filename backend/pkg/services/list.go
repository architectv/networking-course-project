package services

import "yak/backend/pkg/repositories"

type TaskListService struct {
	repo repositories.TaskList
}

func NewTaskListService(repo repositories.TaskList) *TaskListService {
	return &TaskListService{repo: repo}
}
