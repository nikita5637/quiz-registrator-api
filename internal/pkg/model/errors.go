package model

import (
	"context"
	"errors"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Games facade errors
var (
	ErrGameResultAlreadyExists = errors.New("game result already exists")
	ErrGameResultNotFound      = errors.New("game result not found")
	ErrInvalidDate             = errors.New("invalid date")
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

// GetStatus ...
func GetStatus(ctx context.Context, code codes.Code, message, reason string, metadata map[string]string, lexeme i18n.Lexeme) *status.Status {
	st := status.New(code, message)
	ei := &errdetails.ErrorInfo{
		Reason:   reason,
		Metadata: metadata,
	}
	lm := &errdetails.LocalizedMessage{
		Locale:  i18n.GetLangFromContext(ctx),
		Message: i18n.GetTranslator(lexeme)(ctx),
	}
	st, err := st.WithDetails(ei, lm)
	if err != nil {
		panic(err)
	}

	return st
}
