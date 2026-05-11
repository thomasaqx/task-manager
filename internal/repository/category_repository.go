package repository

import (
	"github.com/thomasaqx/task-manager/internal/models"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) CreateCategory(category *models.Category) (*models.Category, error) {
	err := r.db.Create(category).Error
	return category, err
}

func (r *CategoryRepository) FindByCategoryId(id uint) (*models.Category, error) {
	category := models.Category{}
	err := r.db.First(&category, id).Error
	return &category, err
}

func (r *CategoryRepository) FindAll() ([]models.Category, error) {
	var categories = []models.Category{}
	err := r.db.Find(&categories).Error
	return categories, err

}

func (r *CategoryRepository) DeleteById(id uint) (error) {
	category := models.Category{}
	err := r.db.Delete(&category, id).Error
	return err
}
