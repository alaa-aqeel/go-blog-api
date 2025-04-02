package user_controller

import (
	"github.com/alaa-aqeel/govalid/src/domain/dtos"
	"github.com/alaa-aqeel/govalid/src/helpers"
	"github.com/alaa-aqeel/govalid/src/pkgs/validator"
	"github.com/gin-gonic/gin"
)

func (u *UserController) Store(ctx *gin.Context) {

	request, errs := validator.MakeValidate[dtos.CreateUserDto](ctx.Request.Body)
	if errs.HasErrors() {
		ctx.JSON(400, helpers.ErrorResponse("Error creating user", errs))
		return
	}

	user, err := u.service.Create(request)
	if err != nil {

		ctx.JSON(400, helpers.DatabaseErrorResponse(err))
		return
	}

	ctx.JSON(200, helpers.SuccessResponse("success creating user", user.ToDTO()))
}
