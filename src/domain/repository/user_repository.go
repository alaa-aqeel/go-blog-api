package repository

import (
	"errors"

	"github.com/alaa-aqeel/govalid/src/domain/models"
	"github.com/alaa-aqeel/govalid/src/helpers"
	"github.com/google/uuid"
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

// HashPassword hashes the given password using bcrypt.
func hashPassword(data map[string]any) error {
	password, ok := data["password"].(string)
	if !ok {
		return errors.New("password is required")
	}
	hash, err := helpers.Hash().Make(password)
	if err != nil {
		return err
	}
	data["password"] = hash
	return nil
}

func (r *UserRepository) Create(data map[string]any) error {
	data["id"] = uuid.New()
	hashPassword(data)
	return r.db.Model(&models.User{}).Create(data).Error
}

func (r *UserRepository) Update(id string, data map[string]any) error {
	hashPassword(data)
	return r.db.Model(&models.User{}).Where("id = ?", id).Updates(data).Error
}

func (r *UserRepository) Delete(id string) error {
	return r.db.Delete(&models.User{}, "id = ?", id).Error
}

func (r *UserRepository) Find(key string, value any) (*models.User, error) {
	var user models.User
	err := r.db.Where(key+" = ?", value).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
