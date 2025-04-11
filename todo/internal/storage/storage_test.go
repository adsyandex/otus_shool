package storage

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/adsyandex/otus_shool/todo/internal/models"
)

func TestCSVStorage(t *testing.T) {
	ctx := context.Background()
	
	// Создаем testdata если не существует
	testDir := filepath.Join("testdata")
	if _, err := os.Stat(testDir); os.IsNotExist(err) {
		if err := os.Mkdir(testDir, 0755); err != nil {
			t.Fatalf("Failed to create testdata dir: %v", err)
		}
	}

	testFile := filepath.Join(testDir, "temp_tasks.csv")
	defer os.Remove(testFile) // Очистка после тестов

	store := NewCSVStorage(testFile)

	t.Run("Create and Get Task", func(t *testing.T) {
		task := &models.Task{
			Title:       "Test Task",
			Description: "Test Description",
		}

		if err := store.CreateTask(ctx, task); err != nil {
			t.Fatalf("CreateTask failed: %v", err)
		}

		found, err := store.GetTask(ctx, task.ID)
		if err != nil {
			t.Fatalf("GetTask failed: %v", err)
		}
		if found.Title != task.Title {
			t.Errorf("Expected title %q, got %q", task.Title, found.Title)
		}
	})

	t.Run("List Tasks", func(t *testing.T) {
		tasks, err := store.ListTasks(ctx)
		if err != nil {
			t.Fatalf("ListTasks failed: %v", err)
		}
		if len(tasks) == 0 {
			t.Error("Expected at least one task, got 0")
		}
	})
}