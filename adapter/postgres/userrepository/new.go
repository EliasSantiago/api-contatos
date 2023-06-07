package userrepository

import (
	"github.com/EliasSantiago/api-contatos/adapter/postgres"
	"github.com/EliasSantiago/api-contatos/core/domain"
)

type repository struct {
	db postgres.PoolInterface
}

func New(db postgres.PoolInterface) domain.UserRepository {
	return &repository{
		db: db,
	}
}
