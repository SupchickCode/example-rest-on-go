package main

import (
	"log"

	"github.com/SupchickCode/simpleRestAPI"
	"github.com/SupchickCode/simpleRestAPI/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(simpleRestAPI.Server)

	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error while running http serve : %s", err.Error())
	}
}
