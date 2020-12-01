package services

import (
	"errors"
	"testing"
	"yak/backend/pkg/models"

	mock_repositories "yak/backend/pkg/repositories/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type LabelBuilder struct {
	Label *models.Label
}

func NewLabelBuilder() *LabelBuilder {
	label := &models.Label{
		BoardId: 1,
		Name:    "Default Label Name",
		Color:   255,
	}
	return &LabelBuilder{Label: label}
}

func (b *LabelBuilder) build() *models.Label {
	return b.Label
}

func (b *LabelBuilder) withName(name string) *LabelBuilder {
	b.Label.Name = name
	return b
}

func (b *LabelBuilder) withBoard(id int) *LabelBuilder {
	b.Label.BoardId = id
	return b
}

func (b *LabelBuilder) withColor(color uint32) *LabelBuilder {
	b.Label.Color = color
	return b
}

func TestLabelService_Create(t *testing.T) {
	type args struct {
		userId    int
		projectId int
		boardId   int
		label     *models.Label
	}
	type mockBehavior func(r *mock_repositories.MockLabel, label *models.Label)
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
				label:     NewLabelBuilder().withName("Label Builder").build(),
			},
			projectMock: func(r *mock_repositories.MockProject, userId, projectId int) {
				r.EXPECT().GetPermissions(userId, projectId).Return(&models.Permission{true, true, true}, nil)
			},
			boardMock: func(r *mock_repositories.MockBoard, userId, boardId int) {
				r.EXPECT().GetPermissions(userId, boardId).Return(&models.Permission{true, true, true}, nil)
			},
			mock: func(r *mock_repositories.MockLabel, label *models.Label) {
				r.EXPECT().Create(label).Return(1, nil)
			},
			expectedApiResponse: &models.ApiResponse{
				Code: StatusOK,
				Data: Map{"labelId": 1},
			},
		},
		{
			name: "Project Perm Failed",
			input: args{
				userId:    1,
				projectId: 1,
				boardId:   1,
				label:     NewLabelBuilder().withName("Label Builder").build(),
			},
			projectMock: func(r *mock_repositories.MockProject, userId, projectId int) {
				r.EXPECT().GetPermissions(userId, projectId).Return(nil, errors.New("Forbidden"))
			},
			boardMock: func(r *mock_repositories.MockBoard, userId, boardId int) {},
			mock:      func(r *mock_repositories.MockLabel, label *models.Label) {},
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
				label:     NewLabelBuilder().withName("Label Builder").build(),
			},
			projectMock: func(r *mock_repositories.MockProject, userId, projectId int) {
				r.EXPECT().GetPermissions(userId, projectId).Return(&models.Permission{true, true, true}, nil)
			},
			boardMock: func(r *mock_repositories.MockBoard, userId, boardId int) {
				r.EXPECT().GetPermissions(userId, boardId).Return(nil, errors.New("Forbidden"))
			},
			mock: func(r *mock_repositories.MockLabel, label *models.Label) {},
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
				label:     NewLabelBuilder().withName("Label Builder").build(),
			},
			projectMock: func(r *mock_repositories.MockProject, userId, projectId int) {
				r.EXPECT().GetPermissions(userId, projectId).Return(&models.Permission{true, true, true}, nil)
			},
			boardMock: func(r *mock_repositories.MockBoard, userId, boardId int) {
				r.EXPECT().GetPermissions(userId, boardId).Return(&models.Permission{true, true, true}, nil)
			},
			mock: func(r *mock_repositories.MockLabel, label *models.Label) {
				r.EXPECT().Create(label).Return(0, errors.New("repo error"))
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

			repo := mock_repositories.NewMockLabel(c)
			projectRepo := mock_repositories.NewMockProject(c)
			boardRepo := mock_repositories.NewMockBoard(c)
			test.projectMock(projectRepo, test.input.userId, test.input.projectId)
			test.boardMock(boardRepo, test.input.userId, test.input.boardId)
			test.mock(repo, test.input.label)
			s := &LabelService{repo: repo, projectRepo: projectRepo, boardRepo: boardRepo}

			got := s.Create(test.input.userId, test.input.projectId, test.input.boardId, test.input.label)
			assert.Equal(t, test.expectedApiResponse.Code, got.Code)
			if test.expectedApiResponse.Code == StatusOK {
				assert.Equal(t, test.expectedApiResponse.Data, got.Data)
			}
		})
	}
}
