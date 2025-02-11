package storage

import (
    "errors"
    "github.com/adsyandex/otus_shool/internal/task"
)

var ErrTaskNotFound = errors.New("задача не найдена")

// Storage определяет интерфейс для работы с хранилищем задач.
type Storage interface {
    SaveTask(task.Task) error
    GetTasks() ([]task.Task, error)
    GetTaskByID(id int) (task.Task, error)
    UpdateTask(task.Task) error
    DeleteTask(id int) error
    GetNextID() int
}

