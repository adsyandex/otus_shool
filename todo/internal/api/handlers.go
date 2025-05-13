package api

import (
	//"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"
	"log"

	"github.com/adsyandex/otus_shool/todo/internal/models"
	"github.com/adsyandex/otus_shool/todo/internal/storage"
)

// Handler структура для обработчиков API
type Handler struct {
	storage storage.Storage
	logger  storage.Logger // Используем интерфейс Logger из storage
}

// NewHandler создает новый экземпляр Handler
func NewHandler(storage storage.Storage, logger storage.Logger) *Handler {
	return &Handler{
		storage: storage,
		logger:  logger,
	}
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    
    var task models.Task
    if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Основная операция
    if err := h.storage.SaveTask(ctx, task); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Логирование (отдельно обрабатываем ошибку, не влияя на ответ)
    if err := h.logger.LogAction(ctx, "task_created", 24*time.Hour); err != nil {
        log.Printf("Failed to log action: %v", err) // Логируем ошибку в серверные логи
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(task)
}

func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")

	// Исправленный вызов с контекстом
	task, err := h.storage.GetTask(ctx, id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) { // Используем storage.ErrNotFound
			http.Error(w, "task not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.logger.LogAction(ctx, "task_retrieved", time.Hour); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(task)
}

func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")

	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	task.ID = id

	// Исправленный вызов
	if err := h.storage.UpdateTask(ctx, task); err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			http.Error(w, "task not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.logger.LogAction(ctx, "task_updated", 24*time.Hour); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")

	// Исправленный вызов
	if err := h.storage.DeleteTask(ctx, id); err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			http.Error(w, "task not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.logger.LogAction(ctx, "task_deleted", 24*time.Hour); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) ListTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	completedStr := r.URL.Query().Get("completed")

	var completed *bool
	if completedStr != "" {
		val, err := strconv.ParseBool(completedStr)
		if err != nil {
			http.Error(w, "invalid completed value", http.StatusBadRequest)
			return
		}
		completed = &val
	}

	// Исправленный вызов
	tasks, err := h.storage.ListTasks(ctx, completed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.logger.LogAction(ctx, "tasks_listed", time.Hour); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tasks)
}