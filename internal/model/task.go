package model

// Entity - общий интерфейс для всех сущностей (например, задач)
type Entity interface {
	GetID() int
}

// Task - структура задачи
type Task struct {
	ID     int
	Title  string
	Status string
}

// Реализация метода GetID для Task
func (t Task) GetID() int {
	return t.ID
}
