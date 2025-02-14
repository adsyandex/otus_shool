package main

import (
	"fmt"
	"time"

	"github.com/adsyandex/otus_shool/cmd/internal/repository"
	"github.com/adsyandex/otus_shool/cmd/internal/service"

)

func main() {
	fmt.Println("Запуск генерации задач...")

	// Запускаем генерацию задач в фоновом режиме
	go service.StartTaskGeneration()

	// Ожидаем 10 секунд, пока накопятся данные
	time.Sleep(10 * time.Second)

	fmt.Println("\nСписок задач:")
	for _, task := range repository.GetTasks() {
		fmt.Println(task)
	}
}
