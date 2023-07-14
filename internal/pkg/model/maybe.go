package model

import (
	"database/sql"
)

// MaybeInt32 ...
type MaybeInt32 struct {
	Valid bool
	Value int32
}

// NewMaybeInt32 ...
func NewMaybeInt32(value int32) MaybeInt32 {
	return MaybeInt32{
		Valid: true,
		Value: value,
	}
}

// ToSQL ...
func (v MaybeInt32) ToSQL() sql.NullInt32 {
	return sql.NullInt32{
		Valid: v.Valid,
		Int32: v.Value,
	}
}

// ToSQLNullInt64 ...
func (v MaybeInt32) ToSQLNullInt64() sql.NullInt64 {
	return sql.NullInt64{
		Valid: v.Valid,
		Int64: int64(v.Value),
	}
}

// MaybeString ...
type MaybeString struct {
	Valid bool
	Value string
}

// NewMaybeString ...
func NewMaybeString(value string) MaybeString {
	return MaybeString{
		Valid: true,
		Value: value,
	}
}

// ToSQL ...
func (v MaybeString) ToSQL() sql.NullString {
	return sql.NullString{
		Valid:  v.Valid,
		String: v.Value,
	}
}
