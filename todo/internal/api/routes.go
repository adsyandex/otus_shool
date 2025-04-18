package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/adsyandex/otus_shool/todo/internal/auth"
	"github.com/adsyandex/otus_shool/todo/internal/storage"
)

func SetupRoutes(router *gin.Engine, storage storage.Storage) {
	handler := NewTaskHandler(storage)
	
	// Перенаправление на Swagger UI
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	// Публичные routes
	router.POST("/login", handler.Login)

	// Защищенные routes
	authGroup := router.Group("/").Use(auth.AuthMiddleware())
	{
		authGroup.GET("/tasks", handler.GetAllTasks)
		authGroup.POST("/tasks", handler.CreateTask)
		authGroup.GET("/tasks/:id", handler.GetTask)
		authGroup.PUT("/tasks/:id", handler.UpdateTask)
		authGroup.DELETE("/tasks/:id", handler.DeleteTask)
	}

	// Swagger (оставить только здесь)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}