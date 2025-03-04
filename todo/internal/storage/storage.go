package storage

import (
	"github.com/adsyandex/otus_shool/todo/internal/models"
)

// Storage определяет интерфейс для работы с хранилищем задач
type Storage interface {
	GetTasks() ([]models.Task, error)
	SaveTasks(tasks []models.Task) error
	//Save(task task.Task) error
	//Load() ([]task.Task, error)
}
