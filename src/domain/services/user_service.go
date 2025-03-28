package services

import (
	"github.com/alaa-aqeel/govalid/src/domain/interfaces"
	"github.com/alaa-aqeel/govalid/src/domain/models"
)

type UserService struct {
	repository interfaces.UserRepositoryInterface
}

func NewUserService(repository interfaces.UserRepositoryInterface) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) Create(data map[string]any) error {

	return s.repository.Create(data)
}

func (s *UserService) Update(id string, data map[string]any) error {
	return s.repository.Update(id, data)
}

func (s *UserService) GetAll() ([]models.User, error) {
	return s.repository.GetAll()
}

func (s *UserService) Find(key string, value any) (*models.User, error) {
	return s.repository.Find(key, value)
}
