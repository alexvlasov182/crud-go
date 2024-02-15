package service

import (
	"crypto/sha1"
	"fmt"

	crudgo "github.com/alexvlasov182/crud-go"
	"github.com/alexvlasov182/crud-go/pkg/repository"
)

const salt = "sakmfnasndansfgjva123123r3424356543ksfdkmdmskf"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user crudgo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
