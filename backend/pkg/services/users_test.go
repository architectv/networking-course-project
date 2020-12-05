package services

import (
	"fmt"
	"testing"
	"yak/backend/pkg/builders"
	"yak/backend/pkg/models"
	"yak/backend/pkg/repositories/postgres"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

const (
	UsernameTestDB = "postgres"
	PasswordTestDB = "123matan123"
	HostTestDB     = "localhost"
	PortTestDB     = "5432"
	DBnameTestDB   = "yak_test_real_db"
	SslmodeTestDB  = "disable"
)

func openTestDatabase() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		HostTestDB, PortTestDB, UsernameTestDB, DBnameTestDB, PasswordTestDB, SslmodeTestDB))
	return db, err
}

func prepareTestDatabase() (*sqlx.DB, error) {
	db, err := openTestDatabase()
	if err != nil {
		return nil, err
	}

	fixtures, err := testfixtures.New(
		testfixtures.Database(db.DB),
		testfixtures.Dialect("postgres"),
		testfixtures.Directory("fixtures"),
	)
	if err != nil {
		return nil, err
	}

	err = fixtures.Load()
	return db, err
}

func TestUserService_Create(t *testing.T) {
	db, err := prepareTestDatabase()
	if err != nil {
		t.Fatalf(err.Error())
	}

	type args struct {
		user *models.User
	}

	tests := []struct {
		name                string
		input               args
		expectedApiResponse *models.ApiResponse
	}{
		{
			name: "Ok",
			input: args{
				user: builders.NewUserBuilder().WithNickname("User Builder").Build(),
			},
			expectedApiResponse: &models.ApiResponse{
				Code: StatusOK,
				Data: Map{"uid": 10001},
			},
		},
		{
			name: "Already Exists",
			input: args{
				user: builders.NewUserBuilder().WithNickname("User Builder").Build(),
			},
			expectedApiResponse: &models.ApiResponse{
				Code: StatusConflict,
			},
		},
		{
			name: "Repo Error",
			input: args{
				user: builders.NewUserBuilder().WithNickname("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA").Build(),
			},
			expectedApiResponse: &models.ApiResponse{
				Code: StatusInternalServerError,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			repo := postgres.NewUserPg(db)
			s := &UserService{repo: repo}

			got := s.Create(test.input.user)
			assert.Equal(t, test.expectedApiResponse.Code, got.Code)
			if test.expectedApiResponse.Code == StatusOK {
				assert.Equal(t, test.expectedApiResponse.Data, got.Data)
			}
		})
	}
}

func TestUserService_Get(t *testing.T) {
	db, err := prepareTestDatabase()
	if err != nil {
		t.Fatalf(err.Error())
	}

	type args struct {
		id int
	}

	tests := []struct {
		name                string
		input               args
		expectedApiResponse *models.ApiResponse
	}{
		{
			name: "Ok",
			input: args{
				id: 1,
			},
			expectedApiResponse: &models.ApiResponse{
				Code: StatusOK,
				Data: Map{"user": &models.User{
					Id:       1,
					Nickname: "test",
					Email:    "test@.mail.ru",
					Password: "",
					Avatar:   "photo1",
				}},
			},
		},
		{
			name: "Repo err",
			input: args{
				id: 3,
			},
			expectedApiResponse: &models.ApiResponse{
				Code: StatusInternalServerError,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			repo := postgres.NewUserPg(db)
			s := &UserService{repo: repo}

			got := s.Get(test.input.id)
			assert.Equal(t, test.expectedApiResponse.Code, got.Code)
			if test.expectedApiResponse.Code == StatusOK {
				assert.Equal(t, test.expectedApiResponse.Data, got.Data)
			}
		})
	}
}

func TestUserService_Update(t *testing.T) {
	db, err := prepareTestDatabase()
	if err != nil {
		t.Fatalf(err.Error())
	}

	type args struct {
		id   int
		user *models.UpdateUser
	}

	tests := []struct {
		name                string
		input               args
		expectedApiResponse *models.ApiResponse
	}{
		{
			name: "Ok",
			input: args{
				id:   1,
				user: builders.NewUpdUserBuilder().Build(),
			},
			expectedApiResponse: &models.ApiResponse{
				Code: StatusOK,
				Data: Map{},
			},
		},
		{
			name: "User is not exists",
			input: args{
				id:   3,
				user: builders.NewUpdUserBuilder().Build(),
			},
			expectedApiResponse: &models.ApiResponse{
				Code: StatusConflict,
			},
		},
		{
			name: "Repo err",
			input: args{
				id:   3,
				user: builders.NewUpdUserBuilder().WithNickname("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA").Build(),
			},
			expectedApiResponse: &models.ApiResponse{
				Code: StatusInternalServerError,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			repo := postgres.NewUserPg(db)
			s := &UserService{repo: repo}

			got := s.Update(test.input.id, test.input.user)
			assert.Equal(t, test.expectedApiResponse.Code, got.Code)
			if test.expectedApiResponse.Code == StatusOK {
				assert.Equal(t, test.expectedApiResponse.Data, got.Data)
			}
		})
	}
}
