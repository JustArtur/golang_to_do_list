package types

import (
	"github.com/guregu/null/v5"
	"time"
)

type TaskPayload struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"body"`
	DueDate     null.Time `json:"due_date"`
}

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"body"`
	DueDate     null.Time `json:"due_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
