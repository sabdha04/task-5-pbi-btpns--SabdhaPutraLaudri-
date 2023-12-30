package routes

import (
	"github.com/gin-gonic/gin"
	"FinalProject/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// User Endpoints
	userGroup := r.Group("/users")
	{
		userGroup.POST("/register", controllers.RegisterUser)
		userGroup.POST("/login", controllers.LoginUser)
		userGroup.PUT("/:userId", controllers.UpdateUser)
		userGroup.DELETE("/:userId", controllers.DeleteUser)
	}

	// Photos Endpoints
	photoGroup := r.Group("/photos")
	{
		photoGroup.POST("/", controllers.CreatePhoto)
		photoGroup.GET("/", controllers.GetPhotos)
		photoGroup.PUT("/:photoId", controllers.UpdatePhoto)
		photoGroup.DELETE("/:photoId", controllers.DeletePhoto)
	}

	return r
}
