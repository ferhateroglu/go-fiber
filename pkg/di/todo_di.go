package di

import (
	"github.com/ferhateroglu/go-fiber/internal/configs"
	"github.com/ferhateroglu/go-fiber/internal/controllers"
	"github.com/ferhateroglu/go-fiber/internal/repositories"
	"github.com/ferhateroglu/go-fiber/internal/routes"
	"github.com/ferhateroglu/go-fiber/internal/services"
	"github.com/ferhateroglu/go-fiber/pkg/databases"
	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	container.Provide(configs.LoadConfig)
	container.Provide(databases.NewMongoDatabase)
	container.Provide(repositories.NewTodoRepository)
	container.Provide(services.NewTodoService)
	container.Provide(controllers.NewTodoController)
	container.Provide(routes.NewTodoRouter)

	return container
}
