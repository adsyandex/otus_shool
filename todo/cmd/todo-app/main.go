package main

import (
	"log"
	"net/http" // Добавлен импорт
	"github.com/adsyandex/otus_shool/todo/internal/api"
	"github.com/adsyandex/otus_shool/todo/internal/storage"
	"github.com/adsyandex/otus_shool/todo/internal/task"
)

func main() {
	storage := storage.NewMemoryStorage("data/tasks.json")
	service := task.NewService(storage)
	router := api.RegisterRoutes(service)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}