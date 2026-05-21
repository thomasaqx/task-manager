package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thomasaqx/task-manager/internal/controller"
)

func RegisterUserRoutes(r *gin.Engine, c *controller.UserController) {
	r.POST("/users", c.CreateUser)
	r.GET("/users", c.GetAllUsers)
	r.GET("/users/:id", c.GetUserById)
	r.DELETE("/users/:id", c.DeleteUser)
}
