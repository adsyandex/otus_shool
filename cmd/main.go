package main

import (
	"fmt"
	"time"

	"project-root/internal/repository"
	"project-root/internal/service"
)

func main() {
	fmt.Println("Запуск генерации данных...")
	service.StartDataGeneration()

	// Ожидание, пока накопятся данные
	time.Sleep(10 * time.Second)

	fmt.Println("\nСписок пользователей:")
	for _, user := range repository.GetUsers() {
		fmt.Println(user)
	}

	fmt.Println("\nСписок продуктов:")
	for _, product := range repository.GetProducts() {
		fmt.Println(product)
	}
}
