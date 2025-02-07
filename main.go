package main

import (
	"fmt"
	"log"
	"time"

	"example.com/otus_shool/project/internal/model/account"
	"example.com/otus_shool/project/internal/model/task"
	"example.com/otus_shool/project/internal/storage"
)

func main() {
	// Создаем пользователя
	user := account.User{ID: 1, Username: "admin"}

	// Устанавливаем пароль
	err := user.SetPassword("securepassword")
	if err != nil {
		log.Fatal("Ошибка хеширования пароля")
	}

	// Проверяем пароль
	if user.CheckPassword("securepassword") {
		fmt.Println("Пароль верный")
	} else {
		fmt.Println("Неверный пароль")
	}

	// Создаем задачи и связываем с пользователем (по UserID)
	t1 := task.Task{ID: 1, UserID: user.ID, Title: "Написать код", Description: "Реализовать структуры"}
	t1.SetCreatedAt(time.Now())

	t2 := task.Task{ID: 2, UserID: user.ID, Title: "Проверить код", Description: "Протестировать функционал"}
	t2.SetCreatedAt(time.Now())

	tasks := []task.Task{t1, t2}

	// Инициализируем файловое хранилище
	storage := storage.NewFileStorage("tasks.json")

	// Сохраняем задачи в файл
	err = storage.SaveTasks(tasks)
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
		fmt.Printf("- [%d] Пользователь %d: %s - %s (Создана: %v)\n", t.ID, t.UserID, t.Title, t.Description, t.CreatedAt())
	}
}
