package main

import (
	"context"
	"log"
	"time"

	"github.com/adsyandex/otus_shool/todo/internal/models"
	"github.com/adsyandex/otus_shool/todo/internal/storage/mongo"
	"github.com/adsyandex/otus_shool/todo/internal/storage/redis"
)

func main() {
	ctx := context.Background()

	// Подключение к MongoDB
	mongoStorage, err := mongo.NewMongoStorage(
		"mongodb://root:example@localhost:27017",
		"todo",
		"tasks",
	)
	if err != nil {
		log.Fatal(err)
	}
	defer mongoStorage.Close(ctx)

	// Подключение к Redis
	redisLogger, err := redis.NewRedisLogger("localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	defer redisLogger.Close()

	// Пример сохранения задачи
	task := models.Task{
		Title: "Learn Go", 
		Done: false,
	}
	
	if err := mongoStorage.SaveTask(ctx, task); err != nil {
		log.Fatal(err)
	}

	// Логирование действия
	if err := redisLogger.LogAction(ctx, "task_created", 24*time.Hour); err != nil {
		log.Fatal(err)
	}
}