package di

import (
	"github.com/EliasSantiago/api-contatos/adapter/http/userservice"
	"github.com/EliasSantiago/api-contatos/adapter/postgres"
	"github.com/EliasSantiago/api-contatos/adapter/postgres/userrepository"
	"github.com/EliasSantiago/api-contatos/core/domain"
	"github.com/EliasSantiago/api-contatos/core/domain/usecase/userusecase"
)

func ConfigUserDI(conn postgres.PoolInterface) domain.UserService {
	userRepository := userrepository.New(conn)
	userUseCase := userusecase.New(userRepository)
	userService := userservice.New(userUseCase)
	return userService
}
