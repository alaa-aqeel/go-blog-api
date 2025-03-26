package services

import (
	"github.com/alaa-aqeel/govalid/src/domain/interfaces"
)

type UserService struct {
	repository interfaces.UserRepositoryInterface
}

func NewUserService(repository interfaces.UserRepositoryInterface) *UserService {
	return &UserService{
		repository: repository,
	}
}
