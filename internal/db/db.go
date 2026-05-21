package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/thomasaqx/task-manager/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Creating the database connection
func ConnectDb() *gorm.DB {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro to load .env")
	}

	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		log.Fatal("DB_URL is not defined")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	fmt.Println("Successfully connected!")

	db.AutoMigrate(&models.User{}, &models.Category{}, &models.Task{})

	return db
}
