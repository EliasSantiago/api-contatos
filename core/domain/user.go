package domain

import (
	"net/http"

	"github.com/EliasSantiago/api-contatos/core/dto"
)

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	DateBirth string `json:"date_birth"`
	Cpf       string `json:"cpf"`
}

type UserService interface {
	Publish(response http.ResponseWriter, request *http.Request)
	Consumer()
}

type UserUseCase interface {
	Publish(userRequest *dto.CreateUserRequest) error
	Consumer()
}

type UserRepository interface {
	Create(userRequest *dto.CreateUserStore) error
}
