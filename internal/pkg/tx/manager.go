package tx

import (
	"context"
	"database/sql"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/logger"
	timeutils "github.com/nikita5637/quiz-registrator-api/utils/time"
)

// Manager - transaction manager
type Manager struct {
	db *sql.DB
}

// NewManager - return new transaction manager
func NewManager(db *sql.DB) *Manager {
	return &Manager{
		db: db,
	}
}

// Async - returns async read database with fallback
func (m *Manager) Async(ctx context.Context) Client {
	return m.getDB(ctx)
}

// Close - implements io.Closer
func (m *Manager) Close() error {
	return m.db.Close()
}

// Master - returns write database
func (m *Manager) Master(ctx context.Context) Client {
	return m.getDB(ctx)
}

// Sync - returns sync read database with fallback
func (m *Manager) Sync(ctx context.Context) Client {
	return m.getDB(ctx)
}

// RunTX - run transaction
func (m *Manager) RunTX(ctx context.Context, name string, op func(ctx context.Context) error) error {
	ctx, tx, err := m.begin(ctx, name)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	err = op(ctx)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (m *Manager) begin(ctx context.Context, name string) (context.Context, *TX, error) {
	if m.isInTx(ctx) {
		return ctx, noopTX, nil
	}

	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, noopTX, err
	}
	txval := txValue{
		name:           name,
		beginTs:        timeutils.TimeNow(),
		sqlTransaction: tx,
	}

	txHolder := &TX{
		commitFn: tx.Commit,
		rollbackFn: func() {
			err := tx.Rollback()
			if err != nil {
				logger.Errorf(ctx, "rollback %v failed, error: %v", name, err)
			}
		},
	}

	return context.WithValue(ctx, txKey{}, &txval), txHolder, nil
}

func (m *Manager) getDB(ctx context.Context) Client {
	tx, ok := m.getTx(ctx)
	if !ok {
		return getDefaultDB(ctx, m.db)
	}

	tx.mtx.RLock()
	defer tx.mtx.RUnlock()
	if tx.sqlTransaction == nil {
		return getDefaultDB(ctx, m.db)
	}

	return tx.sqlTransaction
}

func (m *Manager) getTx(ctx context.Context) (res *txValue, ok bool) {
	return txFromContext(ctx)
}

func (m *Manager) isInTx(ctx context.Context) bool {
	tx, ok := m.getTx(ctx)
	if !ok {
		return false
	}

	tx.mtx.RLock()
	ok = tx.sqlTransaction != nil
	tx.mtx.RUnlock()

	return ok
}

func getDefaultDB(ctx context.Context, db *sql.DB) Client {
	return db
}

func txFromContext(ctx context.Context) (res *txValue, ok bool) {
	v := ctx.Value(txKey{})
	if v == nil {
		return
	}
	res, ok = ctx.Value(txKey{}).(*txValue)
	return
}
