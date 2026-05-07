package model

import (
	"time"
)

type TaskStatus string
type Priority string


const(
	InProgress TaskStatus = "In progress"
	Pending TaskStatus = "Pending"
	Done TaskStatus = "Done"

	High Priority = "High"
	Medium Priority = "Medium"
	Low Priority = "Low"

)

type Task struct {
	Id uint `json:"id" db:"id"`
	Title string `json:"title" db:"title"` 
	Description string `json:"description" db:"description"` 
	Status TaskStatus `json:"status" db:"status"` 
	Priority Priority `json:"priority" db:"priority"`
	DueDate time.Time `json:"due_date" db:"due_date"`
	UserId *uint `json:"user_id" db:"user_id"`
	CategoryId *uint `json:"category_id" db:"category_id"`
}