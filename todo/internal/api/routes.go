package api

import (
	"github.com/adsyandex/otus_shool/todo/internal/task"

	"github.com/gin-gonic/gin"
)

// SetupRouter настраивает маршруты для API
func SetupRouter(r *gin.Engine, tm *task.TaskManager) {
    r.GET("/tasks", func(c *gin.Context) {
        tasks, err := tm.GetTasks()
        if err != nil {
            c.JSON(500, gin.H{"error": err.Error()})
            return
        }
        c.JSON(200, tasks)
    })
}
