package storage

import (
    "encoding/json"
    //"errors"
    "os"
    "sync"
    "github.com/adsyandex/otus_shool/internal/task"
    
)

// FileStorage реализует Storage с использованием файла JSON.
type FileStorage struct {
    filePath string
    tasks    []task.Task
    mu       sync.Mutex
}

// NewFileStorage создает новый экземпляр FileStorage.
func NewFileStorage(filePath string) *FileStorage {
    fs := &FileStorage{filePath: filePath}
    fs.loadTasks()
    return fs
}

// loadTasks загружает задачи из файла.
func (fs *FileStorage) loadTasks() {
    data, err := os.ReadFile(fs.filePath)
    if err != nil {
        fs.tasks = []task.Task{}
        return
    }
    json.Unmarshal(data, &fs.tasks)
}

// saveTasks сохраняет задачи в файл.
func (fs *FileStorage) saveTasks() error {
    data, err := json.Marshal(fs.tasks)
    if err != nil {
        return err
    }
    return os.WriteFile(fs.filePath, data, 0644)
}

// SaveTask сохраняет задачу.
func (fs *FileStorage) SaveTask(t task.Task) error {
    fs.mu.Lock()
    defer fs.mu.Unlock()
    fs.tasks = append(fs.tasks, t)
    return fs.saveTasks()
}

// GetTasks возвращает все задачи.
func (fs *FileStorage) GetTasks() ([]task.Task, error) {
    return fs.tasks, nil
}

// GetTaskByID возвращает задачу по ID.
func (fs *FileStorage) GetTaskByID(id int) (task.Task, error) {
    for _, t := range fs.tasks {
        if t.ID == id {
            return t, nil
        }
    }
    return task.Task{}, ErrTaskNotFound
}

// UpdateTask обновляет задачу.
func (fs *FileStorage) UpdateTask(t task.Task) error {
    fs.mu.Lock()
    defer fs.mu.Unlock()
    for i, task := range fs.tasks {
        if task.ID == t.ID {
            fs.tasks[i] = t
            return fs.saveTasks()
        }
    }
    return ErrTaskNotFound
}

// DeleteTask удаляет задачу по ID.
func (fs *FileStorage) DeleteTask(id int) error {
    fs.mu.Lock()
    defer fs.mu.Unlock()
    for i, t := range fs.tasks {
        if t.ID == id {
            fs.tasks = append(fs.tasks[:i], fs.tasks[i+1:]...)
            return fs.saveTasks()
        }
    }
    return ErrTaskNotFound
}

// GetNextID возвращает следующий ID для задачи.
func (fs *FileStorage) GetNextID() int {
    if len(fs.tasks) == 0 {
        return 1
    }
    return fs.tasks[len(fs.tasks)-1].ID + 1
}

