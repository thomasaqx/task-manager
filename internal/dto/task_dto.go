package dto

import(
	"github.com/thomasaqx/task-manager/internal/model"
)

type TaskRequest struct{
	Title string `json:"title" binding:"required"` 
	Description string `json:"description"` 
	Status model.TaskStatus `json:"taskstatus" binding:"required"` 
	Priority model.Priority `json:"priority" binding:"required"`
	DueDate time.Time `json:"due_date"`
	UserId uint `json:"user_id" binding:"required"`
}

type TaskStatusResponse { 
	Task model.Task `json:"task" binding:"task"` 
}