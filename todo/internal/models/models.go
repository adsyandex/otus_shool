// internal/storage/file_storage.go
package storage

import (
    "context"
    "github.com/adsyandex/otus_shool/todo/internal/models"
)

type FileStorage struct {
    // реализация
}

func (s *FileStorage) AddTask(ctx context.Context, task models.Task) error {
    // логика сохранения в файл
}

func (s *FileStorage) GetTasks(ctx context.Context, limit, offset int) ([]models.Task, error) {
    // логика чтения из файла
}
