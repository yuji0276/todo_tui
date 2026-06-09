package repository

import (
	"context"

	"example.com/todo_tui/internal/domain"
)

type TaskRepository interface {
	List(ctx context.Context) ([]domain.Task, error)
	GetByID(ctx context.Context, id string) (domain.Task, error)
	Create(ctx context.Context, newTask domain.Task) (domain.Task, error)
	Update(ctx context.Context, targetTask domain.Task) (domain.Task, error)
	Delete(ctx context.Context) error
}
