package main

import (
	"log"

	crudgo "github.com/alexvlasov182/crud-go"
	"github.com/alexvlasov182/crud-go/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(crudgo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running server %s", err.Error())
	}
}
