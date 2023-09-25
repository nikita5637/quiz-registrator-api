package croupier

import "errors"

var (
	// ErrLotteryNotAvailable ...
	ErrLotteryNotAvailable = errors.New("lottery not available")
	// ErrLotteryNotImplemented ...
	ErrLotteryNotImplemented = errors.New("lottery not implemented")
)
