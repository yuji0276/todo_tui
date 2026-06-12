package domain

import "time"

type Task struct {
	ID          string
	Title       string
	Description string
	DueDate     *time.Time
	Priority    Priority
	Done        bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Priority int

const (
	PriorityNone Priority = 0
	PriorityHigh Priority = 1
	PriorityMid  Priority = 2
	PriorityLow  Priority = 3
)
