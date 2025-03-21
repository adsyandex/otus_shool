// internal/storage/storage.go
package storage

import "context"
import "github.com/adsyandex/otus_shool/todo/internal/models"

type Storage interface {
    AddTask(ctx context.Context, task models.Task) error
    GetTasks(ctx context.Context, limit, offset int) ([]models.Task, error)
}

// internal/storage/file_storage.go
package storage

type FileStorage struct {
    // Реализация
}

func (s *FileStorage) AddTask(ctx context.Context, task models.Task) error {
    // Логика
    return nil
}

func (s *FileStorage) GetTasks(ctx context.Context, limit, offset int) ([]models.Task, error) {
    // Логика
    return []models.Task{}, nil
}