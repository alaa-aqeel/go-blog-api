package serializer

import (
	"io"

	"github.com/alaa-aqeel/govalid/src/domain/service"
	"github.com/alaa-aqeel/govalid/src/helpers"
	"github.com/alaa-aqeel/govalid/src/pkgs/validator"
)

type UserSerializer struct{}

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

func (u UserSerializer) Create(body io.ReadCloser) (any, bool) {

	request, err := helpers.ParseBody[CreateUserRequest](body)
	if err != nil {
		return err, false
	}

	errors := validator.Validate.ValidateStruct(request)
	if errors.HasErrors() {
		return errors, false
	}

	data := helpers.StructToMap(request)
	err = service.User().Create(data)
	if err != nil {
		return err, false
	}

	return helpers.Map{
		"username": request.Username,
		"name":     request.Name,
	}, true
}
