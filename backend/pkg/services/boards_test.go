package services

import (
	"errors"
	"testing"
	"yak/backend/pkg/models"

	mock_repositories "yak/backend/pkg/repositories/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type BoardBuilder struct {
	Board *models.Board
}

func NewBoardBuilder() *BoardBuilder {
	board := &models.Board{
		ProjectId: 1,
		OwnerId:   1,
		DefaultPermissions: &models.Permission{
			Read:  true,
			Write: true,
			Admin: false,
		},
		Datetimes: &models.Datetimes{
			Created:  1,
			Updated:  1,
			Accessed: 1,
		},
		Title: "Default Title",
	}
	return &BoardBuilder{Board: board}
}

func (b *BoardBuilder) build() *models.Board {
	return b.Board
}

func (b *BoardBuilder) withTitle(title string) *BoardBuilder {
	b.Board.Title = title
	return b
}

func (b *BoardBuilder) withoutPerm() *BoardBuilder {
	b.Board.DefaultPermissions = nil
	return b
}

func (b *BoardBuilder) withOwner(id int) *BoardBuilder {
	b.Board.OwnerId = id
	return b
}

func (b *BoardBuilder) withProject(id int) *BoardBuilder {
	b.Board.ProjectId = id
	return b
}

func (b *BoardBuilder) withPerm(r, w, a bool) *BoardBuilder {
	perm := &models.Permission{
		Read:  r,
		Write: w,
		Admin: a,
	}
	b.Board.DefaultPermissions = perm
	return b
}

func (b *BoardBuilder) withDate(c, u, a int64) *BoardBuilder {
	date := &models.Datetimes{
		Created:  c,
		Updated:  u,
		Accessed: a,
	}
	b.Board.Datetimes = date
	return b
}

func TestBoardService_Create(t *testing.T) {
	type args struct {
		userId    int
		projectId int
		board     *models.Board
	}
	type projectMockBehavior func(r *mock_repositories.MockProject, userId, projectId int)
	type mockBehavior func(r *mock_repositories.MockBoard, userId int, board *models.Board)

	tests := []struct {
		name                string
		input               args
		projectMock         projectMockBehavior
		mock                mockBehavior
		expectedApiResponse *models.ApiResponse
	}{
		{
			name: "Ok",
			input: args{
				userId: 1,
				// board: &models.Board{
				// 	ProjectId: 1,
				// 	DefaultPermissions: &models.Permission{
				// 		Read:  true,
				// 		Write: true,
				// 		Admin: false,
				// 	},
				// 	Datetimes: &models.Datetimes{
				// 		Created:  1,
				// 		Updated:  1,
				// 		Accessed: 1,
				// 	},
				// 	Title: "New Test Board",
				// },
				board: NewBoardBuilder().withTitle("Board Builder").build(),
			},
			projectMock: func(r *mock_repositories.MockProject, userId, projectId int) {
				r.EXPECT().GetPermissions(userId, projectId).Return(&models.Permission{true, true, true}, nil)
			},
			mock: func(r *mock_repositories.MockBoard, userId int, board *models.Board) {
				r.EXPECT().Create(userId, board).Return(1, nil)
			},
			expectedApiResponse: &models.ApiResponse{
				Code: StatusOK,
				Data: Map{"boardId": 1},
			},
		},
		{
			name: "Forbidden",
			input: args{
				userId: 1,
				board:  NewBoardBuilder().withTitle("Board Builder").build(),
			},
			projectMock: func(r *mock_repositories.MockProject, userId, projectId int) {
				r.EXPECT().GetPermissions(userId, projectId).Return(nil, errors.New("Forbidden"))
			},
			mock: func(r *mock_repositories.MockBoard, userId int, board *models.Board) {},
			expectedApiResponse: &models.ApiResponse{
				Code: StatusForbidden,
			},
		},
		{
			name: "Null Default Permissions",
			input: args{
				userId: 1,
				board:  NewBoardBuilder().withTitle("Board Builder").withoutPerm().build(),
			},
			projectMock: func(r *mock_repositories.MockProject, userId, projectId int) {
				r.EXPECT().GetPermissions(userId, projectId).Return(&models.Permission{true, true, true}, nil)
			},
			mock: func(r *mock_repositories.MockBoard, userId int, board *models.Board) {
				r.EXPECT().Create(userId, board).Return(1, nil)
			},
			expectedApiResponse: &models.ApiResponse{
				Code: StatusOK,
				Data: Map{"boardId": 1},
			},
		},
		{
			name: "Board Repo Error",
			input: args{
				userId: 1,
				board:  NewBoardBuilder().withTitle("Board Builder").withoutPerm().build(),
			},
			projectMock: func(r *mock_repositories.MockProject, userId, projectId int) {
				r.EXPECT().GetPermissions(userId, projectId).Return(&models.Permission{true, true, true}, nil)
			},
			mock: func(r *mock_repositories.MockBoard, userId int, board *models.Board) {
				r.EXPECT().Create(userId, board).Return(0, errors.New("repo error"))
			},
			expectedApiResponse: &models.ApiResponse{
				Code: StatusInternalServerError,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_repositories.NewMockBoard(c)
			projectRepo := mock_repositories.NewMockProject(c)
			test.projectMock(projectRepo, test.input.userId, test.input.projectId)
			test.mock(repo, test.input.userId, test.input.board)
			s := &BoardService{repo: repo, projectRepo: projectRepo}

			got := s.Create(test.input.userId, test.input.projectId, test.input.board)
			assert.Equal(t, test.expectedApiResponse.Code, got.Code)
			if test.expectedApiResponse.Code == StatusOK {
				assert.Equal(t, test.expectedApiResponse.Data, got.Data)
			}
		})
	}
}
