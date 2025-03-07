package storage

import (
    "context"
    "encoding/json"
    "os"
    "github.com/adsyandex/otus_shool/todo/internal/models"
)

type FileStorage struct {
    filePath string
}

func NewFileStorage(filePath string) *FileStorage {
    return &FileStorage{filePath: filePath}
}

func (fs *FileStorage) GetTasks(ctx context.Context) ([]models.Task, error) {
    select {
    case <-ctx.Done():
        return nil, ctx.Err()
    default:
        file, err := os.ReadFile(fs.filePath)
        if err != nil {
            return nil, err
        }

        var tasks []models.Task
        if err := json.Unmarshal(file, &tasks); err != nil {
            return nil, err
        }

        return tasks, nil
    }
}

func (fs *FileStorage) SaveTasks(ctx context.Context, tasks []models.Task) error {
    select {
    case <-ctx.Done():
        return ctx.Err()
    default:
        file, err := json.Marshal(tasks)
        if err != nil {
            return err
        }

        return os.WriteFile(fs.filePath, file, 0644)
    }
}