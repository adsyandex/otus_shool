package api

import (
	"net/http"
	"todo-app/internal/task"

	"github.com/gin-gonic/gin"
)

// TaskHandler отвечает за обработку задач
type TaskHandler struct {
	TaskManager *task.TaskManager
}

// NewTaskHandler создает новый обработчик задач
func NewTaskHandler(tm *task.TaskManager) *TaskHandler {
	return &TaskHandler{TaskManager: tm}
}

// AddTask добавляет новую задачу
func (h *TaskHandler) AddTask(c *gin.Context) {
	var newTask task.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	h.TaskManager.AddTask(newTask)
	c.JSON(http.StatusOK, gin.H{"message": "Task added", "task": newTask})
}

// GetTasks возвращает список задач
func (h *TaskHandler) GetTasks(c *gin.Context) {
	tasks := h.TaskManager.GetTasks()
	c.JSON(http.StatusOK, tasks)
}
