package main

import (
	"belajar-gin-gorm/controllers"
	"belajar-gin-gorm/database"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const PORT = ":3000"

func main() {
	fmt.Println("Starting server.......")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	dbPort, _ := strconv.Atoi(os.Getenv("DB_USER"))
	dbConf := database.Database{
		Host:      os.Getenv("DB_Host"),
		Username:  os.Getenv("DB_USER"),
		Password:  os.Getenv("DB_Password"),
		Port:      dbPort,
		Name:      os.Getenv("DB_NAME"),
		DebugMode: os.Getenv("DEBUG_MODE"),
	}

	db := database.Postgres(&dbConf)
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

	r.Run(PORT)
}
