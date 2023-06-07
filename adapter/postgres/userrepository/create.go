package userrepository

import (
	"context"

	"github.com/EliasSantiago/api-contatos/core/domain"
	"github.com/EliasSantiago/api-contatos/core/dto"
)

func (repository repository) Create(userRequest *dto.CreateUserStore) error {
	ctx := context.Background()
	user := domain.User{}
	err := repository.db.QueryRow(
		ctx,
		"INSERT INTO users (id, name, last_name, phone, address, date_birth, cpf) VALUES ($1, $2, $3, $4, $5, $6, $7) returning *",
		userRequest.ID,
		userRequest.Name,
		userRequest.LastName,
		userRequest.Phone,
		userRequest.Address,
		userRequest.DateBirth,
		userRequest.Cpf,
	).Scan(
		&user.ID,
		&user.Name,
		&user.LastName,
		&user.Phone,
		&user.Address,
		&user.DateBirth,
		&user.Cpf,
	)
	if err != nil {
		return err
	}
	return nil
}
