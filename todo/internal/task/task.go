package task

import (
    "sync"
    "errors"
    "github.com/adsyandex/otus_shool/todo/internal/storage"
    "github.com/adsyandex/otus_shool/todo/internal/models" // Импортируем models
)

// TaskManager управляет списком задач
type TaskManager struct {
    tasks []models.Task // Используем models.Task
    mu    sync.Mutex
    store storage.Storage
}

// NewTaskManager создает новый менеджер задач
func NewTaskManager(store storage.Storage) *TaskManager {
    return &TaskManager{store: store}
}

// AddTask добавляет задачу в список
func (tm *TaskManager) AddTask(task models.Task) { // Используем models.Task
    tm.mu.Lock()
    defer tm.mu.Unlock()
    tm.tasks = append(tm.tasks, task)
}

// GetTasks возвращает копию списка задач и ошибку (если есть)
func (tm *TaskManager) GetTasks() ([]models.Task, error) { // Используем models.Task
    tm.mu.Lock()
    defer tm.mu.Unlock()

    if len(tm.tasks) == 0 {
        return nil, errors.New("список задач пуст")
    }

    return append([]models.Task{}, tm.tasks...), nil
}