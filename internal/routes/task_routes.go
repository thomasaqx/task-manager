package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thomasaqx/task-manager/internal/controller"
)

func RegisterTaskRoutes(r *gin.Engine, c *controller.TaskController) {
	r.POST("/tasks", c.CreateTask)
	r.GET("/tasks", c.GetAllTasks)
	r.GET("/tasks/:id", c.FindByTaskId)
	r.GET("/users/:userId/tasks", c.FindTaskByUserId)
	r.DELETE("/tasks/:id", c.DeleteByTaskId)

}
