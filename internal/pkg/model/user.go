package model

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// UserState ...
type UserState int

const (
	// UserStateInvalid ...
	UserStateInvalid UserState = iota
	// UserStateWelcome ...
	UserStateWelcome
	// UserStateRegistered ...
	UserStateRegistered
	// UserStateChangingEmail ...
	UserStateChangingEmail
	// UserStateChangingName ...
	UserStateChangingName
	// UserStateChangingPhone ...
	UserStateChangingPhone
	// UserStateBanned ...
	UserStateBanned

	numberOfUserStates
)

// ToSQL ...
func (s UserState) ToSQL() int {
	return int(s)
}

// ValidateUserState ...
func ValidateUserState(value interface{}) error {
	v, ok := value.(UserState)
	if !ok {
		return errors.New("must be UserState")
	}

	return validation.Validate(v, validation.Max(numberOfUserStates-1))
}

// User ...
type User struct {
	ID         int32
	Name       string
	TelegramID int64
	Email      MaybeString
	Phone      MaybeString
	State      UserState
}
