package repositories

import "backend-service-api/internal/models"

// TodoRepository defines the storage contract for Todo entities.
type TodoRepository interface {
\tCreate(todo models.Todo) (models.Todo, error)
\tList() ([]models.Todo, error)
\tGetByID(id string) (models.Todo, bool, error)
\tUpdate(todo models.Todo) (models.Todo, bool, error)
\tDelete(id string) (bool, error)
}


