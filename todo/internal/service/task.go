package service

import (
	"context"
	"errors"

	"github.com/adsyandex/otus_shool/todo/internal/models"
	"github.com/adsyandex/otus_shool/todo/internal/storage/contracts"
)

var (
	ErrTaskNotFound = errors.New("task not found")
)

type TaskService struct {
	storage contracts.Storage
}

func NewTaskService(storage contracts.Storage) *TaskService {
	return &TaskService{storage: storage}
}

func (s *TaskService) CreateTask(ctx context.Context, task models.Task) (models.Task, error) {
	if task.Title == "" {
		return models.Task{}, errors.New("title is required")
	}

	return s.storage.CreateTask(ctx, task)
}

func (s *TaskService) GetTask(ctx context.Context, id string) (models.Task, error) {
	if id == "" {
		return models.Task{}, errors.New("id is required")
	}

	return s.storage.GetTask(ctx, id)
}

func (s *TaskService) UpdateTask(ctx context.Context, task models.Task) (models.Task, error) {
	if task.ID == "" {
		return models.Task{}, errors.New("id is required")
	}

	if task.Title == "" {
		return models.Task{}, errors.New("title is required")
	}

	return s.storage.UpdateTask(ctx, task)
}

func (s *TaskService) DeleteTask(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id is required")
	}

	return s.storage.DeleteTask(ctx, id)
}

func (s *TaskService) ListTasks(ctx context.Context, filter models.TaskFilter) ([]models.Task, error) {
	return s.storage.ListTasks(ctx, filter)
}

func (s *TaskService) CompleteTaskWithLog(ctx context.Context, taskID string, logMessage string) error {
	if taskID == "" {
		return errors.New("task id is required")
	}

	if logMessage == "" {
		return errors.New("log message is required")
	}

	return s.storage.WithTransaction(ctx, func(tx contracts.Storage) error {
		// Получаем задачу
		task, err := tx.GetTask(ctx, taskID)
		if err != nil {
			return err
		}

		// Обновляем статус задачи
		task.Completed = true
		if _, err := tx.UpdateTask(ctx, task); err != nil {
			return err
		}

		// Здесь можно добавить логирование в другую таблицу...
		// Например:
		// if err := tx.CreateLog(ctx, models.Log{
		//     TaskID:  taskID,
		//     Message: logMessage,
		//     Time:    time.Now(),
		// }); err != nil {
		//     return err
		// }

		return nil
	})
}
