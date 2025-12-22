package routes

import (
	"gin-user-api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	r *gin.Engine,
	userHandler handlers.UserHandler,
	profileHandler handlers.ProfileHandler, 
) {
	api := r.Group("/api")
	{
		users := api.Group("/users")
		{
			users.GET("", userHandler.GetUsers)
			users.POST("", userHandler.CreateUser)
			users.GET("/:id", userHandler.GetUser)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)

			users.POST("/:id/profile", profileHandler.CreateProfile)
			users.PUT("/:id/profile", profileHandler.UpdateProfile)
		}
	}
}

