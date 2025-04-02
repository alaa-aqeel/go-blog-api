package user_controller

import (
	"github.com/alaa-aqeel/govalid/src/helpers"
	"github.com/gin-gonic/gin"
)

func (u *UserController) Index(ctx *gin.Context) {

	users, err := u.service.GetAll()
	if err != nil {

		ctx.JSON(400, helpers.DatabaseErrorResponse(err))
		return
	}

	ctx.JSON(200, helpers.DtoCollection(users))
}

func (u *UserController) Show(ctx *gin.Context) {

	user, err := u.service.Find("id", ctx.Param("id"))
	if err != nil {

		ctx.JSON(400, helpers.DatabaseErrorResponse(err))
		return
	}

	ctx.JSON(200, user.ToDTO())
}
