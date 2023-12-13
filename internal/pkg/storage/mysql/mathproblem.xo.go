package mysql

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"

	"github.com/go-xorm/builder"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/logger"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// MathProblem represents a row from 'math_problem'.
type MathProblem struct {
	ID        int          `json:"id"`         // id
	FkGameID  int          `json:"fk_game_id"` // fk_game_id
	URL       string       `json:"url"`        // url
	CreatedAt sql.NullTime `json:"created_at"` // created_at
}

// MathProblemStorage is MathProblem service implementation
type MathProblemStorage struct {
	db *tx.Manager
}

// NewMathProblemStorage creates new instance of MathProblemStorage
func NewMathProblemStorage(txManager *tx.Manager) *MathProblemStorage {
	return &MathProblemStorage{
		db: txManager,
	}
}

// GetAll returns all records
func (s *MathProblemStorage) GetAll(ctx context.Context) ([]MathProblem, error) {
	return s.Find(ctx, nil, "")
}

// Find perform find request by params
func (s *MathProblemStorage) Find(ctx context.Context, q builder.Cond, sort string) ([]MathProblem, error) {
	query := `SELECT id, fk_game_id, url, created_at FROM math_problem`

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

	var items []MathProblem

	for rows.Next() {
		var item MathProblem
		if err := rows.Scan(
			&item.ID,
			&item.FkGameID,
			&item.URL,
			&item.CreatedAt,
		); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

// FindWithLimit perform find request by params, offset and limit
func (s *MathProblemStorage) FindWithLimit(ctx context.Context, q builder.Cond, sort string, offset, limit uint64) ([]MathProblem, error) {
	query := `SELECT id, fk_game_id, url, created_at FROM math_problem`

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

	var items []MathProblem

	for rows.Next() {
		var item MathProblem
		if err := rows.Scan(
			&item.ID,
			&item.FkGameID,
			&item.URL,
			&item.CreatedAt,
		); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

// Total return count(*) by params
func (s *MathProblemStorage) Total(ctx context.Context, q builder.Cond) (uint64, error) {
	query := `SELECT count(*) FROM math_problem`

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

// Insert inserts the MathProblem to the database.
func (s *MathProblemStorage) Insert(ctx context.Context, item MathProblem) (int, error) {
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO math_problem (` +
		`fk_game_id, url, created_at` +
		`) VALUES (` +
		`?, ?, UTC_TIMESTAMP()` +
		`)`
	// run
	logger.Debugf(ctx, sqlstr, item.FkGameID, item.URL)

	res, err := s.db.Master(ctx).ExecContext(ctx, sqlstr, item.FkGameID, item.URL)
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

// Update updates a MathProblem in the database.
func (s *MathProblemStorage) Update(ctx context.Context, item MathProblem) error {
	// update with primary key
	const sqlstr = `UPDATE math_problem SET ` +
		`fk_game_id = ?, url = ? ` +
		`WHERE id = ?`
	// run
	logger.Debugf(ctx, sqlstr, item.FkGameID, item.URL, item.ID)
	if _, err := s.db.Master(ctx).ExecContext(ctx, sqlstr, item.FkGameID, item.URL, item.ID); err != nil {
		return err
	}

	return nil
}

// Upsert performs an upsert for MathProblem.
func (s *MathProblemStorage) Upsert(ctx context.Context, item MathProblem) error {
	// upsert
	const sqlstr = `INSERT INTO math_problem (` +
		`id, fk_game_id, url, created_at` +
		`) VALUES (` +
		`?, ?, ?, UTC_TIMESTAMP()` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`fk_game_id = VALUES(fk_game_id), url = VALUES(url), created_at = VALUES(created_at)`
	// run
	logger.Debugf(ctx, sqlstr, item.ID, item.FkGameID, item.URL, item.CreatedAt)
	if _, err := s.db.Master(ctx).ExecContext(ctx, sqlstr, item.ID, item.FkGameID, item.URL); err != nil {
		return err
	}

	return nil
}

// Delete deletes the MathProblem from the database.
func (s *MathProblemStorage) Delete(ctx context.Context, id int) error {
	// delete with single primary key
	const sqlstr = `DELETE FROM math_problem ` +
		`WHERE id = ?`
	// run
	logger.Debugf(ctx, sqlstr, id)

	if _, err := s.db.Master(ctx).ExecContext(ctx, sqlstr, id); err != nil {
		return err
	}

	return nil
}

// GetMathProblemByFkGameID retrieves a row from 'math_problem' as a MathProblem.
//
// Generated from index 'fk_game_id'.
func (s *MathProblemStorage) GetMathProblemByFkGameID(ctx context.Context, fkGameID int) ([]*MathProblem, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, fk_game_id, url, created_at ` +
		`FROM math_problem ` +
		`WHERE fk_game_id = ?`
	// run
	logger.Debugf(ctx, sqlstr, fkGameID)
	rows, err := s.db.Sync(ctx).QueryContext(ctx, sqlstr, fkGameID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// process
	var res []*MathProblem
	for rows.Next() {
		mp := MathProblem{}
		// scan
		if err := rows.Scan(&mp.ID, &mp.FkGameID, &mp.URL, &mp.CreatedAt); err != nil {
			return nil, err
		}
		res = append(res, &mp)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

// GetMathProblemByID retrieves a row from 'math_problem' as a MathProblem.
//
// Generated from index 'math_problem_id_pkey'.
func (s *MathProblemStorage) GetMathProblemByID(ctx context.Context, id int) (*MathProblem, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, fk_game_id, url, created_at ` +
		`FROM math_problem ` +
		`WHERE id = ?`
	// run
	logger.Debugf(ctx, sqlstr, id)
	mp := MathProblem{}
	if err := s.db.Sync(ctx).QueryRowContext(ctx, sqlstr, id).Scan(&mp.ID, &mp.FkGameID, &mp.URL, &mp.CreatedAt); err != nil {
		return nil, err
	}
	return &mp, nil
}