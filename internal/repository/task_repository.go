package repository

import (
	"github.com/thomasaqx/task-manager/internal/models"
	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

// NewTaskRepository creates a new TaskRepository with the provided GORM database instance.
// The db instance is injected from outside, keeping the repository decoupled from database configuration.
func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

// CreateTask persists a new Task record to the database.
// It receives a pointer to a Task so GORM can populate auto-generated fields (ID, CreatedAt) back into the struct.
func (r *TaskRepository) CreateTask(task *models.Task) (*models.Task, error) {
	err := r.db.Create(task).Error
	return task, err
}

// FindByTaskID retrieves a single Task by its primary key.
// Returns a pointer to the found Task, or an error if no record exists.
func (r *TaskRepository) FindByTaskID(id uint) (*models.Task, error) {
	task := models.Task{}
	err := r.db.First(&task, id).Error
	return &task, err
}

// GetAllTasks retrieves all Task records from the database.
// Returns a slice of Tasks, which is empty (not nil) when no records are found.
func (r *TaskRepository) GetAllTasks() ([]models.Task, error) {
	var tasks = []models.Task{}
	err := r.db.Find(&tasks).Error
	return tasks, err

}

// DeleteByTaskID performs a soft delete on the Task with the given ID.
// Because the Task model embeds gorm.Model, GORM sets DeletedAt instead of removing the row.
func (r *TaskRepository) DeleteByTaskID(id uint) error {
	task := models.Task{}
	err := r.db.Delete(&task, id).Error
	return err
}

// FindTaskByUserID retrieves all Tasks belonging to a specific user.
//
// db.Where("user_id = ?", userId) builds a SQL WHERE clause — the "?" is a placeholder
// that GORM replaces safely with the userId value, preventing SQL injection.
// db.Find(&task) then executes the full query (WHERE + SELECT *) and scans the
// result rows directly into the task slice via the pointer.
func (r *TaskRepository) FindTaskByUserID(userId uint) ([]models.Task, error) {
	var task = []models.Task{}
	err := r.db.Where("user_id = ?", userId).Find(&task).Error
	return task, err
}
