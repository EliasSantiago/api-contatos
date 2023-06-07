package di

import (
	"github.com/EliasSantiago/api-contatos/adapter/http/userservice"
	"github.com/EliasSantiago/api-contatos/adapter/postgres"
	"github.com/EliasSantiago/api-contatos/adapter/postgres/userrepository"
	"github.com/EliasSantiago/api-contatos/core/domain"
	"github.com/EliasSantiago/api-contatos/core/domain/usecase/userusecase"
)

func ConfigUserQueueDI(conn postgres.PoolInterface) domain.UserService {
	userQueueRepository := userrepository.New(conn)
	userQueueUseCase := userusecase.New(userQueueRepository)
	userQueueService := userservice.New(userQueueUseCase)
	return userQueueService
}
