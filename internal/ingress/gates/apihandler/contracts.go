package apihandler

import (
	"context"
	"github.com/google/uuid"
	"github.com/kingxl111/workmate_service/internal/tasks/service"
)

//go:generate mockgen -source=contracts.go -destination=mocks.go -package=apihandler
type TaskService interface {
	Create(ctx context.Context, task *service.Task) error
	GetByID(ctx context.Context, id uuid.UUID) (*service.Task, error)
}
