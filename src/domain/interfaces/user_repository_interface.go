package interfaces

import (
	"github.com/alaa-aqeel/govalid/src/domain/models"
)

type RepositoryInterface interface {
	Create(data map[string]any) error
	Update(id string, data map[string]any) error
	Delete(id string) error
}

type UserRepositoryInterface interface {
	RepositoryInterface
	Find(key string, value any) (*models.User, error)
	GetAll() ([]models.User, error)
}
