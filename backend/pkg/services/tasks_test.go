package services

import (
	"errors"
	"testing"
	"yak/backend/pkg/models"

	mock_repositories "yak/backend/pkg/repositories/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type TaskBuilder struct {
	Task *models.Task
}

func NewTaskBuilder() *TaskBuilder {
	task := &models.Task{
		ListId:   1,
		Title:    "Default List Title",
		Position: 1,
		Datetimes: &models.Datetimes{
			Created:  1,
			Updated:  1,
			Accessed: 1,
		},
	}
	return &TaskBuilder{Task: task}
}

func (b *TaskBuilder) build() *models.Task {
	return b.Task
}

func (b *TaskBuilder) withTitle(title string) *TaskBuilder {
	b.Task.Title = title
	return b
}

func (b *TaskBuilder) withList(id int) *TaskBuilder {
	b.Task.ListId = id
	return b
}

func (b *TaskBuilder) withPos(id int) *TaskBuilder {
	b.Task.Position = id
	return b
}

func (b *TaskBuilder) withDate(c, u, a int64) *TaskBuilder {
	date := &models.Datetimes{
		Created:  c,
		Updated:  u,
		Accessed: a,
	}
	b.Task.Datetimes = date
	return b
}

func TestTaskService_Create(t *testing.T) {
	type args struct {
		userId    int
		projectId int
		boardId   int
		listId    int
		task      *models.Task
	}
	type mockBehavior func(r *mock_repositories.MockTask, task *models.Task)
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
				task:      NewTaskBuilder().withTitle("Task Builder").build(),
			},
			projectMock: func(r *mock_repositories.MockProject, userId, projectId int) {
				r.EXPECT().GetPermissions(userId, projectId).Return(&models.Permission{true, true, true}, nil)
			},
			boardMock: func(r *mock_repositories.MockBoard, userId, boardId int) {
				r.EXPECT().GetPermissions(userId, boardId).Return(&models.Permission{true, true, true}, nil)
			},
			mock: func(r *mock_repositories.MockTask, task *models.Task) {
				r.EXPECT().Create(task).Return(1, nil)
			},
			expectedApiResponse: &models.ApiResponse{
				Code: StatusOK,
				Data: Map{"taskId": 1},
			},
		},
		{
			name: "Project Perm Failed",
			input: args{
				userId:    1,
				projectId: 1,
				boardId:   1,
				task:      NewTaskBuilder().withTitle("Task Builder").build(),
			},
			projectMock: func(r *mock_repositories.MockProject, userId, projectId int) {
				r.EXPECT().GetPermissions(userId, projectId).Return(nil, errors.New("Forbidden"))
			},
			boardMock: func(r *mock_repositories.MockBoard, userId, boardId int) {},
			mock:      func(r *mock_repositories.MockTask, task *models.Task) {},
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
				task:      NewTaskBuilder().withTitle("Task Builder").build(),
			},
			projectMock: func(r *mock_repositories.MockProject, userId, projectId int) {
				r.EXPECT().GetPermissions(userId, projectId).Return(&models.Permission{true, true, true}, nil)
			},
			boardMock: func(r *mock_repositories.MockBoard, userId, boardId int) {
				r.EXPECT().GetPermissions(userId, boardId).Return(nil, errors.New("Forbidden"))
			},
			mock: func(r *mock_repositories.MockTask, task *models.Task) {},
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
				task:      NewTaskBuilder().withTitle("Task Builder").build(),
			},
			projectMock: func(r *mock_repositories.MockProject, userId, projectId int) {
				r.EXPECT().GetPermissions(userId, projectId).Return(&models.Permission{true, true, true}, nil)
			},
			boardMock: func(r *mock_repositories.MockBoard, userId, boardId int) {
				r.EXPECT().GetPermissions(userId, boardId).Return(&models.Permission{true, true, true}, nil)
			},
			mock: func(r *mock_repositories.MockTask, task *models.Task) {
				r.EXPECT().Create(task).Return(0, errors.New("repo error"))
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

			repo := mock_repositories.NewMockTask(c)
			projectRepo := mock_repositories.NewMockProject(c)
			boardRepo := mock_repositories.NewMockBoard(c)
			test.projectMock(projectRepo, test.input.userId, test.input.projectId)
			test.boardMock(boardRepo, test.input.userId, test.input.boardId)
			test.mock(repo, test.input.task)
			s := &TaskService{repo: repo, projectRepo: projectRepo, boardRepo: boardRepo}

			got := s.Create(test.input.userId, test.input.projectId, test.input.boardId,
				test.input.listId, test.input.task)
			assert.Equal(t, test.expectedApiResponse.Code, got.Code)
			if test.expectedApiResponse.Code == StatusOK {
				assert.Equal(t, test.expectedApiResponse.Data, got.Data)
			}
		})
	}
}
