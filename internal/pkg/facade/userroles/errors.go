package userroles

import "errors"

var (
	// ErrUserNotFound ...
	ErrUserNotFound = errors.New("user not found")
	// ErrUserRoleAlreadyExists ...
	ErrUserRoleAlreadyExists = errors.New("user role already exists")
	// ErrUserRoleNotFound ...
	ErrUserRoleNotFound = errors.New("user role not found")
)
