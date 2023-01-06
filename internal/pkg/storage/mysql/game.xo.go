package mysql

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-xorm/builder"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/logger"
)

// Game represents a row from 'game'.
type Game struct {
	ID          int            `json:"id"`           // id
	ExternalID  sql.NullInt64  `json:"external_id"`  // external_id
	LeagueID    int            `json:"league_id"`    // league_id
	Type        uint8          `json:"type"`         // type
	Number      string         `json:"number"`       // number
	Name        sql.NullString `json:"name"`         // name
	PlaceID     int            `json:"place_id"`     // place_id
	Date        time.Time      `json:"date"`         // date
	Price       uint           `json:"price"`        // price
	PaymentType []byte         `json:"payment_type"` // payment_type
	MaxPlayers  uint8          `json:"max_players"`  // max_players
	Payment     sql.NullInt64  `json:"payment"`      // payment
	Registered  bool           `json:"registered"`   // registered
	CreatedAt   sql.NullTime   `json:"created_at"`   // created_at
	UpdatedAt   sql.NullTime   `json:"updated_at"`   // updated_at
	DeletedAt   sql.NullTime   `json:"deleted_at"`   // deleted_at
}

// GameStorage is Game service implementation
type GameStorage struct {
	db *sql.DB
}

// NewGameStorage creates new instance of GameStorage
func NewGameStorage(db *sql.DB) *GameStorage {
	return &GameStorage{
		db: db,
	}
}

// GetAll returns all records
func (s *GameStorage) GetAll(ctx context.Context) ([]Game, error) {
	return s.Find(ctx, nil, "")
}

