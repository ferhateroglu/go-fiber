package app

import (
	"github.com/ferhateroglu/go-fiber/internal/configs"
	"github.com/ferhateroglu/go-fiber/internal/routes"
	"github.com/ferhateroglu/go-fiber/pkg/di"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.uber.org/dig"
)

// AppDependencies embeds dig.In and defines dependencies to be injected by the DI container.
// It's used to work around dig's limitation on pointers to structs that embed dig.In.
type AppDependencies struct {
	dig.In

	Config     *configs.Config
	TodoRouter routes.TodoRouter
}

type App struct {
	fiber      *fiber.App
	container  *dig.Container
	config     *configs.Config
	todoRouter routes.TodoRouter
}

func NewApp() (*App, error) {
	container := di.BuildToDoContainer()

	var deps AppDependencies
	err := container.Invoke(func(d AppDependencies) {
		deps = d
	})

	if err != nil {
		return nil, err
	}

	app := &App{
		fiber:      fiber.New(),
		container:  container,
		config:     deps.Config,
		todoRouter: deps.TodoRouter,
	}

	app.setupMiddlewares()
	app.setupRoutes()

	return app, nil
}

func (a *App) setupMiddlewares() {
	a.fiber.Use(logger.New())
}

func (a *App) setupRoutes() {
	a.todoRouter.SetupRoutes(a.fiber)
}

func (a *App) Start() error {
	return a.fiber.Listen(":" + a.config.Server.Port)
}
