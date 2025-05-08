package service

import (
    "context"
    "github.com/adsyandex/otus_shool/todo/internal/models"
    "github.com/adsyandex/otus_shool/todo/internal/storage"
)

type TaskService struct {
    repo storage.Storage
}

func NewTaskService(repo storage.Storage) *TaskService {
    return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(title string) (*models.Task, error) {
    task := &models.Task{
        Title: title,
        Done: false,
    }
    err := s.repo.SaveTask(context.Background(), *task)
    return task, err
}