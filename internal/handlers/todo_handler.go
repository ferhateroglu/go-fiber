package handlers

import (
	"github.com/ferhateroglu/go-fiber/internal/models"
	"github.com/ferhateroglu/go-fiber/internal/services"
	"github.com/gofiber/fiber/v2"
)

type TodoHandler struct {
	todoService services.TodoService
}

func NewTodoHandler(todoService services.TodoService) *TodoHandler {
	return &TodoHandler{
		todoService: todoService,
	}
}

func (c *TodoHandler) CreateTodo(ctx *fiber.Ctx) error {
	todo := new(models.Todo)
	if err := ctx.BodyParser(todo); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := c.todoService.CreateTodo(todo); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(todo)
}

func (c *TodoHandler) GetAllTodos(ctx *fiber.Ctx) error {
	todos, err := c.todoService.GetAllTodos()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(todos)
}

func (c *TodoHandler) GetTodoById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	todo, err := c.todoService.GetTodoById(id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Todo not found"})
	}

	return ctx.JSON(todo)
}

func (c *TodoHandler) UpdateTodo(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	todo := new(models.Todo)
	if err := ctx.BodyParser(todo); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := c.todoService.UpdateTodo(id, todo); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(todo)
}

func (c *TodoHandler) DeleteTodo(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := c.todoService.DeleteTodo(id); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
