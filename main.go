package main

import (
	"belajar-gin-gorm/controllers"
	"belajar-gin-gorm/database"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	fmt.Println("Starting server.......")

	db := database.Postgres()
	controllers := controllers.New(db)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Users
	users := r.Group("/v1/users")
	{
		users.POST("/", controllers.CreateUser)
		users.GET("/products", controllers.GetUsersWithProducts)
	}

	r.Run(fmt.Sprintf(":%s", port))
}
