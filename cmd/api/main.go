package main

import (
	"os"

	"github.com/gin-gonic/gin"

	//"github.com/thomasaqx/task-manager/internal/controller"
	"github.com/thomasaqx/task-manager/internal/db"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	r := gin.Default()

	//TODO: Register routes
	// r.GET("", controller.getTodo)

	
	db := db.ConnectDb()
	_ = db

	//booting the server
	r.Run(":" + port)

}
