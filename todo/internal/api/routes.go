package api

import (
    "github.com/gin-gonic/gin"
    "github.com/adsyandex/otus_shool/todo/internal/task"
)

func SetupRouter(r *gin.Engine, tm *task.TaskManager) {
    handler := &TaskHandler{TaskManager: tm}
    r.GET("/tasks", handler.GetTasks)
    r.POST("/tasks", handler.AddTask)
}