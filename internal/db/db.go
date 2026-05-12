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

// Creating the database
func ConnectDb() *gorm.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro to load .env")
	}

	dsn := os.Getenv("DB_URL")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	fmt.Println("Successfully connected!")

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Category{})
	db.AutoMigrate(&models.Task{})

	return db
}
