package model

import (
	"context"
	"errors"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Certificate facade errors
var (
	ErrCertificateNotFound = errors.New("certificate not found")
	ErrWonOnGameNotFound   = errors.New("won on game not found")
	ErrSpentOnGameNotFound = errors.New("spent on game not found")
)

// Games facade errors
var (
	ErrGameNoFreeSlots         = errors.New("game no free slots")
	ErrGameNotFound            = errors.New("game not found")
	ErrGameResultAlreadyExists = errors.New("game result already exists")
	ErrGameResultNotFound      = errors.New("game result not found")
	ErrInvalidDate             = errors.New("invalid date")
	ErrInvalidGameID           = errors.New("invalid game ID")
	ErrInvalidPlayerDegree     = errors.New("invalid player degree")
	ErrInvalidGameNumber       = errors.New("invalid game number")
	ErrInvalidGameType         = errors.New("invalid game type")
	ErrInvalidLeagueID         = errors.New("invalid league ID")
	ErrInvalidMaxPlayers       = errors.New("invalid max players")
	ErrInvalidPlaceID          = errors.New("invalid place id")
	ErrInvalidPrice            = errors.New("invalid price")
)

// Lottery errors
var (
	ErrLotteryNotAvailable   = errors.New("lottery not available")
	ErrLotteryNotImplemented = errors.New("lottery not implemented")
)

// Places facade errors
var (
	ErrPlaceNotFound = errors.New("place not found")
)

// User roles facade errors
var (
	// ErrUserRoleAlreadyExists ...
	ErrUserRoleAlreadyExists = errors.New("user role already exists")
	// ErrUserRoleNotFound ...
	ErrUserRoleNotFound = errors.New("user role not found")
)

// GetStatus ...
func GetStatus(ctx context.Context, code codes.Code, err error, reason string, lexeme i18n.Lexeme) *status.Status {
	st := status.New(code, err.Error())
	ei := &errdetails.ErrorInfo{
		Reason: reason,
	}
	lm := &errdetails.LocalizedMessage{
		Locale:  i18n.GetLangFromContext(ctx),
		Message: i18n.GetTranslator(lexeme)(ctx),
	}
	st, err = st.WithDetails(ei, lm)
	if err != nil {
		panic(err)
	}

	return st
}
