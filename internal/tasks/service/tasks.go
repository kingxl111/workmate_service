package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/kingxl111/workmate_service/internal/storage/postgres"
	"time"
)

type taskService struct {
	repository TaskRepository
}

func NewTaskService(repository TaskRepository) *taskService {
	return &taskService{
		repository: repository,
	}
}

func (t *taskService) Create(ctx context.Context, task *Task) error {
	tsk := postgres.Task{
		ID:        task.ID,
		Type:      task.Type,
		Status:    string(task.Status),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	id, err := t.repository.Create(ctx, &tsk)
	if err != nil {
		return err
	}

	task.ID = id

	return nil
}

func (t *taskService) GetByID(ctx context.Context, id uuid.UUID) (Task, error) {
	return Task{}, nil
}
