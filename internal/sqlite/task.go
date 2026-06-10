package sqlite

import (
	"context"
	"database/sql"

	"example.com/todo_tui/internal/domain"
	"example.com/todo_tui/internal/repository"
)

type sqliteTaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) repository.TaskRepository {
	return &sqliteTaskRepository{db: db}
}

func (r *sqliteTaskRepository) List(ctx context.Context) ([]domain.Task, error) {
	rows, err := r.db.Query("SELECT id, title, description, due_date, priority, is_completed, created_at, updated_at FROM Tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []domain.Task

	for rows.Next() {
		var task domain.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Priority, &task.IsCompleted, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *sqliteTaskRepository) GetByID(ctx context.Context, id string) (domain.Task, error) {
	panic("not implemented")
}

func (r *sqliteTaskRepository) Create(ctx context.Context, newTask domain.Task) (domain.Task, error) {
	panic("not implemented")
}

func (r *sqliteTaskRepository) Update(ctx context.Context, targetTask domain.Task) (domain.Task, error) {
	panic("not implemented")
}

func (r *sqliteTaskRepository) Delete(ctx context.Context, id string) error {
	panic("not implemented")
}
