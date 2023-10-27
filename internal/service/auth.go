package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/sport-ix/auth/internal/entity"
	"github.com/sport-ix/auth/internal/repository"
)

const salt = "02m1d9u431d094uf12mn094fu8"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user entity.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
