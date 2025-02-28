package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/adsyandex/otus_shool/internal/api"
	"github.com/adsyandex/otus_shool/internal/logger"
	"github.com/adsyandex/otus_shool/internal/storage"
	"github.com/adsyandex/otus_shool/internal/task"
)

// Глобальные переменные для синхронизации
var once sync.Once

func main() {
	// Инициализация хранилища
	store := storage.NewFileStorage("tasks.json")

	// Инициализация менеджера задач
	taskManager := task.NewTaskManager(store)

	// Канал для передачи задач между горутинами
	taskChannel := make(chan task.Task, 10)
	var wg sync.WaitGroup

	// Горутина для обработки задач
	wg.Add(1)
	go func() {
		defer wg.Done()
		for t := range taskChannel {
			log.Println("Обрабатываем задачу:", t.Title)
			time.Sleep(500 * time.Millisecond) // Имитация работы
		}
	}()

	// Инициализация логирования
	consoleLogger := &logger.ConsoleLogger{}

	// Запуск логирования в отдельной горутине
	wg.Add(1)
	go func() {
		defer wg.Done()
		logger.StartLogging(taskManager, consoleLogger)
	}()

	// Используем sync.Once для загрузки задач из хранилища один раз
	once.Do(func() {
		tasks, err := store.GetTasks()
		if err != nil {
			log.Fatalf("Ошибка загрузки задач: %v", err)
		}
		for _, t := range tasks {
			taskChannel <- t
		}
	})

	// Запуск API-сервера
	r := gin.Default()
	api.SetupRoutes(r, taskManager)
	fmt.Println("Сервер запущен на http://localhost:8080")
	r.Run(":8080")

	// Закрываем канал задач и ждем завершения горутин
	close(taskChannel)
	wg.Wait()
}
