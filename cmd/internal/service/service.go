package service

import (
	"fmt"
	"time"

	"github.com/adsyandex/otus_shool/cmd/internal/model"
	"github.com/adsyandex/otus_shool/cmd/internal/repository"
)

// StartDataGeneration генерирует данные и передает их в репозиторий по таймеру
func StartDataGeneration() {
	go func() {
		id := 1
		for {
			user := model.User{ID: id, Name: fmt.Sprintf("User %d", id)}
			product := model.Product{ID: id, Title: fmt.Sprintf("Product %d", id), Price: float64(id) * 10.5}

			repository.AddEntity(user)
			repository.AddEntity(product)

			id++
			time.Sleep(2 * time.Second)
		}
	}()
}
