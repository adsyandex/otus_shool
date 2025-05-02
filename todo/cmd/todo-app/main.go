package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv" // Добавьте этот импорт
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/adsyandex/otus_shool/todo/docs"
	"github.com/adsyandex/otus_shool/todo/internal/api"
	"github.com/adsyandex/otus_shool/todo/internal/storage"
)

// @title Todo App API
// @version 1.0
// @description This is a sample todo server with JWT authentication
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
	// 1. Загрузка переменных окружения
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	// 2. Инициализация хранилища
	store := redis_logger.NewCSVStorage("data/tasks.csv")

	// 3. Создание роутера с CORS
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 4. Настройка маршрутов
	api.SetupRoutes(router, store)

	// 5. Запуск сервера
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server starting on :%s", port)
	log.Println("\n=== Проверка переменных окружения ===")
log.Println("PORT:", os.Getenv("PORT"))
if secret := os.Getenv("JWT_SECRET"); secret != "" {
    log.Println("JWT_SECRET: [скрыто] Длина:", len(secret))
} else {
    log.Println("JWT_SECRET: не установлен!")
}

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}