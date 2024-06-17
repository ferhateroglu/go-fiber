package services

import (
	"errors"

	"github.com/ferhateroglu/go-fiber/internal/models"
	"github.com/ferhateroglu/go-fiber/internal/repositories"
)

type TodoService interface {
	CreateTodo(todo *models.Todo) error
	GetAllTodos() ([]models.Todo, error)
	GetTodoById(id string) (*models.Todo, error)
	UpdateTodo(id string, todo *models.Todo) error
	DeleteTodo(id string) error
}

type todoService struct {
	repo repositories.TodoRepository
}

func NewTodoService(repo repositories.TodoRepository) TodoService {
	return &todoService{
		repo: repo,
	}
}

func (s *todoService) CreateTodo(todo *models.Todo) error {
	if todo.Title == "" {
		return errors.New("title is required")
	}

	err := s.repo.Create(todo)
	if err != nil {
		return errors.New("failed to create todo")
	}

	return nil
}

func (s *todoService) GetAllTodos() ([]models.Todo, error) {
	todos, err := s.repo.GetAll()
	if err != nil {
		return nil, errors.New("failed to fetch todos")
	}

	return todos, nil
}

func (s *todoService) GetTodoById(id string) (*models.Todo, error) {
	todo, err := s.repo.GetById(id)
	if err != nil {
		return nil, errors.New("failed to fetch todo")
	}

	return todo, nil
}

func (s *todoService) UpdateTodo(id string, todo *models.Todo) error {
	if todo.Title == "" {
		return errors.New("title is required")
	}

	err := s.repo.Update(id, todo)
	if err != nil {
		return errors.New("failed to update todo")
	}

	return nil
}

func (s *todoService) DeleteTodo(id string) error {
	err := s.repo.Delete(id)
	if err != nil {
		return errors.New("failed to delete todo")
	}

	return nil
}
