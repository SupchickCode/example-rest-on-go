package service

import (
	"github.com/SupchickCode/simpleRestAPI"
	"github.com/SupchickCode/simpleRestAPI/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{
		repo: repo,
	}
}

func (t *TodoListService) Create(userID int, list simpleRestAPI.TodoList) {
	return t.repo.Create(userID, list)
}
