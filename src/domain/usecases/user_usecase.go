package usecase

import (
	"github.com/alaa-aqeel/govalid/src/domain/interfaces"
)

type UserUsecase struct {
	Repo interfaces.RepositoryInterface
}
