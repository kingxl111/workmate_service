package service

import (
	"github.com/google/uuid"
	"time"
)

type Task struct {
	ID        uuid.UUID
	Type      string
	Status    TaskStatus
	Result    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TaskStatus string

const (
	TaskStatusPending    TaskStatus = "pending"
	TaskStatusProcessing TaskStatus = "processing"
	TaskStatusDone       TaskStatus = "done"
	TaskStatusFailed     TaskStatus = "failed"
)
