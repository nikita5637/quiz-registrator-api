package model

import "errors"

// Games facade errors
var (
	ErrGameNoFreeSlots     = errors.New("game no free slots")
	ErrGameNotFound        = errors.New("game not found")
	ErrInvalidDate         = errors.New("invalid date")
	ErrInvalidGameID       = errors.New("invalid game ID")
	ErrInvalidPlayerDegree = errors.New("invalid player degree")
	ErrInvalidGameNumber   = errors.New("invalid game number")
	ErrInvalidGameType     = errors.New("invalid game type")
	ErrInvalidLeagueID     = errors.New("invalid league ID")
	ErrInvalidMaxPlayers   = errors.New("invalid max players")
	ErrInvalidPlaceID      = errors.New("invalid place id")
	ErrInvalidPrice        = errors.New("invalid price")
)

// Leagues facade errors
var (
	ErrLeagueNotFound = errors.New("league not found")
)

// Lottery errors
var (
	ErrLotteryNotAvailable     = errors.New("lottery not available")
	ErrLotteryNotImplemented   = errors.New("lottery not implemented")
	ErrLotteryPermissionDenied = errors.New("permission denied for lottery registration")
)

// Places facade errors
var (
	ErrPlaceNotFound = errors.New("place not found")
)

// User facade errors
var (
	// ErrUserNotFound ...
	ErrUserNotFound = errors.New("user not found")
	// ErrUserAlreadyExists ...
	ErrUserAlreadyExists = errors.New("user already exists")
	// ErrUserEmailValidate ...
	ErrUserEmailValidate = errors.New("invalid email format")
	// ErrUserNameValidateRequired ...
	ErrUserNameValidateRequired = errors.New("user name is required")
	// ErrUserNameValidateLength ...
	ErrUserNameValidateLength = errors.New("name length must be between 1 and 100 characters")
	// ErrUserNameValidateAlphabet ...
	ErrUserNameValidateAlphabet = errors.New("only Russian character set are allowed")
	// ErrUserPhoneValidate ...
	ErrUserPhoneValidate = errors.New("invalid phone format")
	// ErrUserStateValidate ...
	ErrUserStateValidate = errors.New("invalid user state")
)
