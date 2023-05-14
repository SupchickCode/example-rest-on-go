package repository

import (
	"github.com/SupchickCode/simpleRestAPI"
	"github.com/jmoiron/sqlx"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{
		db: db,
	}
}

func (p *AuthPostgres) Create(userID int, list simpleRestAPI.TodoList) (int, error) {
	return 0, nil
}
