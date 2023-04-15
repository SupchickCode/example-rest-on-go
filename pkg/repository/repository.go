package repository

import (
	"github.com/SupchickCode/simpleRestAPI"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user simpleRestAPI.User) (int, error)
	GetUser(username, password string) (simpleRestAPI.User, error)
}

type TodoList struct {
}

type TodoItem struct {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
