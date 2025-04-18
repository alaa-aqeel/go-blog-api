package user_controller

import (
	"github.com/alaa-aqeel/govalid/src/domain/dtos"
	"github.com/alaa-aqeel/govalid/src/helpers"
	"github.com/alaa-aqeel/govalid/src/pkgs/validator"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (u *UserController) Update(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(404, gin.H{"error": "not_found_user"})
		return
	}
	request, errs := validator.MakeValidate[dtos.UpdateUserDto](ctx.Request.Body)
	if errs.HasErrors() {
		ctx.JSON(400, helpers.ErrorResponse("invalid_data", errs))
		return
	}

	user, err := u.service.Update(id, request)
	if err != nil {
		ctx.JSON(400, gin.H{"error": helpers.DatabaseErrorResponse(err)})
		return
	}
	ctx.JSON(200, user.ToDTO())
}
