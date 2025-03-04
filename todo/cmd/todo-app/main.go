package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
    	"github.com/adsyandex/otus_shool/todo/internal/api"
    	"github.com/adsyandex/otus_shool/todo/internal/logger"
    	"github.com/adsyandex/otus_shool/todo/internal/storage"
    	"github.com/adsyandex/otus_shool/todo/internal/task"
	"github.com/adsyandex/otus_shool/todo/internal/models"
)

// Глобальные переменные для синхронизации
var once sync.Once

func main() {
	// Инициализация хранилища
	store := storage.NewFileStorage("tasks.json")

	// Инициализация менеджера задач
	taskManager := task.NewTaskManager(store)

	// Канал для передачи задач между горутинами
	taskChannel := make(chan models.Task, 10)
	var wg sync.WaitGroup

	// Горутина для обработки задач
	wg.Add(1)
	go func() {
		defer wg.Done()
		for t := range taskChannel {
			log.Println("Обрабатываем задачу:", t.Name)
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
		tasks, err := taskManager.GetTasks()
		if err != nil {
			log.Printf("Ошибка загрузки задач: %v", err)
			// Добавляем начальные задачи, если список пуст
			initialTasks := []models.Task{
				{ID: 1, Name: "Первая задача"},
				{ID: 2, Name: "Вторая задача"},
			}
			for _, t := range initialTasks {
				taskManager.AddTask(t)
				taskChannel <- t
			}
			return
		}
		for _, t := range tasks {
			taskChannel <- t
		}
	})

	// Запуск API-сервера
	r := gin.Default()
	api.SetupRouter(r, taskManager)
	fmt.Println("Сервер запущен на http://localhost:8080")
	r.Run(":8080")

	// Закрываем канал задач и ждем завершения горутин
	close(taskChannel)
	wg.Wait()
}
