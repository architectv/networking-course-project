package services

import (
	"errors"
	"testing"
	"yak/backend/pkg/models"

	mock_repositories "yak/backend/pkg/repositories/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type ListBuilder struct {
	List *models.TaskList
}

func NewListBuilder() *ListBuilder {
	list := &models.TaskList{
		BoardId:  1,
		Title:    "Default List Title",
		Position: 1,
	}
	return &ListBuilder{List: list}
}

func (b *ListBuilder) build() *models.TaskList {
	return b.List
}

func (b *ListBuilder) withTitle(title string) *ListBuilder {
	b.List.Title = title
	return b
}

func (b *ListBuilder) withBoard(id int) *ListBuilder {
	b.List.BoardId = id
	return b
}

func (b *ListBuilder) withPos(id int) *ListBuilder {
	b.List.Position = id
	return b
}

func TestTaskListService_Create(t *testing.T) {
	type args struct {
		userId    int
		projectId int
		boardId   int
		list      *models.TaskList
	}
	type mockBehavior func(r *mock_repositories.MockTaskList, list *models.TaskList)
	type projectMockBehavior func(r *mock_repositories.MockProject, userId, projectId int)
	type boardMockBehavior func(r *mock_repositories.MockBoard, userId, boardId int)

	tests := []struct {
		name                string
		input               args
		projectMock         projectMockBehavior
		boardMock           boardMockBehavior
		mock                mockBehavior
		expectedApiResponse *models.ApiResponse
	}{
		{
			name: "Ok",
			input: args{
				userId:    1,
				projectId: 1,
				boardId:   1,
				list:      NewListBuilder().withTitle("List Builder").build(),
			},
			projectMock: func(r *mock_repositories.MockProject, userId, projectId int) {
				r.EXPECT().GetPermissions(userId, projectId).Return(&models.Permission{true, true, true}, nil)
			},
			boardMock: func(r *mock_repositories.MockBoard, userId, boardId int) {
				r.EXPECT().GetPermissions(userId, boardId).Return(&models.Permission{true, true, true}, nil)
			},
			mock: func(r *mock_repositories.MockTaskList, list *models.TaskList) {
				r.EXPECT().Create(list).Return(1, nil)
			},
			expectedApiResponse: &models.ApiResponse{
				Code: StatusOK,
				Data: Map{"listId": 1},
			},
		},
		{
			name: "Project Perm Failed",
			input: args{
				userId:    1,
				projectId: 1,
				boardId:   1,
				list:      NewListBuilder().withTitle("List Builder").build(),
			},
			projectMock: func(r *mock_repositories.MockProject, userId, projectId int) {
				r.EXPECT().GetPermissions(userId, projectId).Return(nil, errors.New("Forbidden"))
			},
			boardMock: func(r *mock_repositories.MockBoard, userId, boardId int) {},
			mock:      func(r *mock_repositories.MockTaskList, list *models.TaskList) {},
			expectedApiResponse: &models.ApiResponse{
				Code: StatusForbidden,
			},
		},
		{
			name: "Board Perm Failed",
			input: args{
				userId:    1,
				projectId: 1,
				boardId:   1,
				list:      NewListBuilder().withTitle("List Builder").build(),
			},
			projectMock: func(r *mock_repositories.MockProject, userId, projectId int) {
				r.EXPECT().GetPermissions(userId, projectId).Return(&models.Permission{true, true, true}, nil)
			},
			boardMock: func(r *mock_repositories.MockBoard, userId, boardId int) {
				r.EXPECT().GetPermissions(userId, boardId).Return(nil, errors.New("Forbidden"))
			},
			mock: func(r *mock_repositories.MockTaskList, list *models.TaskList) {},
			expectedApiResponse: &models.ApiResponse{
				Code: StatusForbidden,
			},
		},
		{
			name: "Repo Error",
			input: args{
				userId:    1,
				projectId: 1,
				boardId:   1,
				list:      NewListBuilder().withTitle("List Builder").build(),
			},
			projectMock: func(r *mock_repositories.MockProject, userId, projectId int) {
				r.EXPECT().GetPermissions(userId, projectId).Return(&models.Permission{true, true, true}, nil)
			},
			boardMock: func(r *mock_repositories.MockBoard, userId, boardId int) {
				r.EXPECT().GetPermissions(userId, boardId).Return(&models.Permission{true, true, true}, nil)
			},
			mock: func(r *mock_repositories.MockTaskList, list *models.TaskList) {
				r.EXPECT().Create(list).Return(0, errors.New("repo error"))
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

			repo := mock_repositories.NewMockTaskList(c)
			projectRepo := mock_repositories.NewMockProject(c)
			boardRepo := mock_repositories.NewMockBoard(c)
			test.projectMock(projectRepo, test.input.userId, test.input.projectId)
			test.boardMock(boardRepo, test.input.userId, test.input.boardId)
			test.mock(repo, test.input.list)
			s := &TaskListService{repo: repo, projectRepo: projectRepo, boardRepo: boardRepo}

			got := s.Create(test.input.userId, test.input.projectId, test.input.boardId, test.input.list)
			assert.Equal(t, test.expectedApiResponse.Code, got.Code)
			if test.expectedApiResponse.Code == StatusOK {
				assert.Equal(t, test.expectedApiResponse.Data, got.Data)
			}
		})
	}
}
