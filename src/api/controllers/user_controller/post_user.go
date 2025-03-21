package user_controller

import "github.com/gin-gonic/gin"

func PostUser(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"message": "success !",
	})
}
