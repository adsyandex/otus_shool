package repository

import (
	"fmt"
	"sync"

	"github.com/adsyandex/otus_shool/internal/model"

)

var (
	tasks []model.Task
	mu    sync.Mutex
)

// AddEntity принимает интерфейс и распределяет объект в нужный слайс
func AddEntity(e interface{}) {
	mu.Lock()
	defer mu.Unlock()

	switch v := e.(type) {
	case model.Task:
		tasks = append(tasks, v)
		fmt.Println("Добавлена задача:", v)
	default:
		fmt.Println("Неизвестный тип данных, не добавлено")
	}
}

// GetTasks возвращает список задач
func GetTasks() []model.Task {
	mu.Lock()
	defer mu.Unlock()
	return tasks
}
