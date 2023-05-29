package service

import (
	"github.com/SupchickCode/simpleRestAPI"
	"github.com/SupchickCode/simpleRestAPI/pkg/repository"
)

type Authorization interface {
	CreateUser(user simpleRestAPI.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userID int, list simpleRestAPI.TodoList) (int, error)
	GetAllLists(userID int) ([]simpleRestAPI.TodoList, error)
	GetListByID(listID int) (simpleRestAPI.TodoList, error)
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		TodoList:      NewTodoListService(repo.TodoList),
	}
}
