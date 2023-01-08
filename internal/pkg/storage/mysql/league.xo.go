package mysql

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"

	"github.com/go-xorm/builder"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/logger"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// League represents a row from 'league'.
type League struct {
	ID        int            `json:"id"`         // id
	Name      string         `json:"name"`       // name
	ShortName sql.NullString `json:"short_name"` // short_name
	LogoLink  sql.NullString `json:"logo_link"`  // logo_link
	WebSite   sql.NullString `json:"web_site"`   // web_site
}

// LeagueStorage is League service implementation
type LeagueStorage struct {
	db *tx.Manager
}

// NewLeagueStorage creates new instance of LeagueStorage
func NewLeagueStorage(db *sql.DB) *LeagueStorage {
	return &LeagueStorage{
		db: tx.NewManager(db),
	}
}

// GetAll returns all records
func (s *LeagueStorage) GetAll(ctx context.Context) ([]League, error) {
	return s.Find(ctx, nil, "")
}

// Find perform find request by params
func (s *LeagueStorage) Find(ctx context.Context, q builder.Cond, sort string) ([]League, error) {
	query := `SELECT id, name, short_name, logo_link, web_site FROM league`

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

	var items []League

	for rows.Next() {
		var item League
		if err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.ShortName,
			&item.LogoLink,
			&item.WebSite,
		); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

// FindWithLimit perform find request by params, offset and limit
func (s *LeagueStorage) FindWithLimit(ctx context.Context, q builder.Cond, sort string, offset, limit uint64) ([]League, error) {
	query := `SELECT id, name, short_name, logo_link, web_site FROM league`

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

	var items []League

	for rows.Next() {
		var item League
		if err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.ShortName,
			&item.LogoLink,
			&item.WebSite,
		); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

// Total return count(*) by params
func (s *LeagueStorage) Total(ctx context.Context, q builder.Cond) (uint64, error) {
	query := `SELECT count(*) FROM league`

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

// Insert inserts the League to the database.
func (s *LeagueStorage) Insert(ctx context.Context, item League) (int, error) {
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO league (` +
		`name, short_name, logo_link, web_site` +
		`) VALUES (` +
		`?, ?, ?, ?` +
		`)`
	// run
	logger.Debugf(ctx, sqlstr, item.Name, item.ShortName, item.LogoLink, item.WebSite)

	res, err := s.db.Master(ctx).ExecContext(ctx, sqlstr, item.Name, item.ShortName, item.LogoLink, item.WebSite)
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

// Update updates a League in the database.
func (s *LeagueStorage) Update(ctx context.Context, item League) error {
	// update with primary key
	const sqlstr = `UPDATE league SET ` +
		`name = ?, short_name = ?, logo_link = ?, web_site = ? ` +
		`WHERE id = ?`
	// run
	logger.Debugf(ctx, sqlstr, item.Name, item.ShortName, item.LogoLink, item.WebSite, item.ID)
	if _, err := s.db.Master(ctx).ExecContext(ctx, sqlstr, item.Name, item.ShortName, item.LogoLink, item.WebSite, item.ID); err != nil {
		return err
	}

	return nil
}

// Upsert performs an upsert for League.
func (s *LeagueStorage) Upsert(ctx context.Context, item League) error {
	// upsert
	const sqlstr = `INSERT INTO league (` +
		`id, name, short_name, logo_link, web_site` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`name = VALUES(name), short_name = VALUES(short_name), logo_link = VALUES(logo_link), web_site = VALUES(web_site)`
	// run
	logger.Debugf(ctx, sqlstr, item.ID, item.Name, item.ShortName, item.LogoLink, item.WebSite)
	if _, err := s.db.Master(ctx).ExecContext(ctx, sqlstr, item.ID, item.Name, item.ShortName, item.LogoLink, item.WebSite); err != nil {
		return err
	}

	return nil
}

// Delete deletes the League from the database.
func (s *LeagueStorage) Delete(ctx context.Context, id int) error {
	// delete with single primary key
	const sqlstr = `DELETE FROM league ` +
		`WHERE id = ?`
	// run
	logger.Debugf(ctx, sqlstr, id)

	if _, err := s.db.Master(ctx).ExecContext(ctx, sqlstr, id); err != nil {
		return err
	}

	return nil
}

// GetLeagueByID retrieves a row from 'league' as a League.
//
// Generated from index 'league_id_pkey'.
func (s *LeagueStorage) GetLeagueByID(ctx context.Context, id int) (*League, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, name, short_name, logo_link, web_site ` +
		`FROM league ` +
		`WHERE id = ?`
	// run
	logger.Debugf(ctx, sqlstr, id)
	l := League{}
	if err := s.db.Sync(ctx).QueryRowContext(ctx, sqlstr, id).Scan(&l.ID, &l.Name, &l.ShortName, &l.LogoLink, &l.WebSite); err != nil {
		return nil, err
	}
	return &l, nil
}

// GetLeagueByName retrieves a row from 'league' as a League.
//
// Generated from index 'name'.
func (s *LeagueStorage) GetLeagueByName(ctx context.Context, name string) (*League, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, name, short_name, logo_link, web_site ` +
		`FROM league ` +
		`WHERE name = ?`
	// run
	logger.Debugf(ctx, sqlstr, name)
	l := League{}
	if err := s.db.Sync(ctx).QueryRowContext(ctx, sqlstr, name).Scan(&l.ID, &l.Name, &l.ShortName, &l.LogoLink, &l.WebSite); err != nil {
		return nil, err
	}
	return &l, nil
}

// GetLeagueByShortName retrieves a row from 'league' as a League.
//
// Generated from index 'short_name'.
func (s *LeagueStorage) GetLeagueByShortName(ctx context.Context, shortName sql.NullString) (*League, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, name, short_name, logo_link, web_site ` +
		`FROM league ` +
		`WHERE short_name = ?`
	// run
	logger.Debugf(ctx, sqlstr, shortName)
	l := League{}
	if err := s.db.Sync(ctx).QueryRowContext(ctx, sqlstr, shortName).Scan(&l.ID, &l.Name, &l.ShortName, &l.LogoLink, &l.WebSite); err != nil {
		return nil, err
	}
	return &l, nil
}

// GetLeagueByWebSite retrieves a row from 'league' as a League.
//
// Generated from index 'web_site'.
func (s *LeagueStorage) GetLeagueByWebSite(ctx context.Context, webSite sql.NullString) (*League, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, name, short_name, logo_link, web_site ` +
		`FROM league ` +
		`WHERE web_site = ?`
	// run
	logger.Debugf(ctx, sqlstr, webSite)
	l := League{}
	if err := s.db.Sync(ctx).QueryRowContext(ctx, sqlstr, webSite).Scan(&l.ID, &l.Name, &l.ShortName, &l.LogoLink, &l.WebSite); err != nil {
		return nil, err
	}
	return &l, nil
}
