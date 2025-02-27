package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/adsyandex/otus_shool/internal/api"
	"github.com/adsyandex/otus_shool/internal/logger"
	"github.com/adsyandex/otus_shool/internal/storage"
	"github.com/adsyandex/otus_shool/internal/task"
)

func main() {
	// Инициализация хранилища
	store := storage.NewFileStorage("tasks.json")

	// Инициализация менеджера задач
	taskManager := task.NewTaskManager(store)

	// Инициализация логирования
	consoleLogger := &logger.ConsoleLogger{}

	// Запуск логирования в горутине
	go logger.StartLogging(taskManager, consoleLogger)

	// Запуск API-сервера
	r := gin.Default()
	api.SetupRoutes(r, taskManager)
	fmt.Println("Сервер запущен на http://localhost:8080")
	r.Run(":8080")
}
