package userusecase

import "github.com/EliasSantiago/api-contatos/core/domain"

type usecase struct {
	repository domain.UserRepository
}

func New(repository domain.UserRepository) domain.UserUseCase {
	return &usecase{
		repository: repository,
	}
}
