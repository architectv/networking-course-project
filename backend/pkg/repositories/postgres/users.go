package postgres

import (
	"fmt"
	"yak/backend/pkg/models"

	"github.com/jmoiron/sqlx"
)

type UserPg struct {
	db *sqlx.DB
}

func NewUserPg(db *sqlx.DB) *UserPg {
	return &UserPg{db: db}
}

func (r *UserPg) GetAll() ([]*models.User, error) {
	var users []*models.User
	query := fmt.Sprintf(`SELECT * FROM %s`, usersTable)
	if err := r.db.Select(&users, query); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserPg) Get(nickname, password string) (*models.User, error) {
	user := &models.User{}
	query := fmt.Sprintf(
		`SELECT * FROM %s WHERE nickname = $1 AND password = $2`, usersTable)
	err := r.db.Get(user, query, nickname, password)

	return user, err
}

func (r *UserPg) Create(user *models.User) (int, error) {
	var id int
	query := fmt.Sprintf(
		`INSERT INTO %s (nickname, email, password, avatar)
		VALUES ($1, $2, $3, $4) RETURNING id`, usersTable)

	row := r.db.QueryRow(query, user.Nickname, user.Email, user.Password, user.Avatar)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserPg) GetByNickname(nickname string) (*models.User, error) {
	user := &models.User{}
	query := fmt.Sprintf(`SELECT * FROM %s WHERE nickname = $1`, usersTable)
	err := r.db.Get(user, query, nickname)

	return user, err
}
