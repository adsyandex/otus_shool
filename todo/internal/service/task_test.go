package service

import (
	"context"
	"testing"
	"time"

	"github.com/adsyandex/otus_shool/todo/internal/models"
	"github.com/adsyandex/otus_shool/todo/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockStorage реализует мок для redis_logger.Storage
type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) GetTasks(ctx context.Context) ([]models.Task, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.Task), args.Error(1)
}

func (m *MockStorage) GetTaskByID(ctx context.Context, id int) (models.Task, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(models.Task), args.Error(1)
}

func (m *MockStorage) SaveTask(ctx context.Context, task models.Task) (models.Task, error) {
	args := m.Called(ctx, task)
	return args.Get(0).(models.Task), args.Error(1)
}

func (m *MockStorage) UpdateTask(ctx context.Context, task models.Task) (models.Task, error) {
	args := m.Called(ctx, task)
	return args.Get(0).(models.Task), args.Error(1)
}

func (m *MockStorage) DeleteTask(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// Тесты для TaskService
func TestTaskService_CreateTask(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	tests := []struct {
		name        string
		input       models.Task
		mockSetup   func(*MockStorage)
		expected    models.Task
		expectedErr string
	}{
		{
			name:  "success",
			input: models.Task{Title: "Valid task"},
			mockSetup: func(m *MockStorage) {
				m.On("SaveTask", ctx, models.Task{Title: "Valid task"}).
					Return(models.Task{ID: 1, Title: "Valid task"}, nil)
			},
			expected: models.Task{ID: 1, Title: "Valid task"},
		},
		{
			name:        "empty title",
			input:       models.Task{Title: ""},
			mockSetup:   func(m *MockStorage) {},
			expectedErr: "title cannot be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStorage := new(MockStorage)
			tt.mockSetup(mockStorage)

			service := NewTaskService(mockStorage)
			result, err := service.CreateTask(ctx, tt.input)

			if tt.expectedErr != "" {
				assert.Contains(t, err.Error(), tt.expectedErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
			mockStorage.AssertExpectations(t)
		})
	}
}

func TestTaskService_GetTask(t *testing.T) {
	ctx := context.Background()
	mockStorage := new(MockStorage)
	service := NewTaskService(mockStorage)

	// Успешный случай
	t.Run("success", func(t *testing.T) {
		mockStorage.On("GetTaskByID", ctx, 1).
			Return(models.Task{ID: 1, Title: "Existing task"}, nil)

		task, err := service.GetTask(ctx, 1)

		assert.NoError(t, err)
		assert.Equal(t, models.Task{ID: 1, Title: "Existing task"}, task)
	})

	// Задача не найдена
	t.Run("not found", func(t *testing.T) {
		mockStorage.On("GetTaskByID", ctx, 999).
			Return(models.Task{}, redis_logger.ErrNotFound)

		_, err := service.GetTask(ctx, 999)

		assert.ErrorIs(t, err, redis_logger.ErrNotFound)
	})

	mockStorage.AssertExpectations(t)
}

// Аналогичные тесты для остальных методов (GetAllTasks, UpdateTask, DeleteTask)
// Пример для UpdateTask:
func TestTaskService_UpdateTask(t *testing.T) {
	ctx := context.Background()
	mockStorage := new(MockStorage)
	service := NewTaskService(mockStorage)

	validTask := models.Task{ID: 1, Title: "Updated"}

	t.Run("success", func(t *testing.T) {
		mockStorage.On("UpdateTask", ctx, validTask).
			Return(validTask, nil)

		result, err := service.UpdateTask(ctx, validTask)

		assert.NoError(t, err)
		assert.Equal(t, validTask, result)
	})

	t.Run("invalid data", func(t *testing.T) {
		_, err := service.UpdateTask(ctx, models.Task{})
		assert.Error(t, err)
	})

	mockStorage.AssertExpectations(t)
}