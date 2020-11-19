package services

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"time"
	"yak/backend/pkg/models"
	"yak/backend/pkg/repositories"

	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId string `json:"_id"`
}

type UserService struct {
	repo repositories.User
}

func NewUserService(repo repositories.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAll() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) Create(ctx context.Context, user *models.User) *models.ApiResponse {
	r := &models.ApiResponse{}
	if err := s.checkByNickname(ctx, user.Nickname); err == nil {
		r.Error(StatusConflict, "User already exists")
		return r
	}

	// TODO: generatePasswordHash(password)
	// user.Password = generatePasswordHash(user.Password)
	id, err := s.repo.Create(ctx, user)
	if err != nil {
		r.Error(StatusInternalServerError, err.Error())
		return r
	}

	r.Set(StatusOK, "OK", Map{userId: id})
	return r
}

func (s *UserService) checkByNickname(ctx context.Context, nickname string) error {
	_, err := s.repo.GetByNickname(ctx, nickname)
	return err
}

func (s *UserService) GenerateToken(ctx context.Context, username, password string) *models.ApiResponse {
	// TODO: generatePasswordHash(password)
	r := &models.ApiResponse{}
	user, err := s.repo.GetUser(ctx, username, password)
	if err != nil {
		r.Error(StatusConflict, err.Error())
		return r
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	encToken, err := token.SignedString([]byte(signingKey))
	if err != nil {
		r.Error(StatusConflict, err.Error())
		return r
	}

	r.Set(StatusOK, "OK", Map{"token": encToken})
	return r
}

func (s *UserService) ParseToken(ctx context.Context, accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}

	err = s.repo.FindToken(ctx, accessToken)
	if err == nil {
		return "", errors.New("Invalid token")
	}

	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *UserService) SignOut(ctx context.Context, token string) *models.ApiResponse {
	r := &models.ApiResponse{}
	err := s.repo.SignOut(ctx, token)
	if err != nil {
		r.Error(StatusInternalServerError, err.Error())
		return r
	}

	r.Set(StatusOK, "OK", nil)
	return r
}
