// storage.go
package storage

import (
    "context"
    "github.com/adsyandex/otus_shool/todo/internal/models"
)

type Storage interface {
    AddTask(ctx context.Context, task models.Task) error
    GetTasks(ctx context.Context, limit, offset int) ([]models.Task, error)
}

// file_storage.go
package storage

import (
    "context"
    "github.com/adsyandex/otus_shool/todo/internal/models"
)

type FileStorage struct {
    // Реализация хранилища
}

func (s *FileStorage) AddTask(ctx context.Context, task models.Task) error {
    // Логика сохранения
    return nil
}

func (s *FileStorage) GetTasks(ctx context.Context, limit, offset int) ([]models.Task, error) {
    // Логика получения
    return []models.Task{}, nil
}