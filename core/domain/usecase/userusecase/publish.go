package userusecase

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/EliasSantiago/api-contatos/core/dto"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishQueue(ch *amqp.Channel, userRequest *dto.CreateUserResponse) error {
	body, err := json.Marshal(userRequest)
	if err != nil {
		return err
	}
	err = ch.PublishWithContext(context.Background(),
		"amq.direct",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (usecase usecase) Publish(userRequest *dto.CreateUserRequest) error {
	uuid := uuid.New()
	response := &dto.CreateUserResponse{
		ID:        strings.Replace(uuid.String(), "-", "", -1),
		Name:      userRequest.Name,
		LastName:  userRequest.LastName,
		Phone:     userRequest.Phone,
		Address:   userRequest.Address,
		DateBirth: userRequest.DateBirth,
		Cpf:       userRequest.Cpf,
	}

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()
	PublishQueue(ch, response)
	return nil
}
