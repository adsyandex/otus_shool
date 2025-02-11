package task

import (
    "errors"
    "time"
    "github.com/adsyandex/otus_shool/internal/storage"
)

// Task представляет структуру задачи.
type Task struct {
    ID          int       `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Status      string    `json:"status"`
    CreatedAt   time.Time `json:"created_at"`
}

// TaskManager управляет задачами.
type TaskManager struct {
    storage storage.Storage
}

// NewTaskManager создает новый экземпляр TaskManager.
func NewTaskManager(storage storage.Storage) *TaskManager {
    return &TaskManager{storage: storage}
}

// AddTask добавляет новую задачу.
func (tm *TaskManager) AddTask(title, description string) error {
    if title == "" {
        return errors.New("заголовок задачи не может быть пустым")
    }
    task := Task{
        ID:          tm.storage.GetNextID(),
        Title:       title,
        Description: description,
        Status:      "в процессе",
        CreatedAt:   time.Now(),
    }
    return tm.storage.SaveTask(task)
}

// GetTasks возвращает список всех задач.
func (tm *TaskManager) GetTasks() ([]Task, error) {
    return tm.storage.GetTasks()
}

// GetTaskByID возвращает задачу по ID.
func (tm *TaskManager) GetTaskByID(id int) (Task, error) {
    return tm.storage.GetTaskByID(id)
}

// UpdateTask обновляет задачу.
func (tm *TaskManager) UpdateTask(id int, title, description, status string) error {
    task, err := tm.storage.GetTaskByID(id)
    if err != nil {
        return err
    }
    if title != "" {
        task.Title = title
    }
    if description != "" {
        task.Description = description
    }
    if status != "" {
        task.Status = status
    }
    return tm.storage.UpdateTask(task)
}

// DeleteTask удаляет задачу по ID.
func (tm *TaskManager) DeleteTask(id int) error {
    return tm.storage.DeleteTask(id)
}

