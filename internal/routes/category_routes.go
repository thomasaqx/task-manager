package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thomasaqx/task-manager/internal/controller"
)

func RegisterCategoryRoutes(r *gin.Engine, c *controller.CategoryController) {
	r.POST("/category", c.CreateCategory)
	r.GET("/category", c.GetAllCategories)
	r.GET("/category/:id", c.FindByCategoryId)
	r.GET("/users/:userId/categories", c.FindCategoryByUserId)
	r.DELETE("/category/:id", c.DeleteByCategoryId)
}
