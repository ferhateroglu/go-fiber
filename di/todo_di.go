package di

import (
	"github.com/ferhateroglu/go-fiber/controllers"
	"github.com/ferhateroglu/go-fiber/databases"
	"github.com/ferhateroglu/go-fiber/repositories"
	"github.com/ferhateroglu/go-fiber/routes"
	"github.com/ferhateroglu/go-fiber/services"
	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	container.Provide(databases.NewMongoDatabase)
	container.Provide(repositories.NewMongoTodoRepository)
	container.Provide(services.NewTodoService)
	container.Provide(controllers.NewTodoController)
	container.Provide(routes.NewTodoRouter)

	return container
}
