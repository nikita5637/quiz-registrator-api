package model

// RegisterPlayerStatus ...
type RegisterPlayerStatus int32

const (
	// RegisterPlayerStatusInvalid ...
	RegisterPlayerStatusInvalid RegisterPlayerStatus = iota
	// RegisterPlayerStatusOK ...
	RegisterPlayerStatusOK
	// RegisterPlayerStatusAlreadyRegistered ...
	RegisterPlayerStatusAlreadyRegistered
)

// UnregisterPlayerStatus ...
type UnregisterPlayerStatus int32

const (
	// UnregisterPlayerStatusInvalid ...
	UnregisterPlayerStatusInvalid UnregisterPlayerStatus = iota
	// UnregisterPlayerStatusOK ...
	UnregisterPlayerStatusOK
	// UnregisterPlayerStatusNotRegistered ...
	UnregisterPlayerStatusNotRegistered
)

// GamePlayer ...
type GamePlayer struct {
	ID           int32
	FkGameID     int32
	FkUserID     int32
	RegisteredBy int32
	Degree       int32
	CreatedAt    DateTime
	DeletedAt    DateTime
}
