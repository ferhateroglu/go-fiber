package services

import (
	"errors"

	"github.com/ferhateroglu/go-fiber/models"
	"github.com/ferhateroglu/go-fiber/repositories"
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
	return s.repo.Create(todo)
}

func (s *todoService) GetAllTodos() ([]models.Todo, error) {
	return s.repo.GetAll()
}

func (s *todoService) GetTodoById(id string) (*models.Todo, error) {
	return s.repo.GetById(id)
}

func (s *todoService) UpdateTodo(id string, todo *models.Todo) error {
	if todo.Title == "" {
		return errors.New("title is required")
	}
	return s.repo.Update(id, todo)
}

func (s *todoService) DeleteTodo(id string) error {
	return s.repo.Delete(id)
}
