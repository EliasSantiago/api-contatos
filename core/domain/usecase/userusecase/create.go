package userusecase

import (
	"strings"

	"github.com/EliasSantiago/api-contatos/core/dto"
	"github.com/google/uuid"
)

func (usecase usecase) Create(userRequest *dto.CreateUserRequest) error {
	uuid := uuid.New()
	store := &dto.CreateUserStore{
		ID:        strings.Replace(uuid.String(), "-", "", -1),
		Name:      userRequest.Name,
		LastName:  userRequest.LastName,
		Phone:     userRequest.Phone,
		Address:   userRequest.Address,
		DateBirth: userRequest.DateBirth,
		Cpf:       userRequest.Cpf,
	}
	err := usecase.repository.Create(store)
	if err != nil {
		return err
	}
	return nil
}
