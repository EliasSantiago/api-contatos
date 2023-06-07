package userusecase

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/EliasSantiago/api-contatos/core/dto"
	"github.com/EliasSantiago/api-contatos/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
)

func (usecase usecase) Consumer() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	out := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, out)
	numWorkers := 2
	for i := 1; i <= numWorkers; i++ {
		for msg := range out {
			inputDTO := &dto.CreateUserRequest{}
			if err := json.Unmarshal(msg.Body, &inputDTO); err != nil {
				panic(err)
			}

			store := &dto.CreateUserRequest{
				Name:      Decrypt(inputDTO.Name, viper.GetString("secret_key")),
				LastName:  Decrypt(inputDTO.LastName, viper.GetString("secret_key")),
				Phone:     Decrypt(inputDTO.Phone, viper.GetString("secret_key")),
				Address:   Decrypt(inputDTO.Address, viper.GetString("secret_key")),
				DateBirth: Decrypt(inputDTO.DateBirth, viper.GetString("secret_key")),
				Cpf:       Decrypt(inputDTO.Cpf, viper.GetString("secret_key")),
			}

			err = usecase.Create(store)
			if err != nil {
				// TODO: TRATAR PARA PUBLICAR EM UMA FILA DE ERROS
				panic(err)
			}
			msg.Ack(false)
			fmt.Printf("Worker %d has processed cpf: %s\n", i, inputDTO.Cpf)
		}
	}
}

func Decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

func Decrypt(text, MySecret string) string {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return ""
	}
	cipherText := Decode(text)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText)
}
