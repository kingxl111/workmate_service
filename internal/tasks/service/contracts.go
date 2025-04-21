package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/kingxl111/workmate_service/internal/storage/postgres"
)

//go:generate mockgen -source=contracts.go -destination=mocks.go -package=service
type TaskRepository interface {
	Create(ctx context.Context, task *postgres.Task) error
	GetByID(ctx context.Context, id uuid.UUID) (*postgres.Task, error)
}
