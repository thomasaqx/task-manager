package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/thomasaqx/task-manager/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDb() *gorm.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	dsn := os.Getenv("DB_URL")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao conectar no banco:", err)
	}

	fmt.Println("Conectado com sucesso via GORM!")

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Category{})
	db.AutoMigrate(&model.Task{})

	return db
}
