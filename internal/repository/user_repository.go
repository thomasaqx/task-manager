package repository

import (
	"github.com/thomasaqx/task-manager/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	//injetando o database
	db *gorm.DB
}

// cria o database
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	err := r.db.Create(user).Error
	return user, err
}

// retornar todos os usuários
func (r *UserRepository) FindAll() ([]models.User, error) {
	users := []models.User{}
	err := r.db.Find(&users).Error
	return users, err
}

// procura id no database
func (r *UserRepository) FindByUserId(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return &user, err
}

// deleta por id
func (r *UserRepository) DeleteById(id uint) (error) {
	var user models.User
	err := r.db.Delete(&user, id).Error
	return err
}
