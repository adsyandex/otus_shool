package api

import (
	//"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/adsyandex/otus_shool/todo/internal/logger"
	"github.com/adsyandex/otus_shool/todo/internal/models"
	"github.com/adsyandex/otus_shool/todo/internal/service"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service *service.TaskService
	log     *logger.Logger
}

func NewHandler(service *service.TaskService, log *logger.Logger) *Handler {
	return &Handler{
		service: service,
		log:     log,
	}
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		h.log.Error("Failed to decode request body", "error", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	createdTask, err := h.service.CreateTask(r.Context(), task)
	if err != nil {
		h.log.Error("Failed to create task", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdTask)
}

func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Task ID is required", http.StatusBadRequest)
		return
	}

	task, err := h.service.GetTask(r.Context(), id)
	if err != nil {
		if err == service.ErrTaskNotFound {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}
		h.log.Error("Failed to get task", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Task ID is required", http.StatusBadRequest)
		return
	}

	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		h.log.Error("Failed to decode request body", "error", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	task.ID = id

	updatedTask, err := h.service.UpdateTask(r.Context(), task)
	if err != nil {
		if err == service.ErrTaskNotFound {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}
		h.log.Error("Failed to update task", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTask)
}

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Task ID is required", http.StatusBadRequest)
		return
	}

	err := h.service.DeleteTask(r.Context(), id)
	if err != nil {
		if err == service.ErrTaskNotFound {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}
		h.log.Error("Failed to delete task", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) ListTasks(w http.ResponseWriter, r *http.Request) {
	completedStr := r.URL.Query().Get("completed")
	priorityStr := r.URL.Query().Get("priority")

	filter := models.TaskFilter{}

	if completedStr != "" {
		completed, err := strconv.ParseBool(completedStr)
		if err == nil {
			filter.Completed = &completed
		}
	}

	if priorityStr != "" {
		priority, err := strconv.Atoi(priorityStr)
		if err == nil {
			filter.Priority = &priority
		}
	}

	tasks, err := h.service.ListTasks(r.Context(), filter)
	if err != nil {
		h.log.Error("Failed to list tasks", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
