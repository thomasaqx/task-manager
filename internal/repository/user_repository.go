package repository

import (
	"github.com/thomasaqx/task-manager/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	//injecting the database
	db *gorm.DB
}

// NewUserRepository creates a new UserRepository with the provided GORM database instance.
// The db instance is injected from outside, keeping the repository decoupled from database configuration.
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// CreateUser persists a new User record to the database.
// It receives a pointer to a User so GORM can populate auto-generated fields (ID, CreatedAt) back into the struct.
func (r *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	//inserts
	err := r.db.Create(user).Error
	return user, err
}

// FindAll retrieves all User records from the database.
// Returns a slice of Users, which is empty (not nil) when no records are found.
func (r *UserRepository) FindAll() ([]models.User, error) {
	users := []models.User{}
	err := r.db.Find(&users).Error
	return users, err
}

// FindByUserId retrieves a single User by its primary key.
// Returns a pointer to the found User, or an error if no record exists.
//
// To filter by non-primary-key fields, combine db.Where with db.First or db.Find:
//
//	db.Where("email = ?", email).First(&user)
//
// The "?" placeholder is safely replaced by GORM, preventing SQL injection.
func (r *UserRepository) FindByUserId(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return &user, err
}

// DeleteById performs a soft delete on the User with the given ID.
// Because the User model embeds gorm.Model, GORM sets DeletedAt instead of removing the row.
func (r *UserRepository) DeleteById(id uint) error {
	var user models.User
	err := r.db.Delete(&user, id).Error
	return err
}
