package api

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
    "github.com/adsyandex/otus_shool/internal/task"
)

type TaskHandler struct {
    TaskManager *task.TaskManager
}

func NewTaskHandler(taskManager *task.TaskManager) *TaskHandler {
    return &TaskHandler{TaskManager: taskManager}
}

func (h *TaskHandler) GetTasks(c *gin.Context) {
    tasks, err := h.TaskManager.GetTasks()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) GetTaskByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID задачи"})
        return
    }

    task, err := h.TaskManager.GetTaskByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
    var newTask struct {
        Title       string `json:"title"`
        Description string `json:"description"`
    }
    if err := c.ShouldBindJSON(&newTask); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := h.TaskManager.AddTask(newTask.Title, newTask.Description)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"message": "Задача успешно создана"})
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID задачи"})
        return
    }

    var updatedTask struct {
        Title       string `json:"title"`
        Description string `json:"description"`
        Status      string `json:"status"`
    }
    if err := c.ShouldBindJSON(&updatedTask); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err = h.TaskManager.UpdateTask(id, updatedTask.Title, updatedTask.Description, updatedTask.Status)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Задача успешно обновлена"})
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID задачи"})
        return
    }

    err = h.TaskManager.DeleteTask(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Задача успешно удалена"})
}