package interfaces

import (
	"github.com/alaa-aqeel/govalid/src/domain/models"
)

type UserRepositoryInterface interface {
	Create(user *models.User) error
	Delete(id string) error
	Update(user *models.User) error
	Find(key string, value any) (*models.User, error)
	GetAll() ([]models.User, error)
}
