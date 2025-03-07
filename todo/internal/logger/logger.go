package logger

import (
    "context"
    "log"
    "time"
    "github.com/adsyandex/otus_shool/todo/internal/models"
    "github.com/adsyandex/otus_shool/todo/internal/task"
)

type Logger interface {
    LogTasks(tasks []models.Task)
}

type ConsoleLogger struct{}

func (cl *ConsoleLogger) LogTasks(tasks []models.Task) {
    log.Printf("Добавленные задачи: %+v\n", tasks)
}

func StartLogging(ctx context.Context, tm *task.TaskManager, logger Logger) {
    var lastCount int
    ticker := time.NewTicker(200 * time.Millisecond)
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            tasks, err := tm.GetTasks(ctx)
            if err != nil {
                log.Printf("Ошибка при получении задач: %v", err)
                continue
            }

            if tasks != nil && len(tasks) > lastCount {
                logger.LogTasks(tasks[lastCount:])
                lastCount = len(tasks)
            }
        case <-ctx.Done():
            log.Println("Завершение логирования")
            return
        }
    }
}