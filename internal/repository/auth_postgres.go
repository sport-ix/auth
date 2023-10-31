package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/sport-ix/auth/internal/entity"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user entity.User) (int, error) {
	var id int
	query := ("INSERT INTO users (name, username, password_hash) values ($1, $2, $3) RETURNING id")
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (entity.User, error) {
	var user entity.User
	query := ("SELECT id FROM users WHERE username=$1 AND password_hash=$2")
	err := r.db.Get(&user, query, username, password)
	return user, err
}
