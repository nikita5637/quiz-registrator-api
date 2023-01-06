package registrator

import (
	"context"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	unauthenticatedRequestReason = "unauthenticated request"
)

var (
	gameNotFoundLexeme = i18n.Lexeme{
		Key:      "game_not_found",
		FallBack: "Game not found",
	}
)

func getGameNotFoundStatus(ctx context.Context, err error, gameID int32) *status.Status {
	reason := fmt.Sprintf("game with id %d not found", gameID)
	return getStatus(ctx, codes.NotFound, err, reason, gameNotFoundLexeme)
}

func getStatus(ctx context.Context, code codes.Code, err error, reason string, lexeme i18n.Lexeme) *status.Status {
	st := status.New(code, err.Error())
	ei := &errdetails.ErrorInfo{
		Reason: reason,
	}
	lm := &errdetails.LocalizedMessage{
		Locale:  i18n.GetLangFromContext(ctx),
		Message: getTranslator(lexeme)(ctx),
	}
	st, err = st.WithDetails(ei, lm)
	if err != nil {
		panic(err)
	}

	return st
}

func getTranslator(lexeme i18n.Lexeme) func(ctx context.Context) string {
	return func(ctx context.Context) string {
		return i18n.Translate(ctx, lexeme.Key, lexeme.FallBack)
	}
}
