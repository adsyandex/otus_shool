package service

import (
	"fmt"
	"time"

	"github.com/adsyandex/otus_shool/internal/model"
	"github.com/adsyandex/otus_shool/internal/repository"
)

// StartTaskGeneration создаёт задачи в основном потоке
func StartTaskGeneration() {
	id := 1
	for {
		task := model.Task{
			ID:     id,
			Title:  fmt.Sprintf("Задача %d", id),
			Status: "в процессе",
		}

		repository.AddEntity(task) // Теперь передаём интерфейс

		id++
		time.Sleep(2 * time.Second) // Пауза перед созданием следующей задачи
	}
}