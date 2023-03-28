package main

import (
	"log"

	"github.com/SupchickCode/simpleRestAPI"
)

func main() {
	srv := new(simpleRestAPI.Server)

	if err := srv.Run(":8000"); err != nil {
		log.Fatalf("Error while running http serve : %s", err.Error())
	}
}
