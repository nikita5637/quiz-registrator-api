package mysql

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"

	"github.com/go-xorm/builder"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/logger"
)

// User represents a row from 'user'.
type User struct {
	ID         int            `json:"id"`          // id
	Name       string         `json:"name"`        // name
	TelegramID int64          `json:"telegram_id"` // telegram_id
	Email      sql.NullString `json:"email"`       // email
	Phone      sql.NullString `json:"phone"`       // phone
	State      int            `json:"state"`       // state
	CreatedAt  sql.NullTime   `json:"created_at"`  // created_at
	UpdatedAt  sql.NullTime   `json:"updated_at"`  // updated_at
}

// UserStorage is User service implementation
type UserStorage struct {
	db *sql.DB
}

// NewUserStorage creates new instance of UserStorage
func NewUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{
		db: db,
	}
}

// GetAll returns all records
func (s *UserStorage) GetAll(ctx context.Context) ([]User, error) {
	return s.Find(ctx, nil, "")
}

// Find perform find request by params
func (s *UserStorage) Find(ctx context.Context, q builder.Cond, sort string) ([]User, error) {
	query := `SELECT id, name, telegram_id, email, phone, state, created_at, updated_at FROM user`

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

	var items []User

	for rows.Next() {
		var item User
		if err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.TelegramID,
			&item.Email,
			&item.Phone,
			&item.State,
			&item.CreatedAt,
			&item.UpdatedAt,
		); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

// FindWithLimit perform find request by params, offset and limit
func (s *UserStorage) FindWithLimit(ctx context.Context, q builder.Cond, sort string, offset, limit uint64) ([]User, error) {
	query := `SELECT id, name, telegram_id, email, phone, state, created_at, updated_at FROM user`

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

	var items []User

	for rows.Next() {
		var item User
		if err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.TelegramID,
			&item.Email,
			&item.Phone,
			&item.State,
			&item.CreatedAt,
			&item.UpdatedAt,
		); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

// Total return count(*) by params
func (s *UserStorage) Total(ctx context.Context, q builder.Cond) (uint64, error) {
	query := `SELECT count(*) FROM user`

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

// Insert inserts the User to the database.
func (s *UserStorage) Insert(ctx context.Context, item User) (int, error) {
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO user (` +
		`name, telegram_id, email, phone, state, created_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, UTC_TIMESTAMP()` +
		`)`
	// run
	logger.Debugf(ctx, sqlstr, item.Name, item.TelegramID, item.Email, item.Phone, item.State)

	res, err := s.db.ExecContext(ctx, sqlstr, item.Name, item.TelegramID, item.Email, item.Phone, item.State)
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

// Update updates a User in the database.
func (s *UserStorage) Update(ctx context.Context, item User) error {
	// update with primary key
	const sqlstr = `UPDATE user SET ` +
		`name = ?, telegram_id = ?, email = ?, phone = ?, state = ?, updated_at = UTC_TIMESTAMP() ` +
		`WHERE id = ?`
	// run
	logger.Debugf(ctx, sqlstr, item.Name, item.TelegramID, item.Email, item.Phone, item.State, item.ID)
	if _, err := s.db.ExecContext(ctx, sqlstr, item.Name, item.TelegramID, item.Email, item.Phone, item.State, item.ID); err != nil {
		return err
	}

	return nil
}

// Upsert performs an upsert for User.
func (s *UserStorage) Upsert(ctx context.Context, item User) error {
	// upsert
	const sqlstr = `INSERT INTO user (` +
		`id, name, telegram_id, email, phone, state, created_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, UTC_TIMESTAMP()` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`name = VALUES(name), telegram_id = VALUES(telegram_id), email = VALUES(email), phone = VALUES(phone), state = VALUES(state), created_at = VALUES(created_at), updated_at = VALUES(updated_at)`
	// run
	logger.Debugf(ctx, sqlstr, item.ID, item.Name, item.TelegramID, item.Email, item.Phone, item.State, item.CreatedAt, item.UpdatedAt)
	if _, err := s.db.ExecContext(ctx, sqlstr, item.ID, item.Name, item.TelegramID, item.Email, item.Phone, item.State); err != nil {
		return err
	}

	return nil
}

// Delete deletes the User from the database.
func (s *UserStorage) Delete(ctx context.Context, id int) error {
	// delete with single primary key
	const sqlstr = `DELETE FROM user ` +
		`WHERE id = ?`
	// run
	logger.Debugf(ctx, sqlstr, id)

	if _, err := s.db.ExecContext(ctx, sqlstr, id); err != nil {
		return err
	}

	return nil
}

// GetUserByEmail retrieves a row from 'user' as a User.
//
// Generated from index 'email'.
func (s *UserStorage) GetUserByEmail(ctx context.Context, email sql.NullString) (*User, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, name, telegram_id, email, phone, state, created_at, updated_at ` +
		`FROM user ` +
		`WHERE email = ?`
	// run
	logger.Debugf(ctx, sqlstr, email)
	u := User{}
	if err := s.db.QueryRowContext(ctx, sqlstr, email).Scan(&u.ID, &u.Name, &u.TelegramID, &u.Email, &u.Phone, &u.State, &u.CreatedAt, &u.UpdatedAt); err != nil {
		return nil, err
	}
	return &u, nil
}

// GetUserByPhone retrieves a row from 'user' as a User.
//
// Generated from index 'phone'.
func (s *UserStorage) GetUserByPhone(ctx context.Context, phone sql.NullString) (*User, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, name, telegram_id, email, phone, state, created_at, updated_at ` +
		`FROM user ` +
		`WHERE phone = ?`
	// run
	logger.Debugf(ctx, sqlstr, phone)
	u := User{}
	if err := s.db.QueryRowContext(ctx, sqlstr, phone).Scan(&u.ID, &u.Name, &u.TelegramID, &u.Email, &u.Phone, &u.State, &u.CreatedAt, &u.UpdatedAt); err != nil {
		return nil, err
	}
	return &u, nil
}

// GetUserByTelegramID retrieves a row from 'user' as a User.
//
// Generated from index 'telegram_id'.
func (s *UserStorage) GetUserByTelegramID(ctx context.Context, telegramID int64) (*User, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, name, telegram_id, email, phone, state, created_at, updated_at ` +
		`FROM user ` +
		`WHERE telegram_id = ?`
	// run
	logger.Debugf(ctx, sqlstr, telegramID)
	u := User{}
	if err := s.db.QueryRowContext(ctx, sqlstr, telegramID).Scan(&u.ID, &u.Name, &u.TelegramID, &u.Email, &u.Phone, &u.State, &u.CreatedAt, &u.UpdatedAt); err != nil {
		return nil, err
	}
	return &u, nil
}

// GetUserByID retrieves a row from 'user' as a User.
//
// Generated from index 'user_id_pkey'.
func (s *UserStorage) GetUserByID(ctx context.Context, id int) (*User, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, name, telegram_id, email, phone, state, created_at, updated_at ` +
		`FROM user ` +
		`WHERE id = ?`
	// run
	logger.Debugf(ctx, sqlstr, id)
	u := User{}
	if err := s.db.QueryRowContext(ctx, sqlstr, id).Scan(&u.ID, &u.Name, &u.TelegramID, &u.Email, &u.Phone, &u.State, &u.CreatedAt, &u.UpdatedAt); err != nil {
		return nil, err
	}
	return &u, nil
}
