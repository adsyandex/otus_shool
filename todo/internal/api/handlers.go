package api

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/adsyandex/otus_shool/todo/internal/models"
    "github.com/adsyandex/otus_shool/todo/internal/task"
)

type TaskHandler struct {
    TaskManager *task.TaskManager
}

// NewTaskHandler создает новый обработчик задач
func NewTaskHandler(tm *task.TaskManager) *TaskHandler {
    return &TaskHandler{TaskManager: tm}
}

// AddTask добавляет новую задачу
func (h *TaskHandler) AddTask(c *gin.Context) {
    var newTask models.Task
    if err := c.ShouldBindJSON(&newTask); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }
    h.TaskManager.AddTask(newTask)
    c.JSON(http.StatusOK, gin.H{"message": "Task added", "task": newTask})
}

// GetTasks возвращает список задач
func (h *TaskHandler) GetTasks(c *gin.Context) {
    tasks, err := h.TaskManager.GetTasks(c.Request.Context())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, tasks)
}