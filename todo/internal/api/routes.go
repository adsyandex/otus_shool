package api

import (
	"github.com/adsyandex/otus_shool/internal/task"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, taskManager *task.TaskManager) {
	taskHandler := NewTaskHandler(taskManager)

	r.GET("/tasks", taskHandler.GetTasks)
	r.GET("/tasks/:id", taskHandler.GetTaskByID)
	r.POST("/tasks", taskHandler.CreateTask)
	r.PUT("/tasks/:id", taskHandler.UpdateTask)
	r.DELETE("/tasks/:id", taskHandler.DeleteTask)
}
