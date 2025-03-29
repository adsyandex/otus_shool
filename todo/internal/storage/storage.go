package storage

import (
	"context"
	"errors"
	"github.com/adsyandex/otus_shool/todo/internal/models"
)

var ErrNotFound = errors.New("task not found") // Одно определение на весь пакет

type Storage interface {
	GetTasks(ctx context.Context) ([]models.Task, error)
	GetTaskByID(ctx context.Context, id int) (models.Task, error)
	SaveTask(ctx context.Context, task models.Task) error
	UpdateTask(ctx context.Context, task models.Task) error
	DeleteTask(ctx context.Context, id int) error
}
