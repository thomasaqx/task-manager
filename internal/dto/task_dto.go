package dto

import(
	"github.com/thomasaqx/task-manager/internal/model"
	"time"
)

type TaskRequest struct{
	Title string `json:"title" binding:"required"` 
	Description string `json:"description"` 
	Status model.TaskStatus `json:"status" binding:"required"` 
	Priority model.Priority `json:"priority" binding:"required"`
	DueDate time.Time `json:"due_date"`
	CategoryId uint `json:"category_id"`
	UserId uint `json:"user_id" binding:"required"`
}

type TaskResponse struct {
	Id uint `json:"id"`
	Title string `json:"title"` 
	Description string `json:"description"` 
	Status model.TaskStatus `json:"status"` 
	Priority model.Priority `json:"priority"`
	DueDate time.Time `json:"due_date"`
}