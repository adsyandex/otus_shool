package main

import (
	"fmt"
	"time"

	"github.com/adsyandex/otus_shool/cmd/internal/repository"
	"github.com/adsyandex/otus_shool/cmd/internal/service"

)

func main() {
	fmt.Println("Запуск генерации данных...")
	service.StartDataGeneration()

	// Запускаем функцию в отдельной горутине
	go service.StartDataGeneration()

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
