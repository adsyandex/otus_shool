// internal/api/handlers.go
package api

import (
    "context"
    "encoding/json"
    "net/http"
    "github.com/adsyandex/otus_shool/todo/internal/models"
    "github.com/adsyandex/otus_shool/todo/internal/task"
)

type TaskHandler struct {
    service *task.Service
}

func (h *TaskHandler) AddTask(w http.ResponseWriter, r *http.Request) {
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

    if err := h.service.AddTask(r.Context(), task); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
    tasks, err := h.service.GetTasks(r.Context())
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(models.TaskCollection{Tasks: tasks})
}