package main

import (
	"log"

	"github.com/SupchickCode/simpleRestAPI"
	"github.com/SupchickCode/simpleRestAPI/pkg/handler"
	"github.com/SupchickCode/simpleRestAPI/pkg/repository"
	"github.com/SupchickCode/simpleRestAPI/pkg/service"
)

func main() {
	repo := repository.NewRepository()
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(simpleRestAPI.Server)

	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error while running http serve : %s", err.Error())
	}
}
