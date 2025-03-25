// internal/storage/storage.go
package storage

import (
    "context"
    "github.com/adsyandex/otus_shool/todo/internal/models"
)

type Storage interface {
    AddTask(ctx context.Context, task models.Task) error
    GetTasks(ctx context.Context) ([]models.Task, error)
}