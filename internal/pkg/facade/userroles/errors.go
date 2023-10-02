package userroles

import "errors"

var (
	// ErrRoleIsAlreadyAssigned ...
	ErrRoleIsAlreadyAssigned = errors.New("role is already assigned to user")
	// ErrUserRoleNotFound ...
	ErrUserRoleNotFound = errors.New("user role not found")
)
