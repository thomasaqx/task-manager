package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thomasaqx/task-manager/internal/dto"
	"github.com/thomasaqx/task-manager/internal/service"
)

type CategoryController struct {
	categoryService *service.CategoryService
}

func NewCategoryController(categorySvc *service.CategoryService) *CategoryController {
	return &CategoryController{categoryService: categorySvc}
}

func (c *CategoryController) CreateCategory(ctx *gin.Context) {

	var req dto.CategoryRequest

	//extract data from body request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//creating a category Model
	category, err := c.categoryService.CreateCategory(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//response from API
	ctx.JSON(http.StatusCreated, gin.H{"data": category})
}

func (c *CategoryController) GetAllCategories(ctx *gin.Context) {

	category, err := c.categoryService.GetAllCategories()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": category})

}

func (c *CategoryController) FindByCategoryId(ctx *gin.Context) {
	idStr := ctx.Param("id")

	parsedId, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uint(parsedId)

	category, err := c.categoryService.FindByCategoryId(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": category})

}

func (c *CategoryController) DeleteByCategoryId(ctx *gin.Context) {
	idStr := ctx.Param("id")

	parsedId, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uint(parsedId)

	err = c.categoryService.DeleteByCategoryId(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)

}

func (c *CategoryController) FindCategoryByUserId(ctx *gin.Context) {

	userIdStr := ctx.Param("userId")

	parsedId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := uint(parsedId)

	category, err := c.categoryService.FindCategoryByUserId(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": category})
}
