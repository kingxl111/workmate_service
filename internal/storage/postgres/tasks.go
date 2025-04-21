package postgres

import (
	"context"
	"github.com/go-faster/errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/kingxl111/workmate_service/internal/storage"

	sq "github.com/Masterminds/squirrel"
)

const (
	tasksTable = "tasks"

	idColumn        = "id"
	typeColumn      = "type"
	statusColumn    = "status"
	resultColumn    = "result"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type taskRepository struct {
	db *DB
}

func NewTaskRepository(db *DB) *taskRepository {
	return &taskRepository{db: db}
}

func (t *taskRepository) Create(ctx context.Context, task *Task) (uuid.UUID, error) {
	var id uuid.UUID

	builder := sq.Insert(tasksTable).
		Columns(
			idColumn,
			typeColumn,
			statusColumn,
			resultColumn,
			createdAtColumn,
			updatedAtColumn,
		).
		Values(
			task.ID,
			task.Type,
			task.Status,
			task.Result,
			task.CreatedAt,
			task.UpdatedAt,
		).
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return id, storage.ErrorBuildingInsertQuery
	}

	err = t.db.pool.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return id, storage.ErrNotFound
		}
		return id, storage.ErrorExecutingQuery
	}

	return id, nil
}

func (t *taskRepository) GetByID(ctx context.Context, id uuid.UUID) (Task, error) {
	builder := sq.Select(
		idColumn,
		typeColumn,
		statusColumn,
		resultColumn,
		createdAtColumn,
		updatedAtColumn).
		From(tasksTable).
		Where(sq.Eq{idColumn: id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return Task{}, storage.ErrorBuildingSelectQuery
	}

	var task Task
	err = t.db.pool.QueryRow(ctx, query, args...).Scan(
		&task.ID,
		&task.Type,
		&task.Status,
		&task.Result,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return Task{}, storage.ErrNotFound
		}
		return Task{}, storage.ErrorExecutingQuery
	}

	return task, nil
}
