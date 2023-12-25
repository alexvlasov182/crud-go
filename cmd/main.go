package main

import (
	"log"

	crudgo "github.com/alexvlasov182/crud-go"
	"github.com/alexvlasov182/crud-go/pkg/handler"
	"github.com/alexvlasov182/crud-go/pkg/repository"
	"github.com/alexvlasov182/crud-go/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(crudgo.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running server %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
