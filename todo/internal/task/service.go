package task

import (
	"context"
	"github.com/adsyandex/otus_shool/todo/internal/models"
	"github.com/adsyandex/otus_shool/todo/internal/storage"
)

type Service struct {
	storage storage.Storage
}

func NewService(storage storage.Storage) *Service {
	return &Service{storage: storage}
}

func (s *Service) AddTask(ctx context.Context, task models.Task) error {
	return s.storage.AddTask(ctx, task)
}

func (s *Service) GetTasks(ctx context.Context) ([]models.Task, error) {
	return s.storage.GetTasks(ctx)
}