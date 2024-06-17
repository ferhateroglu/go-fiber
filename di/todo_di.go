package di

import (
	"github.com/ferhateroglu/go-fiber/controllers"
	"github.com/ferhateroglu/go-fiber/database"
	"github.com/ferhateroglu/go-fiber/repositories"
	"github.com/ferhateroglu/go-fiber/routes"
	"github.com/ferhateroglu/go-fiber/services"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	container.Provide(func(md *database.MongoDatabase) *mongo.Database {
		return md.GetDatabase()
	})
	container.Provide(database.NewMongoDatabase)
	container.Provide(repositories.NewMongoTodoRepository)
	container.Provide(services.NewTodoService)
	container.Provide(controllers.NewTodoController)
	container.Provide(routes.NewTodoRouter)

	return container
}
