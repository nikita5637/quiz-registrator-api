package errorwrap

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func okHandler(ctx context.Context, req interface{}) (interface{}, error) {
	return nil, nil
}

func internalErrorHandler(ctx context.Context, req interface{}) (interface{}, error) {
	return nil, status.New(codes.Internal, "internal error").Err()
}

func unavailableErrorHandler(ctx context.Context, req interface{}) (interface{}, error) {
	return nil, status.New(codes.Unavailable, "unavailable").Err()
}

func otherErrorHandler(ctx context.Context, req interface{}) (interface{}, error) {
	return nil, errors.New("some error")
}

func TestMiddleware_ErrorWrap(t *testing.T) {
	t.Run("internal error", func(t *testing.T) {
		fx := tearUp(t)

		fn := fx.middleware.ErrorWrap()
		got, err := fn(fx.ctx, nil, nil, internalErrorHandler)
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 2)

		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, reasonInternalError, errorInfo.Reason)
		assert.Equal(t, map[string]string{
			"error": "internal error",
		}, errorInfo.Metadata)
	})

	t.Run("unavailable error", func(t *testing.T) {
		fx := tearUp(t)

		fn := fx.middleware.ErrorWrap()
		got, err := fn(fx.ctx, nil, nil, unavailableErrorHandler)
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Unavailable, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("other error", func(t *testing.T) {
		fx := tearUp(t)

		fn := fx.middleware.ErrorWrap()
		got, err := fn(fx.ctx, nil, nil, otherErrorHandler)
		assert.Nil(t, got)
		assert.Error(t, err)

		_, ok := status.FromError(err)
		assert.False(t, ok)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fn := fx.middleware.ErrorWrap()
		got, err := fn(fx.ctx, nil, nil, okHandler)
		assert.Nil(t, got)
		assert.NoError(t, err)
	})
}
