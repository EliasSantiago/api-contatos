package dto

import (
	"encoding/json"
	"io"
)

type CreateUserRequest struct {
	Name      string `json:"name" validate:"required,name"`
	LastName  string `json:"last_name" validate:"required,last_name"`
	Phone     string `json:"phone" validate:"required,phone"`
	Address   string `json:"address" validate:"required,address"`
	DateBirth string `json:"date_birth" validate:"required,date_birth"`
	Cpf       string `json:"cpf"`
}

type CreateUserResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	DateBirth string `json:"date_birth"`
	Cpf       string `json:"cpf"`
}

type CreateUserStore struct {
	ID        string `db:"id"`
	Name      string `db:"name"`
	LastName  string `db:"last_name"`
	Phone     string `db:"phone"`
	Address   string `db:"address"`
	DateBirth string `db:"date_birth"`
	Cpf       string `db:"cpf"`
}

func FromJSONCreateUserRequest(body io.Reader) (*CreateUserRequest, error) {
	createUserRequest := CreateUserRequest{}
	if err := json.NewDecoder(body).Decode(&createUserRequest); err != nil {
		return nil, err
	}
	return &createUserRequest, nil
}
