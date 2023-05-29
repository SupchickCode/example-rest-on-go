package repository

import (
	"github.com/SupchickCode/simpleRestAPI"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user simpleRestAPI.User) (int, error)
	GetUser(username, password string) (simpleRestAPI.User, error)
}

type TodoList interface {
	Create(userID int, list simpleRestAPI.TodoList) (int, error)
	GetAllLists(userID int) ([]simpleRestAPI.TodoList, error)
	GetListByID(listID int) (simpleRestAPI.TodoList, error)
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
	}
}
