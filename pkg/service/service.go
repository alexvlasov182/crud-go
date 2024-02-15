package service

import (
	crudgo "github.com/alexvlasov182/crud-go"
	"github.com/alexvlasov182/crud-go/pkg/repository"
)

type Authorization interface {
	CreateUser(user crudgo.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
