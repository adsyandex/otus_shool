package task

import "time"

type Task struct {
    ID          int       `json:"id"`
    UserID      int       `json:"user_id"` // Связь с пользователем
    Title       string    `json:"title"`
    Description string    `json:"description"`
    CreatedAt   time.Time  `json:"created_at"`
}

/*// Метод для получения времени создания
func (t *Task) NewCreatedAt() time.Time {
    return t.CreatedAt
}

// Метод для установки времени создания
func (t *Task) SetCreatedAt(ti time.Time) {
    t.CreatedAt = ti
}*/
