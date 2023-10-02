package model

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/mono83/maybe"
)

// UserState ...
type UserState int

const (
	// UserStateWelcome ...
	UserStateWelcome UserState = iota + 1
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
	// UserStateChangingBirthdate ...
	UserStateChangingBirthdate
	// UserStateChangingSex ...
	UserStateChangingSex

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

	return validation.Validate(v, validation.Required, validation.Min(UserStateWelcome), validation.Max(numberOfUserStates-1))
}

// Sex ...
type Sex int32

const (
	// SexMale ...
	SexMale Sex = iota + 1
	// SexFemale ...
	SexFemale

	numberOfSexes
)

// ValidateSex ...
func ValidateSex(value interface{}) error {
	v, ok := value.(Sex)
	if !ok {
		return errors.New("must be Sex")
	}

	return validation.Validate(v, validation.Required, validation.Min(SexMale), validation.Max(numberOfSexes-1))
}

// User ...
type User struct {
	ID         int32
	Name       string
	TelegramID int64
	Email      maybe.Maybe[string]
	Phone      maybe.Maybe[string]
	State      UserState
	Birthdate  maybe.Maybe[string]
	Sex        maybe.Maybe[Sex]
}
