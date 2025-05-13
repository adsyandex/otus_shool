package storage

import (
	"context"
	"errors"
	"time"

	"github.com/adsyandex/otus_shool/todo/internal/models"
)

var (
	ErrNotFound = errors.New("not found")
)

type Logger interface {
	LogAction(ctx context.Context, action string, ttl time.Duration) error
	Close() error
}

type Storage interface {
	SaveTask(ctx context.Context, task models.Task) error
	GetTask(ctx context.Context, id string) (*models.Task, error)
	UpdateTask(ctx context.Context, task models.Task) error
	DeleteTask(ctx context.Context, id string) error
	ListTasks(ctx context.Context, completed *bool) ([]models.Task, error)
	Close(ctx context.Context) error
}