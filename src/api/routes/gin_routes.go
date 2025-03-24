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
			userRoute.POST("/", user_controller.Store)
			userRoute.GET("/", user_controller.Index)
			userRoute.GET("/:id", user_controller.Show)
			userRoute.Match([]string{"PATCH", "PUT"}, "/:id", user_controller.Update)
		}
	}
}
