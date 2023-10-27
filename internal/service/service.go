package service

import (
	"github.com/sport-ix/auth/internal/entity"
	"github.com/sport-ix/auth/internal/repository"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
