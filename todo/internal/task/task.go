package task

import (
    "context"
    "errors"
    "sync"
    "github.com/adsyandex/otus_shool/todo/internal/models"
    "github.com/adsyandex/otus_shool/todo/internal/storage"
)

type TaskManager struct {
    tasks []models.Task
    mu    sync.Mutex
    store storage.Storage
}

// NewTaskManager создает новый менеджер задач
func NewTaskManager(store storage.Storage) *TaskManager {
    return &TaskManager{store: store}
}

// AddTask добавляет задачу в список
func (tm *TaskManager) AddTask(task models.Task) {
    tm.mu.Lock()
    defer tm.mu.Unlock()
    tm.tasks = append(tm.tasks, task)
}

// GetTasks возвращает список задач
func (tm *TaskManager) GetTasks(ctx context.Context) ([]models.Task, error) {
    tm.mu.Lock()
    defer tm.mu.Unlock()

    select {
    case <-ctx.Done():
        return nil, ctx.Err()
    default:
        if len(tm.tasks) == 0 {
            return nil, errors.New("список задач пуст")
        }
        return append([]models.Task{}, tm.tasks...), nil
    }
}