package services

import "yak/backend/pkg/repositories"

type BoardService struct {
	repo repositories.Board
}

func NewBoardService(repo repositories.Board) *BoardService {
	return &BoardService{repo: repo}
}
