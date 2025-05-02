package service

import (
	"context"
	"errors"
	"github.com/adsyandex/otus_shool/todo/internal/models"
	"github.com/adsyandex/otus_shool/todo/internal/storage"
)

// TaskService реализует бизнес-логику работы с задачами
type TaskService struct {
	redis_logger redis_logger.Storage
}

// NewTaskService создает новый экземпляр TaskService
func NewTaskService(redis_logger redis_logger.Storage) *TaskService {
	return &TaskService{redis_logger: redis_logger}
}

// CreateTask создает новую задачу с валидацией
func (s *TaskService) CreateTask(ctx context.Context, task models.Task) (models.Task, error) {
	if task.Title == "" {
		return models.Task{}, errors.New("title cannot be empty")
	}
	return s.redis_logger.SaveTask(ctx, task)
}

// GetTask возвращает задачу по ID
func (s *TaskService) GetTask(ctx context.Context, id int) (models.Task, error) {
	return s.redis_logger.GetTaskByID(ctx, id)
}

// GetAllTasks возвращает все задачи
func (s *TaskService) GetAllTasks(ctx context.Context) ([]models.Task, error) {
	return s.redis_logger.GetTasks(ctx)
}

// UpdateTask обновляет существующую задачу
func (s *TaskService) UpdateTask(ctx context.Context, task models.Task) (models.Task, error) {
	if task.ID == 0 || task.Title == "" {
		return models.Task{}, errors.New("invalid task data")
	}
	return s.redis_logger.UpdateTask(ctx, task)
}

// DeleteTask удаляет задачу по ID
func (s *TaskService) DeleteTask(ctx context.Context, id int) error {
	return s.redis_logger.DeleteTask(ctx, id)
}