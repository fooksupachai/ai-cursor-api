package repositories

import "backend-service-api/internal/models"

type UserRepository interface {
	GetByEmail(email string) (models.User, bool, error)
	GetByID(id string) (models.User, bool, error)
	Create(user models.User) (models.User, error)
	Update(user models.User) (models.User, error)
}


