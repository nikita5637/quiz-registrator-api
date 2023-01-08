package tx

import (
	"context"
	"sync"
	"time"

	"database/sql"
)

// Client - database client facade
type Client interface {
	Exec(query string, args ...any) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	Prepare(query string) (*sql.Stmt, error)
}

type txKey struct{}

type txValue struct {
	name           string
	beginTs        time.Time
	mtx            sync.RWMutex
	sqlTransaction Client
}

// TX - transaction implementation
type TX struct {
	commitFn   func() error
	rollbackFn func()

	mtx  sync.Mutex
	done bool
}

var noopTX = &TX{}

// Commit - commit transaction if called in transaction context
func (tx *TX) Commit() error {
	tx.mtx.Lock()
	defer func() {
		tx.done = true
		tx.mtx.Unlock()
	}()

	if tx.commitFn == nil {
		return nil
	}

	return tx.commitFn()
}

// Rollback - rollback transaction if called in transaction context, safe for call any times
func (tx *TX) Rollback() {
	if tx.rollbackFn == nil {
		return
	}

	tx.mtx.Lock()
	defer func() {
		tx.done = true
		tx.mtx.Unlock()
	}()

	if tx.done {
		return
	}

	tx.rollbackFn()
}
