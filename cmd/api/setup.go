package main

import (
	"github.com/ferhateroglu/go-fiber/internal/routes"
	"github.com/ferhateroglu/go-fiber/pkg/di"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.uber.org/dig"
)

func SetupApp() (*fiber.App, error) {
	app := fiber.New()
	app.Use(logger.New())

	container := di.BuildContainer()

	err := setupRoutes(app, container)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func setupRoutes(app *fiber.App, container *dig.Container) error {
	return container.Invoke(func(todoRouter *routes.TodoRouter) {
		todoRouter.SetupRoutes(app)
	})
}
