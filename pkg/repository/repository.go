package repository

import (
	crudgo "github.com/alexvlasov182/crud-go"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user crudgo.User) (int, error)
}

type TodoList interface {
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
	}
}
