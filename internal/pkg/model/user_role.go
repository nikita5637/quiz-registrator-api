package model

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Role ...
type Role uint8

const (
	// RoleAdmin ...
	RoleAdmin Role = iota + 1
	// RoleManagement ...
	RoleManagement
	// RoleUser ...
	RoleUser
	// ...
	numberOfRoles

	admin      = "admin"
	management = "management"
	user       = "user"
)

// String ...
func (r Role) String() string {
	switch r {
	case RoleAdmin:
		return admin
	case RoleManagement:
		return management
	case RoleUser:
		return user
	}

	return "invalid"
}

// ValidateRole ...
func ValidateRole(value interface{}) error {
	v, ok := value.(Role)
	if !ok {
		return errors.New("must be Role")
	}

	return validation.Validate(v, validation.Required, validation.Min(RoleAdmin), validation.Max(numberOfRoles-1))
}

// UserRole ...
type UserRole struct {
	ID     int32
	UserID int32
	Role   Role
}
