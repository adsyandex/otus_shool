package logger

import (
    "log"
    "time"
    "github.com/adsyandex/otus_shool/todo/internal/models" // Импортируем models
    "github.com/adsyandex/otus_shool/todo/internal/task" 
)

// Logger определяет интерфейс логирования
type Logger interface {
    LogTasks(tasks []models.Task) // Используем models.Task
}

// ConsoleLogger выводит лог в консоль
type ConsoleLogger struct{}

// LogTasks логирует новые задачи в консоль
func (cl *ConsoleLogger) LogTasks(tasks []models.Task) { // Используем models.Task
    log.Printf("Добавленные задачи: %+v\n", tasks)
}

// StartLogging запускает процесс логирования
func StartLogging(taskManager *task.TaskManager, logger Logger) {
    var lastCount int
    ticker := time.NewTicker(200 * time.Millisecond)
    defer ticker.Stop()

    for range ticker.C {
        tasks, err := taskManager.GetTasks()
        if err != nil {
            log.Printf("Ошибка при получении задач: %v", err)
            continue
        }

        if tasks != nil && len(tasks) > lastCount {
            logger.LogTasks(tasks[lastCount:])
            lastCount = len(tasks)
        }
    }
}
