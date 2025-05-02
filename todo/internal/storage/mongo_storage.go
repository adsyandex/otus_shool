package storage

import (
	"context"
	"os"
	"testing"
	
	"github.com/adsyandex/otus_shool/todo/internal/models"
	"github.com/adsyandex/otus_shool/todo/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestCSVStorage(t *testing.T) {
	// Создаем временный файл
	tmpFile, err := os.CreateTemp("", "test-*.csv")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	store := redis_logger.NewCSVStorage(tmpFile.Name())

	ctx := context.Background()

	t.Run("Create and Get", func(t *testing.T) {
		task := models.Task{Title: "Test task", Completed: false}
		
		// Create
		created, err := store.SaveTask(ctx, task)
		assert.NoError(t, err)
		assert.NotZero(t, created.ID)
		
		// Get
		found, err := store.GetTaskByID(ctx, created.ID)
		assert.NoError(t, err)
		assert.Equal(t, created.ID, found.ID)
		assert.Equal(t, "Test task", found.Title)
	})

	t.Run("Update", func(t *testing.T) {
		task := models.Task{Title: "Update test", Completed: true}
		created, _ := store.SaveTask(ctx, task)
		
		updatedTask := created
		updatedTask.Title = "Updated title"
		updated, err := store.UpdateTask(ctx, updatedTask)
		
		assert.NoError(t, err)
		assert.Equal(t, "Updated title", updated.Title)
	})

	t.Run("Delete", func(t *testing.T) {
		task := models.Task{Title: "To delete"}
		created, _ := store.SaveTask(ctx, task)
		
		err := store.DeleteTask(ctx, created.ID)
		assert.NoError(t, err)
		
		_, err = store.GetTaskByID(ctx, created.ID)
		assert.ErrorIs(t, err, redis_logger.ErrNotFound)
	})
}