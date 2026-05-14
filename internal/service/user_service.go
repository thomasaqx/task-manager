package service

import (
	"github.com/thomasaqx/task-manager/internal/dto"
	"github.com/thomasaqx/task-manager/internal/models"
	"github.com/thomasaqx/task-manager/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{userRepository: repo}
}

func (s *UserService) CreateUser(req dto.UserRequest) (*models.User, error) {
	password, err := HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: password,
	}

	userRepo, err := s.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return userRepo, err
}

func (s *UserService) FindAll() ([]models.User, error) {
	return s.userRepository.FindAll()

}

func (s *UserService) FindByUserId(id uint) (*models.User, error) {
	return s.userRepository.FindByUserId(id)
}

func (s *UserService) Delete(id uint) error {
	return s.userRepository.DeleteById(id)

}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
