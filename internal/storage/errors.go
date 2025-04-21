package storage

import "github.com/go-faster/errors"

var (
	ErrorBuildingInsertQuery = errors.New("error inserting record: can't build query")
	ErrorExecutingQuery      = errors.New("error executing query: can't insert record")
	ErrorScanningQuery       = errors.New("error scanning query: can't scan record")
	ErrNotFound              = errors.New("record not found")
	ErrorBuildingSelectQuery = errors.New("error get query: can't build select query")
)
