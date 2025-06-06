package models

import "time"

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DueDate     time.Time `json:"due_date,omitempty"`
	Priority    int       `json:"priority,omitempty"`
}

type TaskFilter struct {
	Completed *bool `json:"completed,omitempty"`
	Priority  *int  `json:"priority,omitempty"`
}
