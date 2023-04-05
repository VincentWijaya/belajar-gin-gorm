package database

import (
	"fmt"

	"belajar-gin-gorm/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Host      string
	Username  string
	Password  string
	Port      int
	Name      string
	DebugMode string
}

var (
	DB_HOST     = "localhost"
	DB_USER     = "postgres"
	DB_PASSWORD = "test123456"
	DB_PORT     = 5432
	DB_NAME     = "test"
	DEBUG_MODE  = false // true/false
	db          *gorm.DB
	err         error
)

func Postgres(conf *Database) *gorm.DB {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)

	if conf.Host != "" {
		config = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
			conf.Host, conf.Username, conf.Password, conf.Name, conf.Port)
	}

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if conf.DebugMode == "true" || DEBUG_MODE {
		fmt.Println(conf.DebugMode, DEBUG_MODE)
		db.Debug().AutoMigrate(models.Product{}, models.User{})
		return db
	}

	db.AutoMigrate(models.User{}, models.Product{})
	return db
}
