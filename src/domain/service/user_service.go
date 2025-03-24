package service

import (
	"github.com/alaa-aqeel/govalid/src/domain/models"
	"github.com/alaa-aqeel/govalid/src/helpers"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func (s *UserService) hashPassword(data map[string]any) error {
	if _, ok := data["password"]; ok {
		password, err := helpers.Hash().Make(data["password"].(string))
		if err != nil {
			return err
		}
		data["password"] = password
	}

	return nil
}

func (s *UserService) Create(data map[string]any) error {

	s.hashPassword(data)
	data["id"] = uuid.New()
	tx := s.db.Model(&models.User{}).Create(&data)

	return tx.Error
}

func (s *UserService) Update(id string, data map[string]any) error {
	s.hashPassword(data)
	tx := s.db.Model(&models.User{}).Where("id = ?", id).Updates(&data)

	return tx.Error
}

func (s *UserService) GetAll() []models.User {
	var users []models.User
	s.db.Model(&models.User{}).Find(&users)

	return users
}

func (s *UserService) Filter(data map[string]any) []models.User {
	var users []models.User
	s.db.Model(&models.User{}).Where(data).Find(&users)

	return users
}

func (s *UserService) Find(id string) (models.User, error) {
	var user models.User
	tx := s.db.Model(&models.User{}).Where("id = ?", id).First(&user)

	return user, tx.Error
}
