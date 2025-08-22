package memory

import (
	"sync"

	"backend-service-api/internal/models"
	"backend-service-api/internal/repositories"
)

// Ensure implementation satisfies interface at compile time
var _ repositories.TodoRepository = (*TodoMemoryRepository)(nil)

type TodoMemoryRepository struct {
	mu    sync.RWMutex
	items map[string]models.Todo
}

func NewTodoMemoryRepository() *TodoMemoryRepository {
	return &TodoMemoryRepository{items: make(map[string]models.Todo)}
}

func (r *TodoMemoryRepository) Create(todo models.Todo) (models.Todo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.items[todo.ID] = todo
	return todo, nil
}

func (r *TodoMemoryRepository) List() ([]models.Todo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	result := make([]models.Todo, 0, len(r.items))
	for _, t := range r.items {
		result = append(result, t)
	}
	return result, nil
}

func (r *TodoMemoryRepository) GetByID(id string) (models.Todo, bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	t, ok := r.items[id]
	return t, ok, nil
}

func (r *TodoMemoryRepository) Update(todo models.Todo) (models.Todo, bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.items[todo.ID]; !ok {
		return models.Todo{}, false, nil
	}
	r.items[todo.ID] = todo
	return todo, true, nil
}

func (r *TodoMemoryRepository) Delete(id string) (bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.items[id]; !ok {
		return false, nil
	}
	delete(r.items, id)
	return true, nil
}


