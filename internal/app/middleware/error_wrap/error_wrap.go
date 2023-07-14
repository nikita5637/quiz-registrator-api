package errorwrap

import (
	"context"
	"errors"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	errInternalError       = errors.New("internal error")
	errInternalErrorLexeme = i18n.Lexeme{
		Key:      "err_internal_error",
		FallBack: "Internal error",
	}
	reasonInternalError = "INTERNAL_ERROR"
)

// ErrorWrap ...
func (m *Middleware) ErrorWrap() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, req)
		if err != nil {
			if st, ok := status.FromError(err); ok {
				switch st.Code() {
				case codes.Internal:
					message := st.Message()

					st = status.New(st.Code(), errInternalError.Error())
					errorInfo := &errdetails.ErrorInfo{
						Reason: reasonInternalError,
						Metadata: map[string]string{
							"error": message,
						},
					}
					localizedMessage := &errdetails.LocalizedMessage{
						Locale:  i18n.GetLangFromContext(ctx),
						Message: i18n.GetTranslator(errInternalErrorLexeme)(ctx),
					}
					st, _ = st.WithDetails(errorInfo, localizedMessage)
				}

				return nil, st.Err()
			}

			return nil, err
		}

		return resp, nil
	}
}
