package mysql

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"

	"github.com/go-xorm/builder"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/logger"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// Certificate represents a row from 'certificate'.
type Certificate struct {
	ID        int            `json:"id"`         // id
	Type      uint8          `json:"type"`       // type
	WonOn     int            `json:"won_on"`     // won_on
	SpentOn   sql.NullInt64  `json:"spent_on"`   // spent_on
	Info      sql.NullString `json:"info"`       // info
	CreatedAt sql.NullTime   `json:"created_at"` // created_at
	UpdatedAt sql.NullTime   `json:"updated_at"` // updated_at
	DeletedAt sql.NullTime   `json:"deleted_at"` // deleted_at
}

// CertificateStorage is Certificate service implementation
type CertificateStorage struct {
	db *tx.Manager
}

// NewCertificateStorage creates new instance of CertificateStorage
func NewCertificateStorage(txManager *tx.Manager) *CertificateStorage {
	return &CertificateStorage{
		db: txManager,
	}
}

// GetAll returns all records
func (s *CertificateStorage) GetAll(ctx context.Context) ([]Certificate, error) {
	return s.Find(ctx, nil, "")
}

// Find perform find request by params
func (s *CertificateStorage) Find(ctx context.Context, q builder.Cond, sort string) ([]Certificate, error) {
	query := `SELECT id, type, won_on, spent_on, info, created_at, updated_at, deleted_at FROM certificate`

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

	var items []Certificate

	for rows.Next() {
		var item Certificate
		if err := rows.Scan(
			&item.ID,
			&item.Type,
			&item.WonOn,
			&item.SpentOn,
			&item.Info,
			&item.CreatedAt,
			&item.UpdatedAt,
			&item.DeletedAt,
		); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

// FindWithLimit perform find request by params, offset and limit
func (s *CertificateStorage) FindWithLimit(ctx context.Context, q builder.Cond, sort string, offset, limit uint64) ([]Certificate, error) {
	query := `SELECT id, type, won_on, spent_on, info, created_at, updated_at, deleted_at FROM certificate`

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
		query += ` OFFSET ? LIMIT ?`
		args = append(args, offset)
		args = append(args, limit)
	}

	rows, err := s.db.Sync(ctx).QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Certificate

	for rows.Next() {
		var item Certificate
		if err := rows.Scan(
			&item.ID,
			&item.Type,
			&item.WonOn,
			&item.SpentOn,
			&item.Info,
			&item.CreatedAt,
			&item.UpdatedAt,
			&item.DeletedAt,
		); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

// Total return count(*) by params
func (s *CertificateStorage) Total(ctx context.Context, q builder.Cond) (uint64, error) {
	query := `SELECT count(*) FROM certificate`

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

// Insert inserts the Certificate to the database.
func (s *CertificateStorage) Insert(ctx context.Context, item Certificate) (int, error) {
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO certificate (` +
		`type, won_on, spent_on, info, created_at` +
		`) VALUES (` +
		`?, ?, ?, ?, UTC_TIMESTAMP()` +
		`)`
	// run
	logger.Debugf(ctx, sqlstr, item.Type, item.WonOn, item.SpentOn, item.Info)

	res, err := s.db.Master(ctx).ExecContext(ctx, sqlstr, item.Type, item.WonOn, item.SpentOn, item.Info)
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

// Update updates a Certificate in the database.
func (s *CertificateStorage) Update(ctx context.Context, item Certificate) error {
	// update with primary key
	const sqlstr = `UPDATE certificate SET ` +
		`type = ?, won_on = ?, spent_on = ?, info = ?, updated_at = UTC_TIMESTAMP() ` +
		`WHERE id = ?`
	// run
	logger.Debugf(ctx, sqlstr, item.Type, item.WonOn, item.SpentOn, item.Info, item.ID)
	if _, err := s.db.Master(ctx).ExecContext(ctx, sqlstr, item.Type, item.WonOn, item.SpentOn, item.Info, item.ID); err != nil {
		return err
	}

	return nil
}

// Upsert performs an upsert for Certificate.
func (s *CertificateStorage) Upsert(ctx context.Context, item Certificate) error {
	// upsert
	const sqlstr = `INSERT INTO certificate (` +
		`id, type, won_on, spent_on, info, created_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, UTC_TIMESTAMP()` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`type = VALUES(type), won_on = VALUES(won_on), spent_on = VALUES(spent_on), info = VALUES(info), created_at = VALUES(created_at), updated_at = VALUES(updated_at), deleted_at = VALUES(deleted_at)`
	// run
	logger.Debugf(ctx, sqlstr, item.ID, item.Type, item.WonOn, item.SpentOn, item.Info, item.CreatedAt, item.UpdatedAt, item.DeletedAt)
	if _, err := s.db.Master(ctx).ExecContext(ctx, sqlstr, item.ID, item.Type, item.WonOn, item.SpentOn, item.Info); err != nil {
		return err
	}

	return nil
}

// Delete deletes the Certificate from the database.
func (s *CertificateStorage) Delete(ctx context.Context, id int) error {
	// update with primary key
	const sqlstr = `UPDATE certificate SET updated_at = UTC_TIMESTAMP(), deleted_at = UTC_TIMESTAMP() WHERE id = ?`
	// run
	logger.Debugf(ctx, sqlstr, id)

	if _, err := s.db.Master(ctx).ExecContext(ctx, sqlstr, id); err != nil {
		return err
	}

	return nil
}

// GetCertificateByID retrieves a row from 'certificate' as a Certificate.
//
// Generated from index 'certificate_id_pkey'.
func (s *CertificateStorage) GetCertificateByID(ctx context.Context, id int) (*Certificate, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, type, won_on, spent_on, info, created_at, updated_at, deleted_at ` +
		`FROM certificate ` +
		`WHERE id = ?`
	// run
	logger.Debugf(ctx, sqlstr, id)
	c := Certificate{}
	if err := s.db.Sync(ctx).QueryRowContext(ctx, sqlstr, id).Scan(&c.ID, &c.Type, &c.WonOn, &c.SpentOn, &c.Info, &c.CreatedAt, &c.UpdatedAt, &c.DeletedAt); err != nil {
		return nil, err
	}
	return &c, nil
}

// GetCertificateBySpentOn retrieves a row from 'certificate' as a Certificate.
//
// Generated from index 'spent_on'.
func (s *CertificateStorage) GetCertificateBySpentOn(ctx context.Context, spentOn sql.NullInt64) ([]*Certificate, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, type, won_on, spent_on, info, created_at, updated_at, deleted_at ` +
		`FROM certificate ` +
		`WHERE spent_on = ?`
	// run
	logger.Debugf(ctx, sqlstr, spentOn)
	rows, err := s.db.Sync(ctx).QueryContext(ctx, sqlstr, spentOn)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// process
	var res []*Certificate
	for rows.Next() {
		c := Certificate{}
		// scan
		if err := rows.Scan(&c.ID, &c.Type, &c.WonOn, &c.SpentOn, &c.Info, &c.CreatedAt, &c.UpdatedAt, &c.DeletedAt); err != nil {
			return nil, err
		}
		res = append(res, &c)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

// GetCertificateByWonOn retrieves a row from 'certificate' as a Certificate.
//
// Generated from index 'won_on'.
func (s *CertificateStorage) GetCertificateByWonOn(ctx context.Context, wonOn int) ([]*Certificate, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, type, won_on, spent_on, info, created_at, updated_at, deleted_at ` +
		`FROM certificate ` +
		`WHERE won_on = ?`
	// run
	logger.Debugf(ctx, sqlstr, wonOn)
	rows, err := s.db.Sync(ctx).QueryContext(ctx, sqlstr, wonOn)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// process
	var res []*Certificate
	for rows.Next() {
		c := Certificate{}
		// scan
		if err := rows.Scan(&c.ID, &c.Type, &c.WonOn, &c.SpentOn, &c.Info, &c.CreatedAt, &c.UpdatedAt, &c.DeletedAt); err != nil {
			return nil, err
		}
		res = append(res, &c)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return res, nil
}