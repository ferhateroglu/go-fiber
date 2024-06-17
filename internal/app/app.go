package app

import (
	"github.com/ferhateroglu/go-fiber/internal/routes"
	"github.com/ferhateroglu/go-fiber/pkg/di"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.uber.org/dig"
)

type App struct {
	fiber     *fiber.App
	container *dig.Container
}

func New() (*App, error) {
	app := &App{
		fiber:     fiber.New(),
		container: di.BuildContainer(),
	}

	app.fiber.Use(logger.New())

	err := app.setupRoutes()
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) setupRoutes() error {
	return a.container.Invoke(func(todoRouter *routes.TodoRouter) {
		todoRouter.SetupRoutes(a.fiber)
	})
}

func (a *App) Start(addr string) error {
	return a.fiber.Listen(addr)
}
