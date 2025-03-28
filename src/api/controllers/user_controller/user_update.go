package user_controller

import (
	"github.com/gin-gonic/gin"
)

type UpdateUserRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *UserController) Update(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	var request UpdateUserRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := u.service.Update(id, map[string]any{
		"name":     request.Name,
		"username": request.Username,
		"password": request.Password,
	}); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "User updated successfully",
		"data": map[string]any{
			"name":     request.Name,
			"username": request.Username,
		},
	})
}
