package mysql

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"

	"github.com/go-xorm/builder"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/logger"
)

// Place represents a row from 'place'.
type Place struct {
	ID        int             `json:"id"`         // id
	Name      string          `json:"name"`       // name
	Address   string          `json:"address"`    // address
	ShortName sql.NullString  `json:"short_name"` // short_name
	Latitude  sql.NullFloat64 `json:"latitude"`   // latitude
	Longitude sql.NullFloat64 `json:"longitude"`  // longitude
	MenuLink  sql.NullString  `json:"menu_link"`  // menu_link
}

// PlaceStorage is Place service implementation
type PlaceStorage struct {
	db *sql.DB
}

// NewPlaceStorage creates new instance of PlaceStorage
func NewPlaceStorage(db *sql.DB) *PlaceStorage {
	return &PlaceStorage{
		db: db,
	}
}

// GetAll returns all records
func (s *PlaceStorage) GetAll(ctx context.Context) ([]Place, error) {
	return s.Find(ctx, nil, "")
}

// Find perform find request by params
func (s *PlaceStorage) Find(ctx context.Context, q builder.Cond, sort string) ([]Place, error) {
	query := `SELECT id, name, address, short_name, latitude, longitude, menu_link FROM place`

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

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Place

	for rows.Next() {
		var item Place
		if err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Address,
			&item.ShortName,
			&item.Latitude,
			&item.Longitude,
			&item.MenuLink,
		); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

// FindWithLimit perform find request by params, offset and limit
func (s *PlaceStorage) FindWithLimit(ctx context.Context, q builder.Cond, sort string, offset, limit uint64) ([]Place, error) {
	query := `SELECT id, name, address, short_name, latitude, longitude, menu_link FROM place`

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

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Place

	for rows.Next() {
		var item Place
		if err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Address,
			&item.ShortName,
			&item.Latitude,
			&item.Longitude,
			&item.MenuLink,
		); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

// Total return count(*) by params
func (s *PlaceStorage) Total(ctx context.Context, q builder.Cond) (uint64, error) {
	query := `SELECT count(*) FROM place`

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

	rows, err := s.db.QueryContext(ctx, query, args...)
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

// Insert inserts the Place to the database.
func (s *PlaceStorage) Insert(ctx context.Context, item Place) (int, error) {
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO place (` +
		`name, address, short_name, latitude, longitude, menu_link` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?` +
		`)`
	// run
	logger.Debugf(ctx, sqlstr, item.Name, item.Address, item.ShortName, item.Latitude, item.Longitude, item.MenuLink)

	res, err := s.db.ExecContext(ctx, sqlstr, item.Name, item.Address, item.ShortName, item.Latitude, item.Longitude, item.MenuLink)
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

// Update updates a Place in the database.
func (s *PlaceStorage) Update(ctx context.Context, item Place) error {
	// update with primary key
	const sqlstr = `UPDATE place SET ` +
		`name = ?, address = ?, short_name = ?, latitude = ?, longitude = ?, menu_link = ? ` +
		`WHERE id = ?`
	// run
	logger.Debugf(ctx, sqlstr, item.Name, item.Address, item.ShortName, item.Latitude, item.Longitude, item.MenuLink, item.ID)
	if _, err := s.db.ExecContext(ctx, sqlstr, item.Name, item.Address, item.ShortName, item.Latitude, item.Longitude, item.MenuLink, item.ID); err != nil {
		return err
	}

	return nil
}

// Upsert performs an upsert for Place.
func (s *PlaceStorage) Upsert(ctx context.Context, item Place) error {
	// upsert
	const sqlstr = `INSERT INTO place (` +
		`id, name, address, short_name, latitude, longitude, menu_link` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`name = VALUES(name), address = VALUES(address), short_name = VALUES(short_name), latitude = VALUES(latitude), longitude = VALUES(longitude), menu_link = VALUES(menu_link)`
	// run
	logger.Debugf(ctx, sqlstr, item.ID, item.Name, item.Address, item.ShortName, item.Latitude, item.Longitude, item.MenuLink)
	if _, err := s.db.ExecContext(ctx, sqlstr, item.ID, item.Name, item.Address, item.ShortName, item.Latitude, item.Longitude, item.MenuLink); err != nil {
		return err
	}

	return nil
}

// Delete deletes the Place from the database.
func (s *PlaceStorage) Delete(ctx context.Context, id int) error {
	// delete with single primary key
	const sqlstr = `DELETE FROM place ` +
		`WHERE id = ?`
	// run
	logger.Debugf(ctx, sqlstr, id)

	if _, err := s.db.ExecContext(ctx, sqlstr, id); err != nil {
		return err
	}

	return nil
}

// GetPlaceByID retrieves a row from 'place' as a Place.
//
// Generated from index 'place_id_pkey'.
func (s *PlaceStorage) GetPlaceByID(ctx context.Context, id int) (*Place, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, name, address, short_name, latitude, longitude, menu_link ` +
		`FROM place ` +
		`WHERE id = ?`
	// run
	logger.Debugf(ctx, sqlstr, id)
	p := Place{}
	if err := s.db.QueryRowContext(ctx, sqlstr, id).Scan(&p.ID, &p.Name, &p.Address, &p.ShortName, &p.Latitude, &p.Longitude, &p.MenuLink); err != nil {
		return nil, err
	}
	return &p, nil
}