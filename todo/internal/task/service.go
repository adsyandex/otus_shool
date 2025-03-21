package task

import (
    "context"
    "fmt"
    "github.com/adsyandex/otus_shool/todo/internal/models"
    "github.com/adsyandex/otus_shool/todo/internal/storage"
)

type Service struct {
    store storage.Storage
}

func New(store storage.Storage) *Service {
    return &Service{store: store}
}

func (s *Service) AddTask(ctx context.Context, req models.TaskRequest) (models.TaskResponse, error) {
    if req.Name == "" || len(req.Name) > 100 {
        return models.TaskResponse{}, fmt.Errorf("invalid task name")
    }

    task := models.Task{Name: req.Name}
    if err := s.store.AddTask(ctx, task); err != nil {
        return models.TaskResponse{}, fmt.Errorf("storage error: %w", err)
    }

    return models.TaskResponse{
        ID:   task.ID,
        Name: task.Name,
    }, nil
}

func (s *Service) GetTasks(ctx context.Context, limit, offset int) ([]models.TaskResponse, error) {
    if limit <= 0 || limit > 100 {
        return nil, fmt.Errorf("invalid limit value: %d", limit)
    }

    tasks, err := s.store.GetTasks(ctx, limit, offset)
    if err != nil {
        return nil, fmt.Errorf("storage error: %w", err)
    }

    responses := make([]models.TaskResponse, len(tasks))
    for i, t := range tasks {
        responses[i] = models.TaskResponse{
            ID:   t.ID,
            Name: t.Name,
        }
    }
    return responses, nil
}