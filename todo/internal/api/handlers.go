package api

import (
    "context"
    "encoding/json"
    "net/http"
    "time"
    
    "github.com/adsyandex/otus_shool/todo/internal/models"
    "github.com/adsyandex/otus_shool/todo/internal/task"
)

type TaskHandler struct {
    service *task.Service
}

func (h *TaskHandler) AddTask(w http.ResponseWriter, r *http.Request) {
    // Явное использование context (даже если просто логируем)
    ctx, cancel := context.WithCancel(r.Context())
    defer cancel()
    
    var req models.TaskRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    task := models.Task{
        Title:       req.Title,
        Description: req.Description,
        Completed:   false,
    }

    if err := h.service.AddTask(ctx, task); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
    ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second) // Явное использование
    defer cancel()
    
    tasks, err := h.service.GetTasks(ctx)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(models.TaskCollection{Tasks: tasks})
}