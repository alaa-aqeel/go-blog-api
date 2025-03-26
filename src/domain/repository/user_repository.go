package repository

import (
	"github.com/alaa-aqeel/govalid/src/domain/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) Create(user *models.User) error {
	return u.db.Create(user).Error
}

func (u *UserRepository) Update(user *models.User) error {

	return nil
}

func (u *UserRepository) Delete(id string) error {

	return nil
}

func (u *UserRepository) Find(key string, value any) (*models.User, error) {

	return nil, nil
}

func (u *UserRepository) GetAll() ([]models.User, error) {

	return nil, nil
}
