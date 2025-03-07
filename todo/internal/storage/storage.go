package storage

import (
    "context"
    "github.com/adsyandex/otus_shool/todo/internal/models"
)

// Storage определяет интерфейс для работы с хранилищем задач
type Storage interface {
    GetTasks(ctx context.Context) ([]models.Task, error)
    SaveTasks(ctx context.Context, tasks []models.Task) error
}