package api

import (
	"github.com/adsyandex/otus_shool/internal/task"

	"github.com/gin-gonic/gin"
)

// SetupRouter настраивает маршруты API
func SetupRouter(taskManager *task.TaskManager) *gin.Engine {
	router := gin.Default()
	handler := NewTaskHandler(taskManager)

	router.POST("/tasks", handler.AddTask)
	router.GET("/tasks", handler.GetTasks)

	return router
}
