package repository

import (
	"fmt"
	"sync"

	"github.com/adsyandex/otus_shool/cmd/internal/model"

)

var (
	tasks []model.Task
	mu    sync.Mutex
)

// AddTask добавляет задачу в хранилище
func AddTask(task model.Task) {
	mu.Lock()
	defer mu.Unlock()
	tasks = append(tasks, task)
	fmt.Println("Добавлена задача:", task)
}

// GetTasks возвращает список задач
func GetTasks() []model.Task {
	mu.Lock()
	defer mu.Unlock()
	return tasks
}
