package database

import (
	"fmt"
	"os"

	"belajar-gin-gorm/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB_HOST     = os.Getenv("DB_HOST")
	DB_USER     = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_PORT     = os.Getenv("DB_PORT")
	DB_NAME     = os.Getenv("DB_NAME")
	DEBUG_MODE  = os.Getenv("DEBUG_MODE") // true/false
	db          *gorm.DB
	err         error
)

func Postgres() *gorm.DB {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if DEBUG_MODE == "true" {
		db.Debug().AutoMigrate(models.Product{}, models.User{})
		return db
	}

	db.AutoMigrate(models.User{}, models.Product{})
	return db
}
