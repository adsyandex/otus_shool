// internal/api/routes.go
package api

import (
    "net/http"
    "github.com/adsyandex/otus_shool/todo/internal/task"
)

func RegisterRoutes(service *task.Service) http.Handler {
    mux := http.NewServeMux()
    handler := &TaskHandler{service: service}

    mux.HandleFunc("GET /tasks", handler.GetTasks)
    mux.HandleFunc("POST /tasks", handler.AddTask)

    return mux
}