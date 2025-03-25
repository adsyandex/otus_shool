// internal/models/task.go
package models

type Task struct {
    ID          string `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Completed   bool   `json:"completed"`
}

type TaskRequest struct {
    Title       string `json:"title"`
    Description string `json:"description"`
}

type TaskCollection struct {
    Tasks []Task `json:"tasks"`
}