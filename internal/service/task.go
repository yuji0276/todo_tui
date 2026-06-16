package service

import (
	"context"

	"example.com/todo_tui/internal/domain"
	"example.com/todo_tui/internal/repository"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) GetAllTasks(ctx context.Context) ([]domain.Task, error) {
	tasks, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *TaskService) CreateTask(ctx context.Context, newTask domain.Task) (domain.Task, error) {
	task, err := s.repo.Create(ctx, newTask)
	newTask.Done = false
	if err != nil {
		return domain.Task{}, err
	}
	return task, nil
}
