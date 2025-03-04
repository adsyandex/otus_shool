package storage

import (
    "encoding/json"
    "os"
    "github.com/adsyandex/otus_shool/todo/internal/models" // Импортируем models
)

type FileStorage struct {
    filePath string
}

func NewFileStorage(filePath string) *FileStorage {
    return &FileStorage{filePath: filePath}
}

func (fs *FileStorage) GetTasks() ([]models.Task, error) { // Используем models.Task
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

func (fs *FileStorage) SaveTasks(tasks []models.Task) error { // Используем models.Task
    file, err := json.Marshal(tasks)
    if err != nil {
        return err
    }

    return os.WriteFile(fs.filePath, file, 0644)
}