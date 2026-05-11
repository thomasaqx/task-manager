package repository

import (
	"github.com/thomasaqx/task-manager/internal/models"
	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) CreateTask(task *models.Task) (*models.Task, error) {
	err := r.db.Create(task).Error
	return task, err
}

func (r *TaskRepository) FindByTaskId(id uint) (*models.Task, error) {
	task := models.Task{}
	err := r.db.First(&task, id).Error
	return &task, err
}

func (r *TaskRepository) FindAll() ([]models.Task, error) {
	var tasks = []models.Task{}
	err := r.db.Find(&tasks).Error
	return tasks, err

}

func (r *TaskRepository) DeleteById(id uint) (error) {
	task := models.Task{}
	err := r.db.Delete(&task, id).Error
	return err
}

