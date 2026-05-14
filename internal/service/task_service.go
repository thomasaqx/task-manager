package service

import (
	"github.com/thomasaqx/task-manager/internal/dto"
	"github.com/thomasaqx/task-manager/internal/models"
	"github.com/thomasaqx/task-manager/internal/repository"
)

type TaskService struct {
	taskRepository *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository ) *TaskService{
	return &TaskService{taskRepository: repo}
}

func (s *TaskService) CreateTask(req dto.TaskRequest) (*models.Task, error) {
	task := models.Task{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		Priority:    req.Priority,
		DueDate:     req.DueDate,
		UserId:      &req.UserId,
		CategoryId:  &req.CategoryId,
	}

	taskRepo, err := s.taskRepository.CreateTask(&task)
	if err != nil {
		return nil, err
	}

	return taskRepo, err
}

func (s *TaskService) DeleteByTaskId(id uint) error{
	err := s.taskRepository.DeleteByTaskId(id)
	return err
}

func (s *TaskService) FindAll() ([]models.Task, error){
	return s.taskRepository.FindAll()
}

func (s *TaskService) FindByTaskId(id uint) (*models.Task, error){
	return s.taskRepository.FindByTaskId(id)
}

func (s *TaskService) FindTaskByUser(userId uint) ([]models.Task, error){
	return s.taskRepository.FindTaskByUser(userId)
}