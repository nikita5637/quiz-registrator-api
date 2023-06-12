package log

import (
	"context"
	"time"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/logger"
	timeutils "github.com/nikita5637/quiz-registrator-api/utils/time"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// Log ...
func (m *Middleware) Log() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := timeutils.TimeNow()

		h, err := handler(ctx, req)

		if err != nil {
			st := status.Convert(err)
			logger.Errorf(ctx, "Request - Method:%s Duration:%s Error:%v Details: %v",
				info.FullMethod,
				time.Since(start),
				err,
				st.Details(),
			)
		} else {
			logger.Debugf(ctx, "Request - Method:%s Duration:%s",
				info.FullMethod,
				time.Since(start),
			)
		}

		return h, err
	}
}
