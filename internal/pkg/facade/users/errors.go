package users

import "errors"

const (
	// ReasonUserNotFound ...
	ReasonUserNotFound = "USER_NOT_FOUND"
)

var (
	// ErrUserNotFound ...
	ErrUserNotFound = errors.New("user not found")
	// ErrUserTelegramIDAlreadyExists ...
	ErrUserTelegramIDAlreadyExists = errors.New("user Telegram ID already exists")
	// ErrUserEmailAlreadyExists ...
	ErrUserEmailAlreadyExists = errors.New("user email already exists")
)
