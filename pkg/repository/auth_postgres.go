package repository

import (
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
	return 1, nil
}
