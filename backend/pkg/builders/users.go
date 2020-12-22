package builders

import (
	"github.com/architectv/networking-course-project/backend/pkg/models"
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

func (b *UserBuilder) Build() *models.User {
	return b.User
}

func (b *UserBuilder) WithNickname(nickname string) *UserBuilder {
	b.User.Nickname = nickname
	return b
}

func (b *UserBuilder) WithEmail(email string) *UserBuilder {
	b.User.Email = email
	return b
}

func (b *UserBuilder) WithPassword(password string) *UserBuilder {
	b.User.Password = password
	return b
}

func (b *UserBuilder) WithAvatar(avatar string) *UserBuilder {
	b.User.Avatar = avatar
	return b
}

type UpdUserBuilder struct {
	UpdUser *models.UpdateUser
}

func NewUpdUserBuilder() *UpdUserBuilder {
	nickname := "Update Nickname"
	email := "Update Email"
	avatar := "Update Avatar"
	user := &models.UpdateUser{
		Nickname: &nickname,
		Email:    &email,
		Avatar:   &avatar,
	}
	return &UpdUserBuilder{UpdUser: user}
}

func (b *UpdUserBuilder) Build() *models.UpdateUser {
	return b.UpdUser
}

func (b *UpdUserBuilder) WithNickname(nickname string) *UpdUserBuilder {
	b.UpdUser.Nickname = &nickname
	return b
}

func (b *UpdUserBuilder) WithEmail(email string) *UpdUserBuilder {
	b.UpdUser.Email = &email
	return b
}

func (b *UpdUserBuilder) WithAvatar(avatar string) *UpdUserBuilder {
	b.UpdUser.Avatar = &avatar
	return b
}
