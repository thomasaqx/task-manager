package dto

import (
	"time"

	"github.com/thomasaqx/task-manager/internal/models"
)

type TaskRequest struct {
	Title       string           `json:"title" binding:"required"`
	Description string           `json:"description"`
	Status      models.TaskStatus `json:"status" binding:"required"`
	Priority    models.Priority   `json:"priority" binding:"required"`
	DueDate     time.Time        `json:"due_date"`
	CategoryId  uint             `json:"category_id"`
	UserId      uint             `json:"user_id" binding:"required"`
}

type TaskResponse struct {
	Id          uint             `json:"id"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Status      models.TaskStatus `json:"status"`
	Priority    models.Priority   `json:"priority"`
	DueDate     time.Time        `json:"due_date"`
}
