package main

import (
	_ "github.com/adsyandex/otus_shool/todo/docs"
	"github.com/adsyandex/otus_shool/todo/internal/api"
	"github.com/adsyandex/otus_shool/todo/internal/storage"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

// @title Todo App API
// @version 1.0
// @description This is a sample todo server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// Инициализация хранилища
	store := storage.NewCSVStorage("data/tasks.csv")  // Обратите внимание на CS"V"Storage
	
	// Создание роутера
	router := gin.Default()
	
	// Инициализация обработчиков
	api.SetupRoutes(router, store)
	
	// Добавляем Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	// Запуск сервера
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}