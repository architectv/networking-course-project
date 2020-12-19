// +build integration
package services

import (
	"errors"
	"testing"
	"yak/backend/pkg/models"
	"yak/backend/pkg/repositories/postgres"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func Test_Integration_ProjectService_Create(t *testing.T) {
	type args struct {
		userId  int
		project *models.Project
	}
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := postgres.NewProjectPg(db)
	s := NewProjectService(r)
	type mockBehavior func(args args, id int)

	tests := []struct {
		name                string
		mock                mockBehavior
		input               args
		want                int
		expectedApiResponse *models.ApiResponse
	}{
		{
			name: "Ok",
			input: args{
				userId: 1,
				project: &models.Project{
					OwnerId: 1,
					DefaultPermissions: &models.Permission{
						Read:  true,
						Write: true,
						Admin: true,
					},
					Datetimes: &models.Datetimes{
						Created:  1,
						Updated:  1,
						Accessed: 1,
					},
					Title:       "New Test Project",
					Description: "Some Description",
				},
			},
			mock: func(args args, id int) {
				mock.ExpectBegin()

				defPermissionId := 1
				datetimesId := 1

				defPerm := args.project.DefaultPermissions
				defPermRows := sqlmock.NewRows([]string{"id"}).AddRow(defPermissionId)
				mock.ExpectQuery("INSERT INTO permissions").
					WithArgs(defPerm.Read, defPerm.Write, defPerm.Admin).
					WillReturnRows(defPermRows)

				dateRows := sqlmock.NewRows([]string{"id"}).AddRow(datetimesId)
				mock.ExpectQuery("INSERT INTO datetimes").
					WillReturnRows(dateRows)

				project := args.project
				projectRows := sqlmock.NewRows([]string{"id"}).AddRow(id)
				mock.ExpectQuery("INSERT INTO projects").WithArgs(project.OwnerId, defPermissionId,
					datetimesId, project.Title, project.Description).
					WillReturnRows(projectRows)

				perm := &models.Permission{
					Read:  true,
					Write: true,
					Admin: true,
				}
				permissionId := defPermissionId + 1
				permRows := sqlmock.NewRows([]string{"id"}).AddRow(permissionId)
				mock.ExpectQuery("INSERT INTO permissions").
					WithArgs(perm.Read, perm.Write, perm.Admin).
					WillReturnRows(permRows)

				mock.ExpectExec("INSERT INTO project_users").WithArgs(project.OwnerId, id, permissionId).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
			want: 1,
			expectedApiResponse: &models.ApiResponse{
				Code: StatusOK,
				Data: Map{"projectId": 1},
			},
		},
		{
			name: "Empty Fields",
			input: args{
				userId: 1,
				project: &models.Project{
					OwnerId: 1,
					DefaultPermissions: &models.Permission{
						Read:  true,
						Write: true,
						Admin: true,
					},
					Datetimes: &models.Datetimes{
						Created:  1,
						Updated:  1,
						Accessed: 1,
					},
					Title:       "",
					Description: "Some Description",
				},
			},
			mock: func(args args, id int) {
				mock.ExpectBegin()

				defPermissionId := 1
				datetimesId := 1

				defPerm := args.project.DefaultPermissions
				defPermRows := sqlmock.NewRows([]string{"id"}).AddRow(defPermissionId)
				mock.ExpectQuery("INSERT INTO permissions").
					WithArgs(defPerm.Read, defPerm.Write, defPerm.Admin).
					WillReturnRows(defPermRows)

				dateRows := sqlmock.NewRows([]string{"id"}).AddRow(datetimesId)
				mock.ExpectQuery("INSERT INTO datetimes").
					WillReturnRows(dateRows)

				project := args.project
				projectRows := sqlmock.NewRows([]string{"id"}).AddRow(id).RowError(0, errors.New("insert error"))
				mock.ExpectQuery("INSERT INTO projects").WithArgs(project.OwnerId, defPermissionId,
					datetimesId, project.Title, project.Description).
					WillReturnRows(projectRows)

				mock.ExpectRollback()
			},
			want: 1,
			expectedApiResponse: &models.ApiResponse{
				Code: StatusInternalServerError,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			test.mock(test.input, test.want)

			got := s.Create(test.input.userId, test.input.project)
			assert.Equal(t, test.expectedApiResponse.Code, got.Code)
			if test.expectedApiResponse.Code == StatusOK {
				assert.Equal(t, test.expectedApiResponse.Data, got.Data)
			}
		})
	}
}
