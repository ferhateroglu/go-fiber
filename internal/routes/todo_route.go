package routes

import (
	"github.com/ferhateroglu/go-fiber/internal/handlers"
	"github.com/ferhateroglu/go-fiber/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

type TodoRouter interface {
	SetupRoutes(app *fiber.App)
}

type todoRouter struct {
	todoHandler handlers.TodoHandler
}

func NewTodoRouter(todoHandler handlers.TodoHandler) TodoRouter {
	return &todoRouter{
		todoHandler: todoHandler,
	}
}

func (tr *todoRouter) SetupRoutes(app *fiber.App) {
	todoGroup := app.Group("/api/todos")

	todoGroup.Get("/", tr.todoHandler.GetAll)
	todoGroup.Post("/", middlewares.ValidateRequest(middlewares.TodoRequest{}), tr.todoHandler.Create)
	todoGroup.Get("/:id", tr.todoHandler.GetByID)
	todoGroup.Put("/:id", middlewares.ValidateRequest(middlewares.UpdateTodoRequest{}), tr.todoHandler.Update)
	todoGroup.Delete("/:id", middlewares.ValidateRequest(middlewares.DeleteTodoRequest{}), tr.todoHandler.Delete)
}
