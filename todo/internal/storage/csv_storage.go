package storage

import (
	"encoding/csv"
	"os"
	"strconv"
	"sync"
	"time"
	"context"

	"github.com/adsyandex/otus_shool/todo/internal/models"
	"github.com/google/uuid"
)

type CSVStorage struct {
	mu     sync.Mutex
	file   string
	tasks  []*models.Task
	//nextID int
}

func NewCSVStorage(filename string) *CSVStorage {
	s := &CSVStorage{
		file:  filename,
		tasks: make([]*models.Task, 0),
	}
	s.loadFromFile()
	return s
}

func (s *CSVStorage) CreateTask(ctx context.Context, task *models.Task) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	task.ID = uuid.New().String() // Генерация UUID вместо числового ID
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	s.tasks = append(s.tasks, task)
	return s.saveToFile()
}

func (s *CSVStorage) GetTask(ctx context.Context,id string) (*models.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, task := range s.tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return nil, os.ErrNotExist
}

// Добавьте остальные методы (UpdateTask, DeleteTask, ListTasks) по аналогии
func (s *CSVStorage) UpdateTask(ctx context.Context, task *models.Task) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, t := range s.tasks {
		if t.ID == task.ID {
			s.tasks[i] = task
			return s.saveToFile()
		}
	}
	return os.ErrNotExist
}

func (s *CSVStorage) DeleteTask(ctx context.Context, id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, task := range s.tasks {
		if task.ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return s.saveToFile()
		}
	}
	return ErrNotFound
}

func (s *CSVStorage) ListTasks(ctx context.Context) ([]*models.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	tasks := make([]*models.Task, len(s.tasks))
	copy(tasks, s.tasks)
	return tasks, nil
}

func (s *CSVStorage) loadFromFile() {
	file, err := os.Open(s.file)
	if err != nil {
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	for _, record := range records {
		if len(record) < 5 {
			continue
		}

		completed, _ := strconv.ParseBool(record[3])
		createdAt, _ := time.Parse(time.RFC3339, record[4])
		updatedAt := createdAt
		if len(record) > 5 {
			updatedAt, _ = time.Parse(time.RFC3339, record[5])
		}

		s.tasks = append(s.tasks, &models.Task{
			ID:          record[0],
			Title:       record[1],
			Description: record[2],
			Completed:   completed,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
		})
	}
}

func (s *CSVStorage) saveToFile() error {
	file, err := os.Create(s.file)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, task := range s.tasks {
		record := []string{
			task.ID,
			task.Title,
			task.Description,
			strconv.FormatBool(task.Completed),
			task.CreatedAt.Format(time.RFC3339),
			task.UpdatedAt.Format(time.RFC3339),
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}
	return nil
}