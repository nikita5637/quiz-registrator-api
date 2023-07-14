package users

import "errors"

var (
	// ErrUserNotFound ...
	ErrUserNotFound = errors.New("user not found")
	// ErrUserTelegramIDAlreadyExists ...
	ErrUserTelegramIDAlreadyExists = errors.New("user Telegram ID already exists")
	// ErrUserEmailAlreadyExists ...
	ErrUserEmailAlreadyExists = errors.New("user email already exists")
)
