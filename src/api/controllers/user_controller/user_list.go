package user_controller

import (
	"github.com/gin-gonic/gin"
)

func (u *UserController) Index(ctx *gin.Context) {
	users, err := u.service.GetAll()
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": users})
}

func (u *UserController) Show(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	user, err := u.service.Find("id", id)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(200, gin.H{"data": user})
}
