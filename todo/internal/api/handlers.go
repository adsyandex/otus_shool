// internal/api/handlers.go
package api

import (
    "encoding/json"
    "net/http"
    "github.com/adsyandex/otus_shool/todo/internal/models"
    "github.com/adsyandex/otus_shool/todo/internal/task"
)

type TaskHandler struct {
    service *task.Service // Исправлено с TaskManager на Service
}

func (h *TaskHandler) AddTask(w http.ResponseWriter, r *http.Request) {
    var req models.TaskRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }
    
    resp, err := h.service.AddTask(r.Context(), req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    json.NewEncoder(w).Encode(resp)
}

// internal/api/routes.go
package api

import (
    "github.com/gorilla/mux"
    "github.com/adsyandex/otus_shool/todo/internal/task"
)

func RegisterRoutes(r *mux.Router, s *task.Service) {
    handler := &TaskHandler{service: s}
    r.HandleFunc("/tasks", handler.AddTask).Methods("POST")
}