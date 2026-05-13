package service

import (
	"log"

	"github.com/thomasaqx/task-manager/internal/dto"
	"github.com/thomasaqx/task-manager/internal/models"
	"github.com/thomasaqx/task-manager/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{userRepo: repo}
}

func (s *UserService) CreateUser(req dto.UserRequest) (*models.User, error) {
	password, err := HashPassword(req.Password)
	if err != nil {
		log.Printf("Password is null")
	}

	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: password,
	}
	
	repoUser, err := s.userRepo.CreateUser(user)
	if err != nil {
		log.Printf("Something is wrong with user")
	}
	_ = repoUser

	return repoUser, err
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
