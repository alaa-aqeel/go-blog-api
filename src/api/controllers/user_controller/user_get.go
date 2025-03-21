package user_controller

import (
	"github.com/alaa-aqeel/govalid/src/domain/service"
	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {

	users := service.User().GetAll()

	ctx.JSON(200, gin.H{
		"data": users,
	})
}

func Show(ctx *gin.Context) {

	user, err := service.User().Find(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{
			"status": "error",
			"errors": err.Error(),
		})
		return
	}

	ctx.JSON(200, user)
}
