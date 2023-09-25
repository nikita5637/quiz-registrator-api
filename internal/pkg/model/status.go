package model

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
