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

	// Загрузка задач из хранилища
	tasks, err := taskManager.GetTasks()
	if err != nil {
			log.Fatalf("Ошибка загрузки задач: %v", err)
	}
	for _, t := range tasks {
			taskChannel <- t
	}
		
		

	// Запуск API-сервера
	r := gin.Default()
	api.SetupRouter(r, taskManager)
	fmt.Println("Сервер запущен на http://localhost:8080")
	r.Run(":8080")

	// Закрываем канал задач и ждем завершения горутин
	close(taskChannel)
	wg.Wait()
}
