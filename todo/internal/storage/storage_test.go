package storage

import (
	"context"
	"testing"
	"os"

	"github.com/adsyandex/otus_shool/todo/internal/models"
)

func TestCSVStorage(t *testing.T) {
	ctx := context.Background()
	
	// Используем временный файл
	testFile := "testdata/temp_tasks.csv"
	
	store := NewCSVStorage(testFile)
	defer func() {
		// Очистка после тестов
		_ = os.Remove(testFile)
	}()

	t.Run("Create and Get Task", func(t *testing.T) {
		task := &models.Task{
			Title:       "Test Task",
			Description: "Test Description",
		}

		// Тест CreateTask с контекстом
		if err := store.CreateTask(ctx, task); err != nil {
			t.Fatalf("CreateTask failed: %v", err)
		}

		// Тест GetTask с контекстом
		found, err := store.GetTask(ctx, task.ID)
		if err != nil {
			t.Fatalf("GetTask failed: %v", err)
		}
		if found.Title != task.Title {
			t.Errorf("Expected title %q, got %q", task.Title, found.Title)
		}
	})

	t.Run("List Tasks", func(t *testing.T) {
		// Тест ListTasks с контекстом
		tasks, err := store.ListTasks(ctx)
		if err != nil {
			t.Fatalf("ListTasks failed: %v", err)
		}
		if len(tasks) == 0 {
			t.Error("Expected at least one task, got 0")
		}
	})
}