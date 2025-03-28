package user_controller

import (
	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *UserController) Store(ctx *gin.Context) {
	var request CreateUserRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := u.service.Create(map[string]any{
		"name":     request.Name,
		"username": request.Username,
		"password": request.Password,
	}); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "User created successfully",
		"data": map[string]any{
			"name":     request.Name,
			"username": request.Username,
		},
	})
}