// Find perform find request by params
func (s *GameStorage) Find(ctx context.Context, q builder.Cond, sort string) ([]Game, error) {
	query := `SELECT id, external_id, league_id, type, number, name, place_id, date, price, payment_type, max_players, payment, registered, created_at, updated_at, deleted_at FROM game`

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

	var items []Game

	for rows.Next() {
		var item Game
		if err := rows.Scan(
			&item.ID,
			&item.ExternalID,
			&item.LeagueID,
			&item.Type,
			&item.Number,
			&item.Name,
			&item.PlaceID,
			&item.Date,
			&item.Price,
			&item.PaymentType,
			&item.MaxPlayers,
			&item.Payment,
			&item.Registered,
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
func (s *GameStorage) FindWithLimit(ctx context.Context, q builder.Cond, sort string, offset, limit uint64) ([]Game, error) {
	query := `SELECT id, external_id, league_id, type, number, name, place_id, date, price, payment_type, max_players, payment, registered, created_at, updated_at, deleted_at FROM game`

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

	var items []Game

	for rows.Next() {
		var item Game
		if err := rows.Scan(
			&item.ID,
			&item.ExternalID,
			&item.LeagueID,
			&item.Type,
			&item.Number,
			&item.Name,
			&item.PlaceID,
			&item.Date,
			&item.Price,
			&item.PaymentType,
			&item.MaxPlayers,
			&item.Payment,
			&item.Registered,
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
func (s *GameStorage) Total(ctx context.Context, q builder.Cond) (uint64, error) {
	query := `SELECT count(*) FROM game`

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

// Insert inserts the Game to the database.
func (s *GameStorage) Insert(ctx context.Context, item Game) (int, error) {
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO game (` +
		`external_id, league_id, type, number, name, place_id, date, price, payment_type, max_players, payment, registered, created_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, UTC_TIMESTAMP()` +
		`)`
	// run
	logger.Debugf(ctx, sqlstr, item.ExternalID, item.LeagueID, item.Type, item.Number, item.Name, item.PlaceID, item.Date, item.Price, item.PaymentType, item.MaxPlayers, item.Payment, item.Registered)

	res, err := s.db.ExecContext(ctx, sqlstr, item.ExternalID, item.LeagueID, item.Type, item.Number, item.Name, item.PlaceID, item.Date, item.Price, item.PaymentType, item.MaxPlayers, item.Payment, item.Registered)
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

// Update updates a Game in the database.
func (s *GameStorage) Update(ctx context.Context, item Game) error {
	// update with primary key
	const sqlstr = `UPDATE game SET ` +
		`external_id = ?, league_id = ?, type = ?, number = ?, name = ?, place_id = ?, date = ?, price = ?, payment_type = ?, max_players = ?, payment = ?, registered = ?, updated_at = UTC_TIMESTAMP() ` +
		`WHERE id = ?`
	// run
	logger.Debugf(ctx, sqlstr, item.ExternalID, item.LeagueID, item.Type, item.Number, item.Name, item.PlaceID, item.Date, item.Price, item.PaymentType, item.MaxPlayers, item.Payment, item.Registered, item.ID)
	if _, err := s.db.ExecContext(ctx, sqlstr, item.ExternalID, item.LeagueID, item.Type, item.Number, item.Name, item.PlaceID, item.Date, item.Price, item.PaymentType, item.MaxPlayers, item.Payment, item.Registered, item.ID); err != nil {
		return err
	}

	return nil
}

// Upsert performs an upsert for Game.
func (s *GameStorage) Upsert(ctx context.Context, item Game) error {
	// upsert
	const sqlstr = `INSERT INTO game (` +
		`id, external_id, league_id, type, number, name, place_id, date, price, payment_type, max_players, payment, registered, created_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, UTC_TIMESTAMP()` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`external_id = VALUES(external_id), league_id = VALUES(league_id), type = VALUES(type), number = VALUES(number), name = VALUES(name), place_id = VALUES(place_id), date = VALUES(date), price = VALUES(price), payment_type = VALUES(payment_type), max_players = VALUES(max_players), payment = VALUES(payment), registered = VALUES(registered), created_at = VALUES(created_at), updated_at = VALUES(updated_at), deleted_at = VALUES(deleted_at)`
	// run
	logger.Debugf(ctx, sqlstr, item.ID, item.ExternalID, item.LeagueID, item.Type, item.Number, item.Name, item.PlaceID, item.Date, item.Price, item.PaymentType, item.MaxPlayers, item.Payment, item.Registered, item.CreatedAt, item.UpdatedAt, item.DeletedAt)
	if _, err := s.db.ExecContext(ctx, sqlstr, item.ID, item.ExternalID, item.LeagueID, item.Type, item.Number, item.Name, item.PlaceID, item.Date, item.Price, item.PaymentType, item.MaxPlayers, item.Payment, item.Registered); err != nil {
		return err
	}

	return nil
}

// Delete deletes the Game from the database.
func (s *GameStorage) Delete(ctx context.Context, id int) error {
	// update with primary key
	const sqlstr = `UPDATE game SET updated_at = UTC_TIMESTAMP(), deleted_at = UTC_TIMESTAMP() WHERE id = ?`
	// run
	logger.Debugf(ctx, sqlstr, id)

	if _, err := s.db.ExecContext(ctx, sqlstr, id); err != nil {
		return err
	}

	return nil
}

// GetGameByID retrieves a row from 'game' as a Game.
//
// Generated from index 'game_id_pkey'.
func (s *GameStorage) GetGameByID(ctx context.Context, id int) (*Game, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, external_id, league_id, type, number, name, place_id, date, price, payment_type, max_players, payment, registered, created_at, updated_at, deleted_at ` +
		`FROM game ` +
		`WHERE id = ?`
	// run
	logger.Debugf(ctx, sqlstr, id)
	g := Game{}
	if err := s.db.QueryRowContext(ctx, sqlstr, id).Scan(&g.ID, &g.ExternalID, &g.LeagueID, &g.Type, &g.Number, &g.Name, &g.PlaceID, &g.Date, &g.Price, &g.PaymentType, &g.MaxPlayers, &g.Payment, &g.Registered, &g.CreatedAt, &g.UpdatedAt, &g.DeletedAt); err != nil {
		return nil, err
	}
	return &g, nil
}

// GetGameByLeagueID retrieves a row from 'game' as a Game.
//
// Generated from index 'league'.
func (s *GameStorage) GetGameByLeagueID(ctx context.Context, leagueID int) ([]*Game, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, external_id, league_id, type, number, name, place_id, date, price, payment_type, max_players, payment, registered, created_at, updated_at, deleted_at ` +
		`FROM game ` +
		`WHERE league_id = ?`
	// run
	logger.Debugf(ctx, sqlstr, leagueID)
	rows, err := s.db.QueryContext(ctx, sqlstr, leagueID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// process
	var res []*Game
	for rows.Next() {
		g := Game{}
		// scan
		if err := rows.Scan(&g.ID, &g.ExternalID, &g.LeagueID, &g.Type, &g.Number, &g.Name, &g.PlaceID, &g.Date, &g.Price, &g.PaymentType, &g.MaxPlayers, &g.Payment, &g.Registered, &g.CreatedAt, &g.UpdatedAt, &g.DeletedAt); err != nil {
			return nil, err
		}
		res = append(res, &g)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

// GetGameByPlaceID retrieves a row from 'game' as a Game.
//
// Generated from index 'place_id'.
func (s *GameStorage) GetGameByPlaceID(ctx context.Context, placeID int) ([]*Game, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, external_id, league_id, type, number, name, place_id, date, price, payment_type, max_players, payment, registered, created_at, updated_at, deleted_at ` +
		`FROM game ` +
		`WHERE place_id = ?`
	// run
	logger.Debugf(ctx, sqlstr, placeID)
	rows, err := s.db.QueryContext(ctx, sqlstr, placeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// process
	var res []*Game
	for rows.Next() {
		g := Game{}
		// scan
		if err := rows.Scan(&g.ID, &g.ExternalID, &g.LeagueID, &g.Type, &g.Number, &g.Name, &g.PlaceID, &g.Date, &g.Price, &g.PaymentType, &g.MaxPlayers, &g.Payment, &g.Registered, &g.CreatedAt, &g.UpdatedAt, &g.DeletedAt); err != nil {
			return nil, err
		}
		res = append(res, &g)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return res, nil
}
