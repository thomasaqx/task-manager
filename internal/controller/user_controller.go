package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thomasaqx/task-manager/internal/dto"
	"github.com/thomasaqx/task-manager/internal/service"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userSvc *service.UserService) *UserController {
	return &UserController{userService: userSvc}
}

func (c *UserController) CreateUser(ctx *gin.Context) {

	var req dto.UserRequest

	//extract data from body request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.userService.CreateUser(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//Status response from API
	ctx.JSON(http.StatusCreated, gin.H{"data": user})
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.userService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": users})
}

func (c *UserController) GetUserById(ctx *gin.Context) {
	//Create a String Param Id and convert him to uint
	idStr := ctx.Param("id")
	parsedId, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := uint(parsedId)

	user, err := c.userService.FindByUserId(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

func (c *UserController) DeleteUser(ctx *gin.Context) {

	idStr := ctx.Param("id")

	parsedId, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uint(parsedId)

	err = c.userService.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
