// +build integration

package services

import (
	"testing"

	"github.com/architectv/networking-course-project/backend/pkg/builders"
	"github.com/architectv/networking-course-project/backend/pkg/models"
	"github.com/architectv/networking-course-project/backend/pkg/repositories/postgres"

	"github.com/stretchr/testify/assert"
)

func Test_Integration_BoardService_Create(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	type args struct {
		userId    int
		projectId int
		board     *models.Board
	}
	db, err := prepareTestDatabase()
	if err != nil {
		t.Fatalf(err.Error())
	}

	repo := postgres.NewBoardPg(db)
	projectRepo := postgres.NewProjectPg(db)
	s := NewBoardService(repo, projectRepo)

	tests := []struct {
		name                string
		input               args
		want                int
		expectedApiResponse *models.ApiResponse
	}{
		{
			name: "Ok",
			input: args{
				userId:    1,
				projectId: 1,
				board:     builders.NewBoardBuilder().WithTitle("Board Builder").Build(),
			},
			expectedApiResponse: &models.ApiResponse{
				Code: StatusOK,
				Data: Map{"boardId": 10001},
			},
		},
		{
			name: "Forbidden",
			input: args{
				userId:    3,
				projectId: 2,
				board:     builders.NewBoardBuilder().WithTitle("Board Builder").Build(),
			},
			expectedApiResponse: &models.ApiResponse{
				Code: StatusForbidden,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := s.Create(test.input.userId, test.input.projectId, test.input.board)
			assert.Equal(t, test.expectedApiResponse.Code, got.Code)
			if test.expectedApiResponse.Code == StatusOK {
				assert.Equal(t, test.expectedApiResponse.Data, got.Data)
			}
		})
	}
}
