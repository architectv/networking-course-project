package postgres

import (
	"errors"
	"testing"
	"yak/backend/pkg/data_builders"
	"yak/backend/pkg/models"

	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestObjectpermsPg_Create(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewObjectPermsPg(db)

	type args struct {
		objectId   int
		memberId   int
		objectType int
		perms      *models.Permission
	}
	type mockBehavior func(args args)

	tests := []struct {
		name    string
		mock    mockBehavior
		input   args
		want    int
		wantErr bool
	}{
		{
			name: "Ok for project",
			input: args{
				objectId:   1,
				memberId:   1,
				objectType: IsProject,
				perms:      data_builders.NewPermsBuilder().Build(),
			},
			want: 1,
			mock: func(args args) {
				mock.ExpectBegin()

				permId := 1
				projectUserId := 1

				rows := sqlmock.NewRows([]string{"per.read", "per.write", "per.admin"}).
					RowError(0, errors.New(DbResultNotFound))
				mock.ExpectQuery("SELECT (.+) FROM project_users").
					WithArgs(args.objectId, args.memberId).WillReturnRows(rows)

				perms := args.perms
				defPermRows := sqlmock.NewRows([]string{"id"}).AddRow(permId)
				mock.ExpectQuery("INSERT INTO permissions").
					WithArgs(perms.Read, perms.Write, perms.Admin).
					WillReturnRows(defPermRows)

				dateRows := sqlmock.NewRows([]string{"id"}).AddRow(projectUserId)
				mock.ExpectQuery("INSERT INTO project_users").
					WithArgs(args.memberId, args.objectId, permId).
					WillReturnRows(dateRows)

				mock.ExpectCommit()
			},
		},
		{
			name: "Ok for board",
			input: args{
				objectId:   1,
				memberId:   1,
				objectType: IsBoard,
				perms:      data_builders.NewPermsBuilder().Build(),
			},
			want: 1,
			mock: func(args args) {
				mock.ExpectBegin()

				permId := 1
				projectUserId := 1

				rows := sqlmock.NewRows([]string{"per.read", "per.write", "per.admin"}).
					RowError(0, errors.New(DbResultNotFound))
				mock.ExpectQuery("SELECT (.+) FROM board_users").
					WithArgs(args.objectId, args.memberId).WillReturnRows(rows)

				perms := args.perms
				defPermRows := sqlmock.NewRows([]string{"id"}).AddRow(permId)
				mock.ExpectQuery("INSERT INTO permissions").
					WithArgs(perms.Read, perms.Write, perms.Admin).
					WillReturnRows(defPermRows)

				dateRows := sqlmock.NewRows([]string{"id"}).AddRow(projectUserId)
				mock.ExpectQuery("INSERT INTO board_users").
					WithArgs(args.memberId, args.objectId, permId).
					WillReturnRows(dateRows)

				mock.ExpectCommit()
			},
		},
		{
			name: "bad object type",
			input: args{
				objectId:   1,
				memberId:   1,
				objectType: 0,
				perms:      data_builders.NewPermsBuilder().Build(),
			},
			mock:    func(args args) {},
			wantErr: true,
		},
		{
			name: "Member already has permissions in the project",
			input: args{
				objectId:   1,
				memberId:   1,
				objectType: IsProject,
				perms:      data_builders.NewPermsBuilder().Build(),
			},
			want: 1,
			mock: func(args args) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"per.read", "per.write", "per.admin"}).
					AddRow(true, true, true)
				mock.ExpectQuery("SELECT (.+) FROM project_users").
					WithArgs(args.objectId, args.memberId).WillReturnRows(rows)
			},
			wantErr: true,
		},
		{
			name: "repo error for create in permissions",
			input: args{
				objectId:   1,
				memberId:   1,
				objectType: IsBoard,
				perms:      data_builders.NewPermsBuilder().Build(),
			},
			want: 1,
			mock: func(args args) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"per.read", "per.write", "per.admin"}).
					RowError(0, errors.New(DbResultNotFound))
				mock.ExpectQuery("SELECT (.+) FROM board_users").
					WithArgs(args.objectId, args.memberId).WillReturnRows(rows)

				perms := args.perms
				defPermRows := sqlmock.NewRows([]string{"id"}).
					RowError(0, errors.New("Some error"))
				mock.ExpectQuery("INSERT INTO permissions").
					WithArgs(perms.Read, perms.Write, perms.Admin).
					WillReturnRows(defPermRows)

				mock.ExpectCommit()
			},
			wantErr: true,
		},
		{ //TODO не отображается в покрытии тестов
			name: "repo error for create in project_users",
			input: args{
				objectId:   1,
				memberId:   1,
				objectType: IsBoard,
				perms:      data_builders.NewPermsBuilder().Build(),
			},
			want: 1,
			mock: func(args args) {
				mock.ExpectBegin()

				permId := 1

				rows := sqlmock.NewRows([]string{"per.read", "per.write", "per.admin"}).
					RowError(0, errors.New(DbResultNotFound))
				mock.ExpectQuery("SELECT (.+) FROM board_users").
					WithArgs(args.objectId, args.memberId).WillReturnRows(rows)

				perms := args.perms
				defPermRows := sqlmock.NewRows([]string{"id"}).AddRow(permId)
				mock.ExpectQuery("INSERT INTO permissions").
					WithArgs(perms.Read, perms.Write, perms.Admin).
					WillReturnRows(defPermRows)

				dateRows := sqlmock.NewRows([]string{"id"}).
					RowError(0, errors.New("Some error"))
				mock.ExpectQuery("INSERT INTO project_users").
					WithArgs(args.memberId, args.objectId, permId).
					WillReturnRows(dateRows)

				mock.ExpectCommit()
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.input)

			got, err := r.Create(tt.input.objectId, tt.input.memberId, tt.input.objectType, tt.input.perms)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
