package services

import (
	"github.com/alaa-aqeel/govalid/src/domain/dtos"
	"github.com/alaa-aqeel/govalid/src/domain/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService struct {
	base BaseService
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		base: BaseService{
			Db: db,
		},
	}
}

func (u *UserService) Create(userDto *dtos.CreateUserDto) (user *models.User, err error) {

	user = &models.User{
		ID:       uuid.New(),
		Name:     userDto.Name,
		Username: userDto.Username,
		Password: userDto.Password,
	}
	err = u.base.Db.Create(&user).Error
	return
}

func (u *UserService) Find(key string, value any) (user *models.User, err error) {

	err = u.base.Db.Model(&models.User{}).Where(key, value).First(&user).Error
	return
}

func (u *UserService) GetAll() (users []models.User, err error) {

	err = u.base.Db.Model(&models.User{}).Find(&users).Error
	return
}
