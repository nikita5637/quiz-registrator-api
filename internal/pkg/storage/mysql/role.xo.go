package mysql

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql/driver"
	"fmt"
)

// Role is the 'role' enum type from schema 'registrator_api'.
type Role uint16

// Role values.
const (
	// RoleAdmin is the 'admin' role.
	RoleAdmin Role = 1
	// RoleManagement is the 'management' role.
	RoleManagement Role = 2
	// RoleUser is the 'user' role.
	RoleUser Role = 3
)

// String satisfies the fmt.Stringer interface.
func (r Role) String() string {
	switch r {
	case RoleAdmin:
		return "admin"
	case RoleManagement:
		return "management"
	case RoleUser:
		return "user"
	}
	return fmt.Sprintf("Role(%d)", r)
}

// MarshalText marshals Role into text.
func (r Role) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

// UnmarshalText unmarshals Role from text.
func (r *Role) UnmarshalText(buf []byte) error {
	switch str := string(buf); str {
	case "admin":
		*r = RoleAdmin
	case "management":
		*r = RoleManagement
	case "user":
		*r = RoleUser
	default:
		return ErrInvalidRole(str)
	}
	return nil
}

// Value satisfies the driver.Valuer interface.
func (r Role) Value() (driver.Value, error) {
	return r.String(), nil
}

// Scan satisfies the sql.Scanner interface.
func (r *Role) Scan(v interface{}) error {
	if buf, ok := v.([]byte); ok {
		return r.UnmarshalText(buf)
	}
	return ErrInvalidRole(fmt.Sprintf("%T", v))
}

// NullRole represents a null 'role' enum for schema 'registrator_api'.
type NullRole struct {
	Role Role
	// Valid is true if Role is not null.
	Valid bool
}

// Value satisfies the driver.Valuer interface.
func (nr NullRole) Value() (driver.Value, error) {
	if !nr.Valid {
		return nil, nil
	}
	return nr.Role.Value()
}

// Scan satisfies the sql.Scanner interface.
func (nr *NullRole) Scan(v interface{}) error {
	if v == nil {
		nr.Role, nr.Valid = 0, false
		return nil
	}
	err := nr.Role.Scan(v)
	nr.Valid = err == nil
	return err
}

// ErrInvalidRole is the invalid Role error.
type ErrInvalidRole string

// Error satisfies the error interface.
func (err ErrInvalidRole) Error() string {
	return fmt.Sprintf("invalid Role(%s)", string(err))
}
