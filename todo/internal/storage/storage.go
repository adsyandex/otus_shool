package storage

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/adsyandex/otus_shool/todo/internal/models"
)

type Storage interface {
	AddTask(ctx context.Context, task models.Task) error
	GetTasks(ctx context.Context) ([]models.Task, error)
}

type MemoryStorage struct {
	mu    sync.Mutex
	tasks map[string]models.Task
	file  string
}

func NewMemoryStorage(jsonFile string) *MemoryStorage {
	// Создаем директорию, если ее нет
	if err := os.MkdirAll(filepath.Dir(jsonFile), 0755); err != nil {
		log.Printf("Warning: failed to create data directory: %v", err)
	}

	s := &MemoryStorage{
		tasks: make(map[string]models.Task),
		file:  jsonFile,
	}

	if err := s.loadFromFile(); err != nil {
		log.Printf("Warning: failed to load initial data: %v", err)
	}

	return s
}

func (s *MemoryStorage) AddTask(ctx context.Context, task models.Task) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.tasks[task.ID] = task
	return s.saveToFile()
}

func (s *MemoryStorage) GetTasks(ctx context.Context) ([]models.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	tasks := make([]models.Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (s *MemoryStorage) saveToFile() error {
	file, err := os.Create(s.file)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(s.tasks)
}

func (s *MemoryStorage) loadFromFile() error {
	if _, err := os.Stat(s.file); os.IsNotExist(err) {
		return nil
	}

	file, err := os.Open(s.file)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(&s.tasks)
}