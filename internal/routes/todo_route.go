package routes

import (
	"github.com/ferhateroglu/go-fiber/internal/controllers"
	"github.com/ferhateroglu/go-fiber/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

type TodoRouter struct {
	todoController *controllers.TodoController
}

func NewTodoRouter(todoController *controllers.TodoController) *TodoRouter {
	return &TodoRouter{
		todoController: todoController,
	}
}

func (tr *TodoRouter) SetupRoutes(app *fiber.App) {
	todoGroup := app.Group("/api/todos")

	todoGroup.Get("/", tr.todoController.GetAllTodos)
	todoGroup.Post("/", middlewares.ValidateRequest(middlewares.TodoRequest{}), tr.todoController.CreateTodo)
	todoGroup.Get("/:id", middlewares.ValidateRequest(middlewares.GetTodoRequest{}), tr.todoController.GetTodoById)
	todoGroup.Put("/:id", middlewares.ValidateRequest(middlewares.UpdateTodoRequest{}), tr.todoController.UpdateTodo)
	todoGroup.Delete("/:id", middlewares.ValidateRequest(middlewares.DeleteTodoRequest{}), tr.todoController.DeleteTodo)
}
