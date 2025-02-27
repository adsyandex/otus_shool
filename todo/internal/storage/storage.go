package storage

import "github.com/adsyandex/otus_shool/internal/task"

// Storage определяет интерфейс для работы с хранилищем задач
type Storage interface {
	Save(task task.Task) error
	Load() ([]task.Task, error)
}
