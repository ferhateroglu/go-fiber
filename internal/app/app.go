package app

import (
	"github.com/ferhateroglu/go-fiber/internal/configs"
	"github.com/ferhateroglu/go-fiber/internal/routes"
	"github.com/ferhateroglu/go-fiber/pkg/di"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.uber.org/dig"
)

type App struct {
	fiber     *fiber.App
	container *dig.Container
	config    *configs.Config
}

func New() (*App, error) {
	container := di.BuildContainer()

	app := &App{
		fiber:     fiber.New(),
		container: container,
	}

	err := container.Invoke(func(cfg *configs.Config) {
		app.config = cfg
	})
	if err != nil {
		return nil, err
	}

	app.setupMiddlewares()
	if err := app.setupRoutes(); err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) setupMiddlewares() {
	a.fiber.Use(logger.New())
}

func (a *App) setupRoutes() error {
	return a.container.Invoke(func(todoRouter *routes.TodoRouter) {
		todoRouter.SetupRoutes(a.fiber)
	})
}

func (a *App) Start() error {
	return a.fiber.Listen(":" + a.config.Server.Port)
}
