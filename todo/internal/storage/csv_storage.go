package storage

import (
	"context"
	"encoding/csv"
	"github.com/adsyandex/otus_shool/todo/internal/models"
	"os"
	"strconv"
	"time"
)

type CSVStorage struct {
	filePath string
}

func NewCSVStorage(filePath string) Storage {
	return &CSVStorage{filePath: filePath}
}

func (cs *CSVStorage) GetTasks(ctx context.Context) ([]models.Task, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		file, err := os.Open(cs.filePath)
		if os.IsNotExist(err) {
			return []models.Task{}, nil
		}
		if err != nil {
			return nil, err
		}
		defer file.Close()

		reader := csv.NewReader(file)
		records, err := reader.ReadAll()
		if err != nil {
			return nil, err
		}

		var tasks []models.Task
		for _, record := range records {
			id, _ := strconv.Atoi(record[0])
			createdAt, _ := time.Parse(time.RFC3339, record[4])
			updatedAt, _ := time.Parse(time.RFC3339, record[5])

			tasks = append(tasks, models.Task{
				ID:          id,
				Title:       record[1],
				Description: record[2],
				Status:      record[3],
				CreatedAt:   createdAt,
				UpdatedAt:   updatedAt,
			})
		}
		return tasks, nil
	}
}

func (cs *CSVStorage) GetTaskByID(ctx context.Context, id int) (models.Task, error) {
	tasks, err := cs.GetTasks(ctx)
	if err != nil {
		return models.Task{}, err
	}

	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return models.Task{}, ErrNotFound
}

func (cs *CSVStorage) SaveTask(ctx context.Context, task models.Task) error {
	tasks, err := cs.GetTasks(ctx)
	if err != nil {
		return err
	}

	tasks = append(tasks, task)
	return cs.saveAllTasks(ctx, tasks)
}

func (cs *CSVStorage) UpdateTask(ctx context.Context, updated models.Task) error {
	tasks, err := cs.GetTasks(ctx)
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == updated.ID {
			tasks[i] = updated
			return cs.saveAllTasks(ctx, tasks)
		}
	}
	return ErrNotFound
}

func (cs *CSVStorage) DeleteTask(ctx context.Context, id int) error {
	tasks, err := cs.GetTasks(ctx)
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return cs.saveAllTasks(ctx, tasks)
		}
	}
	return ErrNotFound
}

func (cs *CSVStorage) saveAllTasks(ctx context.Context, tasks []models.Task) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		file, err := os.Create(cs.filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		for _, task := range tasks {
			record := []string{
				strconv.Itoa(task.ID),
				task.Title,
				task.Description,
				task.Status,
				task.CreatedAt.Format(time.RFC3339),
				task.UpdatedAt.Format(time.RFC3339),
			}
			if err := writer.Write(record); err != nil {
				return err
			}
		}
		return nil
	}
}
