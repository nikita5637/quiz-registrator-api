package log

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func errorHandler(_ context.Context, _ interface{}) (interface{}, error) {
	return nil, errors.New("some error")
}

func okHandler(_ context.Context, _ interface{}) (interface{}, error) {
	return nil, nil
}

func TestMiddleware_Log(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		m := New()

		fn := m.Log()
		_, err := fn(context.Background(), nil, &grpc.UnaryServerInfo{}, errorHandler)
		assert.Error(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		m := New()

		fn := m.Log()
		_, err := fn(context.Background(), nil, &grpc.UnaryServerInfo{}, okHandler)
		assert.NoError(t, err)
	})
}
