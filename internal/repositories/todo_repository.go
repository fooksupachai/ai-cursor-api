package repositories

import "backend-service-api/internal/models"

// TodoRepository defines the storage contract for Todo entities.
type TodoRepository interface {
	Create(todo models.Todo) (models.Todo, error)
	List() ([]models.Todo, error)
	GetByID(id string) (models.Todo, bool, error)
	Update(todo models.Todo) (models.Todo, bool, error)
	Delete(id string) (bool, error)
}


