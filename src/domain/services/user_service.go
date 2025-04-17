package services

import (
	"github.com/alaa-aqeel/govalid/src/domain/dtos"
	"github.com/alaa-aqeel/govalid/src/domain/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService struct {
	base BaseService[models.User]
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		base: BaseService[models.User]{
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
	err = u.base.Query().Create(&user).Error
	return
}

func (u *UserService) Find(key string, value any) (user *models.User, err error) {

	err = u.base.Query().Where(key, value).First(&user).Error
	return
}

func (u *UserService) GetAll() (users []models.User, err error) {

	err = u.base.Query().Find(&users).Error
	return
}

func (u *UserService) Update(id uuid.UUID, userDto *dtos.UpdateUserDto) (*models.User, error) {

	var user *models.User
	if err := u.base.Query().First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	user.Name = userDto.Name
	user.Username = userDto.Username

	if userDto.Password != "" {
		user.Password = userDto.Password
	}

	err := u.base.Query().Where("id = ?", id).Updates(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
