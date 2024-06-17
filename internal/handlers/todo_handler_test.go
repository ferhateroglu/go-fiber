package handlers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/ferhateroglu/go-fiber/internal/handlers"
	"github.com/ferhateroglu/go-fiber/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockTodoService struct {
	mock.Mock
}

func (m *MockTodoService) CreateTodo(todo *models.Todo) error {
	args := m.Called(todo)
	return args.Error(0)
}

func (m *MockTodoService) GetAllTodos() ([]models.Todo, error) {
	args := m.Called()
	return args.Get(0).([]models.Todo), args.Error(1)
}

func (m *MockTodoService) GetTodoById(id string) (*models.Todo, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Todo), args.Error(1)
}

func (m *MockTodoService) UpdateTodo(id string, todo *models.Todo) error {
	args := m.Called(id, todo)
	return args.Error(0)
}

func (m *MockTodoService) DeleteTodo(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreate(t *testing.T) {
	mockService := new(MockTodoService)
	handler := handlers.NewTodoHandler(mockService)

	app := fiber.New()
	app.Post("/todos", handler.Create)

	todo := &models.Todo{Title: "Test Todo", Content: "Test Content"}
	jsonTodo, _ := json.Marshal(todo)

	mockService.On("CreateTodo", mock.AnythingOfType("*models.Todo")).Return(nil)

	req := httptest.NewRequest("POST", "/todos", bytes.NewBuffer(jsonTodo))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestGetAll(t *testing.T) {
	mockService := new(MockTodoService)
	handler := handlers.NewTodoHandler(mockService)

	app := fiber.New()
	app.Get("/todos", handler.GetAll)

	todos := []models.Todo{{Id: bson.NewObjectID(), Title: "Test Todo", Content: "Test Content"}}
	mockService.On("GetAllTodos").Return(todos, nil)

	req := httptest.NewRequest("GET", "/todos", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestGetByID(t *testing.T) {
	mockService := new(MockTodoService)
	handler := handlers.NewTodoHandler(mockService)

	app := fiber.New()
	app.Get("/todos/:id", handler.GetByID)

	todo := &models.Todo{Id: bson.NewObjectID(), Title: "Test Todo", Content: "Test Content"}
	mockService.On("GetTodoById", "1").Return(todo, nil)

	req := httptest.NewRequest("GET", "/todos/1", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {
	mockService := new(MockTodoService)
	handler := handlers.NewTodoHandler(mockService)

	app := fiber.New()
	app.Put("/todos/:id", handler.Update)

	todo := &models.Todo{Id: bson.NewObjectID(), Title: "Updated Todo", Content: "Updated Content"}
	jsonTodo, _ := json.Marshal(todo)

	mockService.On("UpdateTodo", "1", mock.AnythingOfType("*models.Todo")).Return(nil)

	req := httptest.NewRequest("PUT", "/todos/1", bytes.NewBuffer(jsonTodo))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	mockService := new(MockTodoService)
	handler := handlers.NewTodoHandler(mockService)

	app := fiber.New()
	app.Delete("/todos/:id", handler.Delete)

	mockService.On("DeleteTodo", "1").Return(nil)

	req := httptest.NewRequest("DELETE", "/todos/1", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusNoContent, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestErrorHandling(t *testing.T) {
	mockService := new(MockTodoService)
	handler := handlers.NewTodoHandler(mockService)

	app := fiber.New()
	app.Get("/todos", handler.GetAll)

	mockService.On("GetAllTodos").Return([]models.Todo{}, errors.New("database error"))

	req := httptest.NewRequest("GET", "/todos", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}
