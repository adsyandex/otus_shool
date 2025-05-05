package api

import (
	"net/http"

	"github.com/adsyandex/otus_shool/todo/internal/storage"
)

func SetupRoutes(router *http.ServeMux, storage storage.Storage, logger storage.Logger) {
	handler := NewHandler(storage, logger)
	
	router.HandleFunc("POST /tasks", handler.CreateTask)
	router.HandleFunc("GET /tasks/{id}", handler.GetTask)
	router.HandleFunc("PUT /tasks/{id}", handler.UpdateTask)
	router.HandleFunc("DELETE /tasks/{id}", handler.DeleteTask)
	router.HandleFunc("GET /tasks", handler.ListTasks)
}