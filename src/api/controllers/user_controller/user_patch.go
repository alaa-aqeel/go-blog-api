package user_controller

import (
	"github.com/gin-gonic/gin"
)

func Update(ctx *gin.Context) {
	// var request request.CreateUserRequest
	// if err := ctx.ShouldBindBodyWithJSON(&request); err != nil {
	// 	ctx.JSON(400, gin.H{
	// 		"status": "error",
	// 		"errors": helpers.ParseValidationErrors(err.(validator.ValidationErrors)),
	// 	})
	// 	return
	// }

	// data := helpers.StructToMap(request)
	// err := service.User().Update(ctx.Param("id"), data)
	// if err != nil {
	// 	ctx.JSON(400, gin.H{
	// 		"status": "error",
	// 		"errors": err.Error(),
	// 	})
	// 	return
	// }

	// ctx.JSON(200, gin.H{
	// 	"data": data,
	// })
}
