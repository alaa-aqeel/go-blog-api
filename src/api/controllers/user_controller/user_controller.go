package user_controller

import "github.com/alaa-aqeel/govalid/src/domain/services"

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{
		service: service,
	}
}
