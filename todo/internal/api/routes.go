package api

import (
	"github.com/gin-gonic/gin"
	"github.com/adsyandex/otus_shool/todo/internal/storage"
)

func SetupRoutes(router *gin.Engine, storage storage.Storage) {
	handler := NewTaskHandler(storage)

	// Убрана группа /api, теперь маршруты будут доступны от корня
	router.GET("/items", handler.GetAllTasks)
	router.POST("/item", handler.CreateTask)
	router.GET("/item/:id", handler.GetTask)
	router.PUT("/item/:id", handler.UpdateTask)
	router.DELETE("/item/:id", handler.DeleteTask)
}