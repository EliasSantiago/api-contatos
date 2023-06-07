package main

import (
	"context"

	"github.com/EliasSantiago/api-contatos/adapter/postgres"
	"github.com/EliasSantiago/api-contatos/di"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	ctx := context.Background()
	conn := postgres.GetConnection(ctx)
	userServiceQueue := di.ConfigUserQueueDI(conn)
	userServiceQueue.Consumer()
}
