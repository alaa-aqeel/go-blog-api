package routes

import (
	"github.com/alaa-aqeel/govalid/src/api/controllers/user_controller"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		userRoute := api.Group("/users")
		{
			userRoute.POST("/", user_controller.PostUser)
		}
	}
}
