package user_controller

import (
	"github.com/alaa-aqeel/govalid/src/api/serializer"
	"github.com/gin-gonic/gin"
)

func Store(ctx *gin.Context) {

	data, isOk := (serializer.UserSerializer{}).Create(ctx.Request.Body)
	if !isOk {
		ctx.JSON(400, gin.H{
			"status": "error",
			"errors": data,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": data,
	})
}
