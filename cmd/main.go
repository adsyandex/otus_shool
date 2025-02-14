package main

import (
	"fmt"
	"time"

	"github.com/adsyandex/otus_shool/internal/repository"
	"github.com/adsyandex/otus_shool/internal/service"

)


func main() {
	fmt.Println("Запуск генерации задач...")

	// Запускаем генерацию задач в основном потоке
	service.StartTaskGeneration()

	// Через 10 секунд выводим задачи
	time.Sleep(10 * time.Second)

	fmt.Println("\nСписок задач:")
	for _, task := range repository.GetTasks() {
		fmt.Println(task)
	}
}