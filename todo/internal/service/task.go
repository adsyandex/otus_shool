package service

import (
	"context"
	"errors"
	"github.com/adsyandex/otus_shool/todo/internal/models"
	"github.com/adsyandex/otus_shool/todo/internal/storage"
)

// TaskService реализует бизнес-логику работы с задачами
type TaskService struct {
	storage storage.Storage
}

// NewTaskService создает новый экземпляр TaskService
func NewTaskService(storage storage.Storage) *TaskService {
	return &TaskService{storage: storage}
}

// CreateTask создает новую задачу с валидацией
func (s *TaskService) CreateTask(ctx context.Context, task models.Task) (models.Task, error) {
	if task.Title == "" {
		return models.Task{}, errors.New("title cannot be empty")
	}
	return s.storage.SaveTask(ctx, task)
}

// GetTask возвращает задачу по ID
func (s *TaskService) GetTask(ctx context.Context, id int) (models.Task, error) {
	return s.storage.GetTaskByID(ctx, id)
}

// GetAllTasks возвращает все задачи
func (s *TaskService) GetAllTasks(ctx context.Context) ([]models.Task, error) {
	return s.storage.GetTasks(ctx)
}

// UpdateTask обновляет существующую задачу
func (s *TaskService) UpdateTask(ctx context.Context, task models.Task) (models.Task, error) {
	if task.ID == 0 || task.Title == "" {
		return models.Task{}, errors.New("invalid task data")
	}
	return s.storage.UpdateTask(ctx, task)
}

// DeleteTask удаляет задачу по ID
func (s *TaskService) DeleteTask(ctx context.Context, id int) error {
	return s.storage.DeleteTask(ctx, id)
}