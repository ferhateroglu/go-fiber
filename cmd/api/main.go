package main

import (
	"log"

	"github.com/ferhateroglu/go-fiber/internal/app"
)

func main() {
	application, err := app.NewApp()
	if err != nil {
		log.Fatalf("Failed to create app: %v", err)
	}

	log.Fatal(application.Start())
}
