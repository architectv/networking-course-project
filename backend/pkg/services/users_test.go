package services

import (
	"errors"
	"testing"
	"yak/backend/pkg/models"

	mock_repositories "yak/backend/pkg/repositories/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type UserBuilder struct {
	User *models.User
}

func NewUserBuilder() *UserBuilder {
	user := &models.User{
		Nickname: "Dafault Nickname",
		Email:    "Default Email",
		Password: "Default Password",
		Avatar:   "Default Avatar",
	}
	return &UserBuilder{User: user}
}

func (b *UserBuilder) build() *models.User {
	return b.User
}

func (b *UserBuilder) withNickname(nickname string) *UserBuilder {
	b.User.Nickname = nickname
	return b
}

func (b *UserBuilder) withEmail(email string) *UserBuilder {
	b.User.Email = email
	return b
}

func (b *UserBuilder) withPassword(password string) *UserBuilder {
	b.User.Password = password
	return b
}

func (b *UserBuilder) withAvatar(avatar string) *UserBuilder {
	b.User.Avatar = avatar
	return b
}

func TestUserService_Create(t *testing.T) {
	type args struct {
		user *models.User
	}
	type mockCheck func(r *mock_repositories.MockUser, nickname string)
	type mockCreate func(r *mock_repositories.MockUser, user *models.User)

	tests := []struct {
		name                string
		input               args
		mockCheck           mockCheck
		mockCreate          mockCreate
		expectedApiResponse *models.ApiResponse
	}{
		{
			name: "Ok",
			input: args{
				user: NewUserBuilder().withNickname("User Builder").build(),
			},
			mockCheck: func(r *mock_repositories.MockUser, nickname string) {
				r.EXPECT().GetByNickname(nickname).Return(nil, errors.New("new nickname"))
			},
			mockCreate: func(r *mock_repositories.MockUser, user *models.User) {
				r.EXPECT().Create(user).Return(1, nil)
			},
			expectedApiResponse: &models.ApiResponse{
				Code: StatusOK,
				Data: Map{"uid": 1},
			},
		},
		{
			name: "Already Exists",
			input: args{
				user: NewUserBuilder().withNickname("User Builder").build(),
			},
			mockCheck: func(r *mock_repositories.MockUser, nickname string) {
				r.EXPECT().GetByNickname(nickname).Return(nil, nil)
			},
			mockCreate: func(r *mock_repositories.MockUser, user *models.User) {},
			expectedApiResponse: &models.ApiResponse{
				Code: StatusConflict,
			},
		},
		{
			name: "Repo Error",
			input: args{
				user: NewUserBuilder().withNickname("User Builder").build(),
			},
			mockCheck: func(r *mock_repositories.MockUser, nickname string) {
				r.EXPECT().GetByNickname(nickname).Return(nil, errors.New("new nickname"))
			},
			mockCreate: func(r *mock_repositories.MockUser, user *models.User) {
				r.EXPECT().Create(user).Return(0, errors.New("repo error"))
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

			repo := mock_repositories.NewMockUser(c)
			test.mockCheck(repo, test.input.user.Nickname)
			test.mockCreate(repo, test.input.user)
			s := &UserService{repo: repo}

			got := s.Create(test.input.user)
			assert.Equal(t, test.expectedApiResponse.Code, got.Code)
			if test.expectedApiResponse.Code == StatusOK {
				assert.Equal(t, test.expectedApiResponse.Data, got.Data)
			}
		})
	}
}
