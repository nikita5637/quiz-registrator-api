package errorwrap

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
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

					st = model.GetStatus(ctx,
						st.Code(),
						"internal error",
						reasonInternalError,
						map[string]string{
							"error": message,
						},
						errInternalErrorLexeme,
					)
				}

				return nil, st.Err()
			}

			return nil, err
		}

		return resp, nil
	}
}
