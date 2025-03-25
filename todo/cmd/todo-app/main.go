// cmd/todo-app/main.go
package main

import (
	"log"
	"net/http"
	"github.com/adsyandex/otus_shool/todo/internal/api"
	"github.com/adsyandex/otus_shool/todo/internal/storage"
	"github.com/adsyandex/otus_shool/todo/internal/task"
)

func main() {
	// Инициализация хранилища
	storage := storage.NewMemoryStorage()
	
	// Создание сервиса
	service := task.NewService(storage)
	
	// Настройка маршрутов
	router := api.RegisterRoutes(service)
	
	// Запуск сервера
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}