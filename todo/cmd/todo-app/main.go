package main

import (
	"github.com/adsyandex/otus_shool/todo/internal/api"
	"github.com/adsyandex/otus_shool/todo/internal/storage"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	store := storage.NewCSVStorage("data/tasks.csv")

	router := gin.Default()
	api.SetupRoutes(router, store)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
