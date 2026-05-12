package repository

import (
	"github.com/thomasaqx/task-manager/internal/models"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

// NewCategoryRepository creates a new CategoryRepository with the provided GORM database instance.
// The db instance is injected from outside, keeping the repository decoupled from database configuration.
func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

// CreateCategory persists a new Category record to the database.
// It receives a pointer to a Category so GORM can populate auto-generated fields (ID, CreatedAt) back into the struct.
func (r *CategoryRepository) CreateCategory(category *models.Category) (*models.Category, error) {
	err := r.db.Create(category).Error
	return category, err
}

// FindByCategoryId retrieves a single Category by its primary key.
// Returns a pointer to the found Category, or an error if no record exists.
func (r *CategoryRepository) FindByCategoryId(id uint) (*models.Category, error) {
	category := models.Category{}
	err := r.db.First(&category, id).Error
	return &category, err
}

// FindAll retrieves all Category records from the database.
// Returns a slice of Categories, which is empty (not nil) when no records are found.
func (r *CategoryRepository) FindAll() ([]models.Category, error) {
	var categories = []models.Category{}
	err := r.db.Find(&categories).Error
	return categories, err

}

// DeleteById performs a soft delete on the Category with the given ID.
// Because the Category model embeds gorm.Model, GORM sets DeletedAt instead of removing the row.
func (r *CategoryRepository) DeleteById(id uint) error {
	category := models.Category{}
	err := r.db.Delete(&category, id).Error
	return err
}

// FindCategoryByUser retrieves all Categories belonging to a specific user.
//
// db.Where("user_id = ?", userId) builds a SQL WHERE clause — the "?" is a placeholder
// that GORM replaces safely with the userId value, preventing SQL injection.
// db.Find(&category) then executes the full query (WHERE + SELECT *) and scans the
// result rows directly into the category slice via the pointer.
func (r *CategoryRepository) FindCategoryByUser(userId uint) ([]models.Category, error) {
	var category = []models.Category{}
	err := r.db.Where("user_id = ?", userId).Find(&category).Error
	return category, err
}
