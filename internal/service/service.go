package service

import (
	"fmt"
	"time"

	"github.com/adsyandex/otus_shool/cmd/internal/model"
	"github.com/adsyandex/otus_shool/cmd/internal/repository"
)

// StartTaskGeneration создаёт задачи и передаёт в репозиторий
func StartTaskGeneration() {
	go func() {
		id := 1
		for {
			task := model.Task{
				ID:     id,
				Title:  fmt.Sprintf("Задача %d", id),
				Status: "в процессе",
			}

			repository.AddTask(task)

			id++
			time.Sleep(2 * time.Second)
		}
	}()
}
