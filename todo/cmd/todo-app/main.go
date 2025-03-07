package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "os"
    "os/signal"
    "sync"
    "syscall"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/adsyandex/otus_shool/todo/internal/api"
    "github.com/adsyandex/otus_shool/todo/internal/logger"
    "github.com/adsyandex/otus_shool/todo/internal/storage"
    "github.com/adsyandex/otus_shool/todo/internal/task"
    "github.com/adsyandex/otus_shool/todo/internal/models"
)

func main() {
    // Создаём контекст с отменой
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

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
        logger.StartLogging(ctx, taskManager, consoleLogger)
    }()

    // Загрузка задач из хранилища
    tasks, err := taskManager.GetTasks(ctx)
    if err != nil {
        log.Fatalf("Ошибка загрузки задач: %v", err)
    }
    for _, t := range tasks {
        taskChannel <- t
    }

    // Запуск API-сервера
    r := gin.Default()
    api.SetupRouter(r, taskManager)

    // HTTP-сервер
    srv := &http.Server{
        Addr:    ":8080",
        Handler: r,
    }

    // Запуск сервера в отдельной горутине
    go func() {
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("Ошибка запуска сервера: %v", err)
        }
    }()

    fmt.Println("Сервер запущен на http://localhost:8080")

    // Ожидание сигналов от ОС
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    log.Println("Получен сигнал завершения")

    // Отмена контекста для завершения горутин
    cancel()

    // Graceful shutdown сервера
    ctxShutdown, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancelShutdown()
    if err := srv.Shutdown(ctxShutdown); err != nil {
        log.Fatalf("Ошибка завершения сервера: %v", err)
    }

    // Закрываем канал задач
    close(taskChannel)

    // Ожидание завершения всех горутин
    wg.Wait()
    log.Println("Приложение завершено")
}