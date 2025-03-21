package user_controller

import (
	"github.com/alaa-aqeel/govalid/src/api/dtos"
	"github.com/alaa-aqeel/govalid/src/domain/service"
	"github.com/alaa-aqeel/govalid/src/helpers"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Store(ctx *gin.Context) {
	var request dtos.CreateUserRequest
	if err := ctx.ShouldBindBodyWithJSON(&request); err != nil {
		ctx.JSON(400, gin.H{
			"status": "error",
			"errors": helpers.ParseValidationErrors(err.(validator.ValidationErrors)),
		})
		return
	}

	err := service.User().Create(request.ToMap())
	if err != nil {
		ctx.JSON(400, gin.H{
			"status": "error",
			"errors": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": request.ToMap(),
	})
}
