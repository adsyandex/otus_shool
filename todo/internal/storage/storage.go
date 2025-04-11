package storage

import (
	"context"
	"errors"
	"github.com/adsyandex/otus_shool/todo/internal/models"
)

var ErrNotFound = errors.New("not found")


type Storage interface {
	CreateTask(ctx context.Context, task *models.Task) error
	GetTask(ctx context.Context, id string) (*models.Task, error)
	UpdateTask(ctx context.Context, task *models.Task) error
	DeleteTask(ctx context.Context, id string) error
	ListTasks(ctx context.Context) ([]*models.Task, error) 
}