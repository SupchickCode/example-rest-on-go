package service

import (
	"crypto/sha1"

	"github.com/SupchickCode/simpleRestAPI"
	"github.com/SupchickCode/simpleRestAPI/pkg/repository"
)

const salt = "n6i49mf24cnm2mf3rg0"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) CreateUser(user simpleRestAPI.User) (int, error) {
	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) (int, error) {
	hash := sha1.New()
	hash.Write([]byte(password))
}
