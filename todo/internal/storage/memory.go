// internal/storage/memory.go
package storage

import (
	"context"
	"sync"
	"github.com/adsyandex/otus_shool/todo/internal/models"
)

type MemoryStorage struct {
	mu    sync.Mutex
	tasks map[string]models.Task
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		tasks: make(map[string]models.Task),
	}
}

func (s *MemoryStorage) AddTask(ctx context.Context, task models.Task) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tasks[task.ID] = task
	return nil
}

func (s *MemoryStorage) GetTasks(ctx context.Context) ([]models.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	result := make([]models.Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		result = append(result, task)
	}
	return result, nil
}