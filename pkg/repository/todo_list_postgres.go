package repository

import (
	"fmt"

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

func (p *TodoListPostgres) Create(userID int, list simpleRestAPI.TodoList) (int, error) {
	tx, err := p.db.Begin()

	if err != nil {
		return 0, err
	}

	var listID int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTables)
	row := p.db.QueryRow(createListQuery, list.Title, list.Desc)
	if err := row.Scan(&listID); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListsQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTables)
	_, errExec := tx.Exec(createUsersListsQuery, userID, listID)
	if errExec != nil {
		tx.Rollback()
		return 0, err
	}

	return listID, tx.Commit()
}

func (t *TodoListPostgres) GetAllLists(userID int) ([]simpleRestAPI.TodoList, error) {
	var lists []simpleRestAPI.TodoList

	query := fmt.Sprintf("SELECT t1.id, t1.title, t1.description FROM %s t1 INNER JOIN %s t2 ON t1.id = t2.list_id WHERE t2.user_id = $1 ORDER BY t1.id DESC",
		todoListsTables, usersListsTables)

	err := t.db.Select(&lists, query, userID)

	return lists, err
}

func (t *TodoListPostgres) GetListByID(listID int) (simpleRestAPI.TodoList, error) {
	var list simpleRestAPI.TodoList

	query := fmt.Sprintf("SELECT id, title, description FROM %s WHERE id = $1", todoListsTables)

	err := t.db.Get(&list, query, listID)

	return list, err
}
