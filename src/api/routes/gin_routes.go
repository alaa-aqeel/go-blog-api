package routes

import (
	"github.com/alaa-aqeel/govalid/src/api/controllers/user_controller"
	"github.com/alaa-aqeel/govalid/src/domain/repository"
	"github.com/alaa-aqeel/govalid/src/domain/services"
	"github.com/alaa-aqeel/govalid/src/pkgs/database"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {

	api := r.Group("/api/v1")
	{
		userRoute := api.Group("/users")
		{
			userRepository := repository.NewUserRepository(database.DB)
			userService := services.NewUserService(userRepository)
			userController := user_controller.NewUserController(userService)

			userRoute.POST("/", userController.Store)
			userRoute.GET("/", userController.Index)
			userRoute.GET("/:id", userController.Show)
			userRoute.Match([]string{"PATCH", "PUT"}, "/:id", userController.Update)
		}
	}
}
