package main

import (
	"log"

	crudgo "github.com/alexvlasov182/crud-go"
	"github.com/alexvlasov182/crud-go/pkg/handler"
	"github.com/alexvlasov182/crud-go/pkg/repository"
	"github.com/alexvlasov182/crud-go/pkg/service"
)

func main() {

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(crudgo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running server %s", err.Error())
	}
}
