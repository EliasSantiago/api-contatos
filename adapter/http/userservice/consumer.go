package userservice

func (service service) Consumer() {
	service.usecase.Consumer()
}
