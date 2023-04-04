package service

import "github.com/SupchickCode/simpleRestAPI/pkg/repository"

type Authorization struct {
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
	return &Service{}
}
