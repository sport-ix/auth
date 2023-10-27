package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/sport-ix/auth/internal/entity"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
