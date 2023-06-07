package userservice

import (
	"github.com/EliasSantiago/api-contatos/core/domain"
)

type service struct {
	usecase domain.UserUseCase
}

func New(usecase domain.UserUseCase) domain.UserService {
	return &service{
		usecase: usecase,
	}
}
