package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	_ "github.com/EliasSantiago/api-contatos/adapter/http/docs"
	"github.com/EliasSantiago/api-contatos/adapter/postgres"
	"github.com/EliasSantiago/api-contatos/di"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	httpSwagger "github.com/swaggo/http-swagger"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

// @title Api Contatos
// @version 1.0.0
// @contact.name Elias Fonseca
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3001
// @BasePath /
func main() {
	ctx := context.Background()
	conn := postgres.GetConnection(ctx)
	defer conn.Close()
	postgres.RunMigrations()
	userService := di.ConfigUserDI(conn)
	router := mux.NewRouter()
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	router.Handle("/user", http.HandlerFunc(userService.Publish)).Methods("POST")

	// TODO: Implementar os métodos
	// router.Handle("/user", http.HandlerFunc(userService.Get)).Methods("GET")
	// router.Handle("/user", http.HandlerFunc(userService.List)).Methods("GET")
	// router.Handle("/user", http.HandlerFunc(userService.Put)).Methods("PUT")
	// router.Handle("/user", http.HandlerFunc(userService.Delete)).Methods("DELETE")

	port := viper.GetString("server.port")
	log.Printf("LISTEN ON PORT: %v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}
