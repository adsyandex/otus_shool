package api

import (
	"github.com/gin-gonic/gin"
	"github.com/adsyandex/otus_shool/todo/internal/storage"
)

func SetupRoutes(router *gin.Engine, storage storage.Storage) {
	handler := NewTaskHandler(storage)

	api := router.Group("/api")
	{
		api.GET("/items", handler.GetAllTasks)      // Изменил GetTasks на GetAllTasks
		api.POST("/item", handler.CreateTask)
		api.GET("/item/:id", handler.GetTask)      // Исправлено название метода
		api.PUT("/item/:id", handler.UpdateTask)   // Исправлено название метода
		api.DELETE("/item/:id", handler.DeleteTask) // Исправлено название метода
	}
}