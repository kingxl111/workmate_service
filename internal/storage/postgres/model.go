package postgres

import (
	"github.com/google/uuid"
	"time"
)

type Task struct {
	ID        uuid.UUID `db:"id"`
	Type      string    `db:"type"`
	Status    string    `db:"status"`
	Result    string    `db:"result"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
