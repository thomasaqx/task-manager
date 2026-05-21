package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/thomasaqx/task-manager/internal/controller"
	"github.com/thomasaqx/task-manager/internal/db"
	"github.com/thomasaqx/task-manager/internal/repository"
	"github.com/thomasaqx/task-manager/internal/routes"
	"github.com/thomasaqx/task-manager/internal/service"
)

func main() {

	//database variable
	db := db.ConnectDb()
	
	//repositories
	userRepo := repository.NewUserRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	taskRepo := repository.NewTaskRepository(db)

	//services
	userSvc := service.NewUserService(userRepo)
	categorySvc := service.NewCategoryService(categoryRepo)
	taskSvc := service.NewTaskService(taskRepo)

	//controllers 
	userCtrl := controller.NewUserController(userSvc)
	categoryCtrl := controller.NewCategoryController(categorySvc)
	taskCtrl := controller.NewTaskController(taskSvc)
	
	r := gin.Default()

	//routes
	routes.RegisterUserRoutes(r, userCtrl)
	routes.RegisterCategoryRoutes(r, categoryCtrl)
	routes.RegisterTaskRoutes(r, taskCtrl)


	//init server
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	//booting the server
	r.Run(":" + port)

}
