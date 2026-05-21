package models

import (
	"time"

	"gorm.io/gorm"
)

type TaskStatus string
type Priority string

const (
	InProgress TaskStatus = "In progress"
	Pending    TaskStatus = "Pending"
	Done       TaskStatus = "Done"

	High   Priority = "High"
	Medium Priority = "Medium"
	Low    Priority = "Low"
)

type Task struct {
	gorm.Model
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	Priority    Priority   `json:"priority"`
	DueDate     time.Time  `json:"due_date"`
	UserID      *uint      `json:"user_id"`
	CategoryID  *uint      `json:"category_id"`
}
