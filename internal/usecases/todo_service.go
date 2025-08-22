package usecases

import (
	"errors"

	"backend-service-api/internal/models"
	"backend-service-api/internal/repositories"
)

var ErrTodoNotFound = errors.New("todo not found")

type TodoService struct {
	repo repositories.TodoRepository
}

func NewTodoService(repo repositories.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) Create(title string) (models.Todo, error) {
	if title == "" {
		return models.Todo{}, errors.New("title is required")
	}
	return s.repo.Create(models.NewTodo(title))
}

func (s *TodoService) List() ([]models.Todo, error) {
	return s.repo.List()
}

func (s *TodoService) Get(id string) (models.Todo, error) {
	t, ok, err := s.repo.GetByID(id)
	if err != nil {
		return models.Todo{}, err
	}
	if !ok {
		return models.Todo{}, ErrTodoNotFound
	}
	return t, nil
}

func (s *TodoService) Update(id string, title string, completed bool) (models.Todo, error) {
	_, err := s.Get(id)
	if err != nil {
		return models.Todo{}, err
	}
	updated := models.Todo{ID: id, Title: title, Completed: completed}
	res, ok, err := s.repo.Update(updated)
	if err != nil {
		return models.Todo{}, err
	}
	if !ok {
		return models.Todo{}, ErrTodoNotFound
	}
	return res, nil
}

func (s *TodoService) Delete(id string) error {
	ok, err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	if !ok {
		return ErrTodoNotFound
	}
	return nil
}


