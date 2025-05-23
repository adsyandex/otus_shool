package contracts

import (
	"context"
	"errors"
	"time"

	"github.com/adsyandex/otus_shool/todo/internal/models"
)

var (
	ErrTaskNotFound = errors.New("task not found")
)

type Logger interface {
	LogAction(ctx context.Context, action string, ttl time.Duration) error
	Close() error
	Error(msg string, args ...interface{})
	Info(msg string, args ...interface{})
}

type Storage interface {
	// Основные CRUD операции
	CreateTask(ctx context.Context, task models.Task) (models.Task, error)
	GetTask(ctx context.Context, id string) (models.Task, error)
	UpdateTask(ctx context.Context, task models.Task) (models.Task, error)
	DeleteTask(ctx context.Context, id string) error
	ListTasks(ctx context.Context, filter models.TaskFilter) ([]models.Task, error)

	// Дополнительные методы
	Close(ctx context.Context) error
	WithTransaction(ctx context.Context, fn func(tx Storage) error) error
}
