package middlewares

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type TodoRequest struct {
	Title   string `json:"title" validate:"required,min=3,max=100"`
	Content string `json:"content" validate:"max=500"`
}

type GetTodoRequest struct {
	Id int `json:"id" validate:"required,min=1"`
}

type UpdateTodoRequest struct {
	Id      int    `json:"id" validate:"required,min=1"`
	Title   string `json:"title" validate:"required,min=3,max=100"`
	Content string `json:"content" validate:"max=500"`
}

type DeleteTodoRequest struct {
	Id int `json:"id" validate:"required,min=1"`
}

func ValidateRequest(requestType interface{}) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		req := reflect.New(reflect.TypeOf(requestType)).Interface()
		if err := ctx.BodyParser(req); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Json parsing error",
			})
		}

		if err := validate.Struct(req); err != nil {
			errors := err.(validator.ValidationErrors)
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": formatValidationErrors(errors),
			})
		}

		ctx.Locals("validatedRequest", req)
		return ctx.Next()
	}
}

func formatValidationErrors(errors validator.ValidationErrors) map[string]string {
	errorMap := make(map[string]string)

	for _, err := range errors {
		switch err.Tag() {
		case "required":
			errorMap[err.Field()] = "Id is required"
		case "min":
			errorMap[err.Field()] = "Must be at least " + err.Param() + " characters long"
		case "max":
			errorMap[err.Field()] = "Must be at most " + err.Param() + " characters long"
		default:
			errorMap[err.Field()] = "Invalid value"
		}
	}

	return errorMap
}
