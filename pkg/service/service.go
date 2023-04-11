package service

import (
	"github.com/SupchickCode/simpleRestAPI"
	"github.com/SupchickCode/simpleRestAPI/pkg/repository"
)

type Authorization interface {
	CreateUser(user simpleRestAPI.User) (int, error)
}

type TodoList struct {
}

type TodoItem struct {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}
