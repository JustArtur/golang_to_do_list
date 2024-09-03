package types

import (
	"github.com/guregu/null/v5"
	"time"
)

// TaskPayload represents the payload used for creating or updating task.
type TaskPayload struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"body"`
	DueDate     null.Time `json:"due_date"`
}

// Task represents the task model retrieved from the database.
type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"body"`
	DueDate     null.Time `json:"due_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
