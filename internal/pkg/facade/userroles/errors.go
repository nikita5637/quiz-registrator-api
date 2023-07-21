package userroles

import "errors"

var (
	// ErrUserNotFound ...
	ErrUserNotFound = errors.New("user not found")
	// ErrRoleIsAlreadyAssigned ...
	ErrRoleIsAlreadyAssigned = errors.New("role is already assigned to user")
	// ErrUserRoleNotFound ...
	ErrUserRoleNotFound = errors.New("user role not found")
)
