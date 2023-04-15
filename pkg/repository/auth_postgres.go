package repository

import (
	"fmt"

	"github.com/SupchickCode/simpleRestAPI"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

func (p *AuthPostgres) CreateUser(user simpleRestAPI.User) (int, error) {
	query := fmt.Sprintf(
		"INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", userTables,
	)

	row := p.db.QueryRow(query, user.Name, user.Username, user.Password)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (p *AuthPostgres) GetUser(username, password string) (simpleRestAPI.User, error) {
	var user simpleRestAPI.User
	query := fmt.Sprintf(
		"SELECT id FROM %s WHERE username=$1 and password_hash=$2 ", userTables,
	)

	if err := p.db.Get(&user, query, username, password); err != nil {
		return user, err
	}

	return user, nil
}
