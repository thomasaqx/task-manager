package service

import (
	"github.com/thomasaqx/task-manager/internal/dto"
	"github.com/thomasaqx/task-manager/internal/models"
	"github.com/thomasaqx/task-manager/internal/repository"
)

type CategoryService struct {
	categoryRepository *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{categoryRepository: repo}
}

func (s *CategoryService) CreateCategory(req dto.CategoryRequest) (*models.Category, error) {

	category := models.Category{
		Name:   req.Name,
		UserID: req.UserID,
	}

	categoryRepo, err := s.categoryRepository.CreateCategory(&category)
	if err != nil {
		return nil, err
	}

	return categoryRepo, err

}

func (s *CategoryService) DeleteByCategoryId(id uint) error {
	return s.categoryRepository.DeleteByCategoryID(id)
}

func (s *CategoryService) GetAllCategories() ([]models.Category, error) {
	return s.categoryRepository.GetAllCategories()
}

func (s *CategoryService) FindByCategoryId(id uint) (*models.Category, error) {
	return s.categoryRepository.FindByCategoryID(id)
}

func (s *CategoryService) FindCategoryByUserId(userId uint) ([]models.Category, error) {
	return s.categoryRepository.FindCategoryByUserID(userId)
}
