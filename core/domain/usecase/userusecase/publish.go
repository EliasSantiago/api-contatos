package userusecase

import (
	"context"
	"encoding/json"

	"github.com/EliasSantiago/api-contatos/core/dto"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"

	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func PublishQueue(ch *amqp.Channel, userRequest *dto.CreateUserRequest) error {
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
	user := &dto.CreateUserRequest{
		Name:      Encrypt(userRequest.Name, viper.GetString("secret_key")),
		LastName:  Encrypt(userRequest.LastName, viper.GetString("secret_key")),
		Phone:     Encrypt(userRequest.Phone, viper.GetString("secret_key")),
		Address:   Encrypt(userRequest.Address, viper.GetString("secret_key")),
		DateBirth: Encrypt(userRequest.DateBirth, viper.GetString("secret_key")),
		Cpf:       Encrypt(userRequest.Cpf, viper.GetString("secret_key")),
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
	PublishQueue(ch, user)
	return nil
}

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Encrypt(text, MySecret string) string {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return ""
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return Encode(cipherText)
}
