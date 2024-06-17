package services_test

import (
	"errors"
	"testing"

	"github.com/ferhateroglu/go-fiber/internal/models"
	"github.com/ferhateroglu/go-fiber/internal/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockTodoRepository struct {
	mock.Mock
}

func (m *MockTodoRepository) Create(todo *models.Todo) error {
	args := m.Called(todo)
	return args.Error(0)
}

func (m *MockTodoRepository) GetAll() ([]models.Todo, error) {
	args := m.Called()
	return args.Get(0).([]models.Todo), args.Error(1)
}

func (m *MockTodoRepository) GetById(id string) (*models.Todo, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Todo), args.Error(1)
}

func (m *MockTodoRepository) Update(id string, todo *models.Todo) error {
	args := m.Called(id, todo)
	return args.Error(0)
}

func (m *MockTodoRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateTodo(t *testing.T) {
	mockRepo := new(MockTodoRepository)
	service := services.NewTodoService(mockRepo)
	todo := &models.Todo{Title: "Test Todo"}

	t.Run("Repository Error", func(t *testing.T) {
		mockRepo.On("Create", mock.AnythingOfType("*models.Todo")).Return(errors.New("db error"))

		err := service.CreateTodo(todo)

		assert.Error(t, err)
		assert.Equal(t, "failed to create todo", err.Error())

		mockRepo.AssertExpectations(t)
	})

	t.Run("Success", func(t *testing.T) {
		mockRepo.ExpectedCalls = nil
		mockRepo.Calls = nil

		mockRepo.On("Create", mock.AnythingOfType("*models.Todo")).Return(nil)

		err := service.CreateTodo(todo)

		assert.NoError(t, err)

		mockRepo.AssertExpectations(t)
	})
}
func TestGetAllTodos(t *testing.T) {
	mockRepo := new(MockTodoRepository)
	service := services.NewTodoService(mockRepo)

	t.Run("Success", func(t *testing.T) {

		expectedTodos := []models.Todo{
			{Id: bson.NewObjectID(), Title: "Test Todo 1"},
			{Id: bson.NewObjectID(), Title: "Test Todo 2"},
		}
		mockRepo.On("GetAll").Return(expectedTodos, nil)

		todos, err := service.GetAllTodos()

		assert.NoError(t, err)
		assert.Equal(t, expectedTodos, todos)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Repository Error", func(t *testing.T) {
		mockRepo.ExpectedCalls = nil
		mockRepo.Calls = nil

		mockRepo.On("GetAll").Return([]models.Todo{}, errors.New("db error"))

		todos, err := service.GetAllTodos()

		assert.Error(t, err)
		assert.Nil(t, todos)
		assert.Equal(t, "failed to fetch todos", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestGetTodoById(t *testing.T) {
	mockRepo := new(MockTodoRepository)
	service := services.NewTodoService(mockRepo)

	t.Run("Success", func(t *testing.T) {

		id := bson.NewObjectID().Hex()
		expectedTodo := &models.Todo{Id: bson.NewObjectID(), Title: "Test Todo"}
		mockRepo.On("GetById", id).Return(expectedTodo, nil)

		todo, err := service.GetTodoById(id)

		assert.NoError(t, err)
		assert.Equal(t, expectedTodo, todo)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Not Found", func(t *testing.T) {

		id := bson.NewObjectID().Hex()
		mockRepo.On("GetById", id).Return((*models.Todo)(nil), errors.New("not found"))

		todo, err := service.GetTodoById(id)

		assert.Error(t, err)
		assert.Nil(t, todo)
		assert.Equal(t, "failed to fetch todo", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestUpdateTodo(t *testing.T) {
	mockRepo := new(MockTodoRepository)
	service := services.NewTodoService(mockRepo)

	t.Run("Success", func(t *testing.T) {
		id := bson.NewObjectID().Hex()
		todo := &models.Todo{Title: "Updated Todo"}
		mockRepo.On("Update", id, todo).Return(nil)

		err := service.UpdateTodo(id, todo)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Empty Title", func(t *testing.T) {
		id := bson.NewObjectID().Hex()
		todo := &models.Todo{Title: ""}

		err := service.UpdateTodo(id, todo)

		assert.Error(t, err)
		assert.Equal(t, "title is required", err.Error())
	})

	t.Run("Repository Error", func(t *testing.T) {
		id := bson.NewObjectID().Hex()
		todo := &models.Todo{Title: "Updated Todo"}
		mockRepo.On("Update", id, todo).Return(errors.New("db error"))

		err := service.UpdateTodo(id, todo)

		assert.Error(t, err)
		assert.Equal(t, "failed to update todo", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestDeleteTodo(t *testing.T) {
	mockRepo := new(MockTodoRepository)
	service := services.NewTodoService(mockRepo)

	t.Run("Success", func(t *testing.T) {
		id := bson.NewObjectID().Hex()
		mockRepo.On("Delete", id).Return(nil)

		err := service.DeleteTodo(id)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Repository Error", func(t *testing.T) {
		id := bson.NewObjectID().Hex()
		mockRepo.On("Delete", id).Return(errors.New("db error"))

		err := service.DeleteTodo(id)

		assert.Error(t, err)
		assert.Equal(t, "failed to delete todo", err.Error())
		mockRepo.AssertExpectations(t)
	})
}
