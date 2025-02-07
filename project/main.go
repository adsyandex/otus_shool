package main

import (
	"fmt"
	"log"
	"time"

	"example.com/otus_shool/project/internal/model/task"
	"example.com/otus_shool/project/internal/storage"
)

func main() {
	// Создаем список задач
	tasks := []task.Task{
		{ID: 1, Title: "Написать код", Description: "Реализовать структуры", CreatedAt: time.Now()},
		{ID: 2, Title: "Проверить код", Description: "Протестировать функционал", CreatedAt: time.Now()},
	}

	// Инициализируем файловое хранилище
	storage := storage.NewFileStorage("tasks.json")

	// Сохраняем задачи в файл
	err := storage.SaveTasks(tasks)
	if err != nil {
		log.Fatalf("Ошибка при сохранении задач: %v", err)
	}

	// Загружаем задачи из файла
	loadedTasks, err := storage.LoadTasks()
	if err != nil {
		log.Fatalf("Ошибка при загрузке задач: %v", err)
	}

	// Вывод загруженных задач
	fmt.Println("Загруженные задачи:")
	for _, t := range loadedTasks {
		fmt.Printf("- [%d] %s: %s\n", t.ID, t.Title, t.Description)
	}
}
