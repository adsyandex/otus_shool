package task

import (
	"sync"
)

// Task представляет задачу
type Task struct {
	ID   int
	Name string
}

// TaskManager управляет списком задач
type TaskManager struct {
	tasks []Task
	mu    sync.Mutex
}

// NewTaskManager создает новый менеджер задач
func NewTaskManager() *TaskManager {
	return &TaskManager{}
}

// AddTask добавляет задачу в список
func (tm *TaskManager) AddTask(task Task) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	tm.tasks = append(tm.tasks, task)
}

// GetTasks возвращает копию списка задач
func (tm *TaskManager) GetTasks() []Task {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	return append([]Task{}, tm.tasks...)
}
