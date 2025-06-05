package routes

import (
	"github.com/gin-gonic/gin"
	"006-go-api-orm-gorm/controllers"
)

func RegisterUserRoutes(r *gin.Engine) {
	api := r.Group("/api/users")
	{
		api.GET("/", controllers.GetUsers)
		api.GET("/:id", controllers.GetUser)
		api.POST("/", controllers.CreateUser)
	}
}
