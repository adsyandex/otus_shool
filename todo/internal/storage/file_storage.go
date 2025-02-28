package storage

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/adsyandex/otus_shool/internal/task"
)

// FileStorage реализует Storage с использованием файла
type FileStorage struct {
	filename string
	mu       sync.Mutex // Добавляем мьютекс для защиты доступа к файлу
}

// NewFileStorage создает новый экземпляр FileStorage
func NewFileStorage(filename string) *FileStorage {
	return &FileStorage{filename: filename}
}

// Save сохраняет задачу в файл (безопасно для нескольких горутин)
func (fs *FileStorage) Save(task task.Task) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	tasks, err := fs.Load()
	if err != nil {
		return err
	}

	tasks = append(tasks, task)

	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	return os.WriteFile(fs.filename, data, 0644)
}

// Load загружает задачи из файла (безопасно для нескольких горутин)
func (fs *FileStorage) Load() ([]task.Task, error) {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	data, err := os.ReadFile(fs.filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []task.Task{}, nil
		}
		return nil, err
	}

	var tasks []task.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}
