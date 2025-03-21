// internal/models/task.go
package models // Исправлено с storage на models

type Task struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

// internal/models/requests.go
package models

type TaskRequest struct {
    Name string `json:"name"`
}