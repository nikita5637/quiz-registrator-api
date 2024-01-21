package mysql

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-xorm/builder"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/logger"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// Log represents a row from 'logs'.
type Log struct {
	ID         int            `json:"id"`          // id
	Timestamp  time.Time      `json:"timestamp"`   // timestamp
	UserID     sql.NullInt64  `json:"user_id"`     // user_id
	ActionID   int            `json:"action_id"`   // action_id
	MessageID  int            `json:"message_id"`  // message_id
	ObjectType sql.NullString `json:"object_type"` // object_type
	ObjectID   sql.NullInt64  `json:"object_id"`   // object_id
	Metadata   sql.NullString `json:"metadata"`    // metadata
}

// LogStorage is Log service implementation
type LogStorage struct {
	db *tx.Manager
}

// NewLogStorage creates new instance of LogStorage
func NewLogStorage(txManager *tx.Manager) *LogStorage {
	return &LogStorage{
		db: txManager,
	}
}

// GetAll returns all records
func (s *LogStorage) GetAll(ctx context.Context) ([]Log, error) {
	return s.Find(ctx, nil, "")
}

// Find perform find request by params
func (s *LogStorage) Find(ctx context.Context, q builder.Cond, sort string) ([]Log, error) {
	query := `SELECT id, timestamp, user_id, action_id, message_id, object_type, object_id, metadata FROM logs`

	var args []interface{}

	if q != nil {
		var where string
		var err error
		where, args, err = builder.ToSQL(q)
		if err != nil {
			return nil, err
		}
		query += ` WHERE ` + where
	}

	if sort != "" {
		query += ` ` + getOrderStmt(sort)
	}

	rows, err := s.db.Sync(ctx).QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Log

	for rows.Next() {
		var item Log
		if err := rows.Scan(
			&item.ID,
			&item.Timestamp,
			&item.UserID,
			&item.ActionID,
			&item.MessageID,
			&item.ObjectType,
			&item.ObjectID,
			&item.Metadata,
		); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

// FindWithLimit perform find request by params, offset and limit
func (s *LogStorage) FindWithLimit(ctx context.Context, q builder.Cond, sort string, offset, limit uint64) ([]Log, error) {
	query := `SELECT id, timestamp, user_id, action_id, message_id, object_type, object_id, metadata FROM logs`

	var args []interface{}

	if q != nil {
		var where string
		var err error
		where, args, err = builder.ToSQL(q)
		if err != nil {
			return nil, err
		}
		query += ` WHERE ` + where
	}

	if sort != "" {
		query += ` ` + getOrderStmt(sort)
	}

	if limit != 0 {
		query += ` LIMIT ? OFFSET ?`
		args = append(args, limit)
		args = append(args, offset)
	}

	rows, err := s.db.Sync(ctx).QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Log

	for rows.Next() {
		var item Log
		if err := rows.Scan(
			&item.ID,
			&item.Timestamp,
			&item.UserID,
			&item.ActionID,
			&item.MessageID,
			&item.ObjectType,
			&item.ObjectID,
			&item.Metadata,
		); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

// Total return count(*) by params
func (s *LogStorage) Total(ctx context.Context, q builder.Cond) (uint64, error) {
	query := `SELECT count(*) FROM logs`

	var args []interface{}

	if q != nil {
		var where string
		var err error
		where, args, err = builder.ToSQL(q)
		if err != nil {
			return 0, err
		}
		query += ` WHERE ` + where
	}

	rows, err := s.db.Sync(ctx).QueryContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var count uint64

	for rows.Next() {
		if err := rows.Scan(
			&count,
		); err != nil {
			return 0, err
		}
	}

	return count, nil
}

// Insert inserts the Log to the database.
func (s *LogStorage) Insert(ctx context.Context, item Log) (int, error) {
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO logs (` +
		`timestamp, user_id, action_id, message_id, object_type, object_id, metadata` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?` +
		`)`
	// run
	logger.Debugf(ctx, sqlstr, item.Timestamp, item.UserID, item.ActionID, item.MessageID, item.ObjectType, item.ObjectID, item.Metadata)

	res, err := s.db.Master(ctx).ExecContext(ctx, sqlstr, item.Timestamp, item.UserID, item.ActionID, item.MessageID, item.ObjectType, item.ObjectID, item.Metadata)
	if err != nil {
		return 0, err
	}
	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Update updates a Log in the database.
func (s *LogStorage) Update(ctx context.Context, item Log) error {
	// update with primary key
	const sqlstr = `UPDATE logs SET ` +
		`timestamp = ?, user_id = ?, action_id = ?, message_id = ?, object_type = ?, object_id = ?, metadata = ? ` +
		`WHERE id = ?`
	// run
	logger.Debugf(ctx, sqlstr, item.Timestamp, item.UserID, item.ActionID, item.MessageID, item.ObjectType, item.ObjectID, item.Metadata, item.ID)
	if _, err := s.db.Master(ctx).ExecContext(ctx, sqlstr, item.Timestamp, item.UserID, item.ActionID, item.MessageID, item.ObjectType, item.ObjectID, item.Metadata, item.ID); err != nil {
		return err
	}

	return nil
}

// Upsert performs an upsert for Log.
func (s *LogStorage) Upsert(ctx context.Context, item Log) error {
	// upsert
	const sqlstr = `INSERT INTO logs (` +
		`id, timestamp, user_id, action_id, message_id, object_type, object_id, metadata` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`timestamp = VALUES(timestamp), user_id = VALUES(user_id), action_id = VALUES(action_id), message_id = VALUES(message_id), object_type = VALUES(object_type), object_id = VALUES(object_id), metadata = VALUES(metadata)`
	// run
	logger.Debugf(ctx, sqlstr, item.ID, item.Timestamp, item.UserID, item.ActionID, item.MessageID, item.ObjectType, item.ObjectID, item.Metadata)
	if _, err := s.db.Master(ctx).ExecContext(ctx, sqlstr, item.ID, item.Timestamp, item.UserID, item.ActionID, item.MessageID, item.ObjectType, item.ObjectID, item.Metadata); err != nil {
		return err
	}

	return nil
}

// Delete deletes the Log from the database.
func (s *LogStorage) Delete(ctx context.Context, id int) error {
	// delete with single primary key
	const sqlstr = `DELETE FROM logs ` +
		`WHERE id = ?`
	// run
	logger.Debugf(ctx, sqlstr, id)

	if _, err := s.db.Master(ctx).ExecContext(ctx, sqlstr, id); err != nil {
		return err
	}

	return nil
}

// GetLogByID retrieves a row from 'logs' as a Log.
//
// Generated from index 'logs_id_pkey'.
func (s *LogStorage) GetLogByID(ctx context.Context, id int) (*Log, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, timestamp, user_id, action_id, message_id, object_type, object_id, metadata ` +
		`FROM logs ` +
		`WHERE id = ?`
	// run
	logger.Debugf(ctx, sqlstr, id)
	l := Log{}
	if err := s.db.Sync(ctx).QueryRowContext(ctx, sqlstr, id).Scan(&l.ID, &l.Timestamp, &l.UserID, &l.ActionID, &l.MessageID, &l.ObjectType, &l.ObjectID, &l.Metadata); err != nil {
		return nil, err
	}
	return &l, nil
}
