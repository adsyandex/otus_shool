package storage

import (
	"context"
	"errors"
	"github.com/adsyandex/otus_shool/todo/internal/models"
)

var (
	ErrNotFound = errors.New("task not found")
)

type Storage interface {
	GetTasks(ctx context.Context) ([]models.Task, error)
	GetTaskByID(ctx context.Context, id int) (models.Task, error)
	SaveTask(ctx context.Context, task models.Task) (models.Task, error) // Теперь возвращает задачу
	UpdateTask(ctx context.Context, task models.Task) (models.Task, error)
	DeleteTask(ctx context.Context, id int) error
}