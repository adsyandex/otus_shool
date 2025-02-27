package storage

import "todo-app/internal/task"

// Storage определяет интерфейс для работы с хранилищем задач
type Storage interface {
	Save(task task.Task) error
	Load() ([]task.Task, error)
}
