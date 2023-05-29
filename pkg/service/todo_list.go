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

func (t *TodoListService) Create(userID int, list simpleRestAPI.TodoList) (int, error) {
	return t.repo.Create(userID, list)
}

func (t *TodoListService) GetAllLists(userID int) ([]simpleRestAPI.TodoList, error) {
	return t.repo.GetAllLists(userID)
}

func (t *TodoListService) GetListByID(listID int) (simpleRestAPI.TodoList, error) {
	return t.repo.GetListByID(listID)
}
