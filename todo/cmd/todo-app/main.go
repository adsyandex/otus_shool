package main

import (
    "fmt"
    "os"
    "github.com/gin-gonic/gin"
    "github.com/adsyandex/otus_shool/internal/api" // Исправленный импорт
    "github.com/adsyandex/otus_shool/internal/menu"
    "github.com/adsyandex/otus_shool/internal/storage"
    "github.com/adsyandex/otus_shool/internal/task"
)

func main() {
    // Инициализация хранилища
    store := storage.NewFileStorage("tasks.json")

    // Инициализация менеджера задач
    taskManager := task.NewTaskManager(store)

    // Выбор режима работы
    if len(os.Args) > 1 && os.Args[1] == "--api" {
        // Сетевой режим
        r := gin.Default()
        api.SetupRoutes(r, taskManager)
        fmt.Println("Сервер запущен на http://localhost:8080")
        r.Run(":8080")
    } else {
        // Консольный режим
        appMenu := menu.NewMenu(taskManager)
        fmt.Println("Добро пожаловать в TODO-лист!")
        appMenu.Run()
    }
}