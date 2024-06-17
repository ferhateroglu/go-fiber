package main

import (
	"log"

	"github.com/ferhateroglu/go-fiber/di"
	"github.com/ferhateroglu/go-fiber/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	container := di.BuildContainer()

	err := container.Invoke(func(todoRouter *routes.TodoRouter) {
		todoRouter.SetupRoutes(app)
	})

	if err != nil {
		log.Fatalf("Failed to invoke container: %v", err)
	}

	log.Fatal(app.Listen(":3000"))
}
