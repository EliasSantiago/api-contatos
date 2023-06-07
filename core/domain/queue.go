package domain

type UserQueueService interface {
	Consumer()
}

type UserQueueUseCase interface {
	Consumer()
}

type UserQueueRepository interface {
}
