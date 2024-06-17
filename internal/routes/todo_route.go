package routes

import (
	"github.com/ferhateroglu/go-fiber/internal/handlers"
	"github.com/ferhateroglu/go-fiber/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

type TodoRouter struct {
	todoHandler *handlers.TodoHandler
}

func NewTodoRouter(todoHandler *handlers.TodoHandler) *TodoRouter {
	return &TodoRouter{
		todoHandler: todoHandler,
	}
}

func (tr *TodoRouter) SetupRoutes(app *fiber.App) {
	todoGroup := app.Group("/api/todos")

	todoGroup.Get("/", tr.todoHandler.GetAllTodos)
	todoGroup.Post("/", middlewares.ValidateRequest(middlewares.TodoRequest{}), tr.todoHandler.CreateTodo)
	todoGroup.Get("/:id", middlewares.ValidateRequest(middlewares.GetTodoRequest{}), tr.todoHandler.GetTodoById)
	todoGroup.Put("/:id", middlewares.ValidateRequest(middlewares.UpdateTodoRequest{}), tr.todoHandler.UpdateTodo)
	todoGroup.Delete("/:id", middlewares.ValidateRequest(middlewares.DeleteTodoRequest{}), tr.todoHandler.DeleteTodo)
}
