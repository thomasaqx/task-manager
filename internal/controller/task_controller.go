package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thomasaqx/task-manager/internal/dto"
	"github.com/thomasaqx/task-manager/internal/service"
)

type TaskController struct {
	taskService *service.TaskService
}

func NewTaskController(TaskSvc *service.TaskService) *TaskController {
	return &TaskController{taskService: TaskSvc}
}

func (c *TaskController) CreateTask(ctx *gin.Context) {
	var req dto.TaskRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := c.taskService.CreateTask(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": task})
}

func (c *TaskController) FindAll(ctx *gin.Context) {

	tasks, err := c.taskService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": tasks})
}

func (c *TaskController) FindByTaskId(ctx *gin.Context) {
	idStr := ctx.Param("id")

	parsedId, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uint(parsedId)

	task, err := c.taskService.FindByTaskId(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": task})
}

func (c *TaskController) FindTaskByUser(ctx *gin.Context) {
	userIdStr := ctx.Param("userId")

	parsedId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := uint(parsedId)

	task, err := c.taskService.FindTaskByUser(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": task})
}

func (c *TaskController) DeleteByTaskId(ctx *gin.Context) {

	idStr := ctx.Param("id")

	parsedId, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uint(parsedId)

	if err = c.taskService.DeleteByTaskId(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)

}
