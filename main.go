package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"006-go-api-orm-gorm/database"
	"006-go-api-orm-gorm/routes"
)

func main() {
	database.InitDB()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to 006-go-api-orm-gorm"})
	})

	routes.RegisterUserRoutes(r)

	r.Run(":8080")
	fmt.Println("Server running at http://localhost:8080")
}
