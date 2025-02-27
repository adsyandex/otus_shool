package storage

import (
	"encoding/json"
	"errors"
	"os"
	"github.com/adsyandex/otus_shool/internal/task"
)

var ErrTaskNotFound = errors.New("задача не найдена")

// Storage определяет интерфейс для работы с хранилищем задач
type Storage interface {
	SaveTask(task.Task) error
	GetTasks() ([]task.Task, error)
	GetTaskByID(id int) (task.Task, error)
	UpdateTask(task.Task) error
	DeleteTask(id int) error
	GetNextID() int
}

// FileStorage реализует интерфейс Storage с использованием файлового хранилища
type FileStorage struct {
	filename string
}

// NewFileStorage создает новый экземпляр FileStorage
func NewFileStorage(filename string) *FileStorage {
	return &FileStorage{filename: filename}
}

// SaveTask сохраняет задачу в файл
func (fs *FileStorage) SaveTask(task task.Task) error {
	tasks, err := fs.GetTasks()
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

// GetTasks возвращает все задачи
func (fs *FileStorage) GetTasks() ([]task.Task, error) {
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

// GetTaskByID возвращает задачу по ID
func (fs *FileStorage) GetTaskByID(id int) (task.Task, error) {
	tasks, err := fs.GetTasks()
	if err != nil {
		return task.Task{}, err
	}

	for _, t := range tasks {
		if t.ID == id {
			return t, nil
		}
	}
	return task.Task{}, ErrTaskNotFound
}

// UpdateTask обновляет задачу
func (fs *FileStorage) UpdateTask(updatedTask task.Task) error {
	tasks, err := fs.GetTasks()
	if err != nil {
		return err
	}

	for i, t := range tasks {
		if t.ID == updatedTask.ID {
			tasks[i] = updatedTask
			data, err := json.Marshal(tasks)
			if err != nil {
				return err
			}
			return os.WriteFile(fs.filename, data, 0644)
		}
	}

	return ErrTaskNotFound
}

// DeleteTask удаляет задачу по ID
func (fs *FileStorage) DeleteTask(id int) error {
	tasks, err := fs.GetTasks()
	if err != nil {
		return err
	}

	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			data, err := json.Marshal(tasks)
			if err != nil {
				return err
			}
			return os.WriteFile(fs.filename, data, 0644)
		}
	}

	return ErrTaskNotFound
}

// GetNextID возвращает следующий ID для задачи
func (fs *FileStorage) GetNextID() int {
	tasks, err := fs.GetTasks()
	if err != nil {
		return 1
	}

	if len(tasks) == 0 {
		return 1
	}
	return tasks[len(tasks)-1].ID + 1
}
