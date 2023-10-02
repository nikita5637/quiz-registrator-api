package authentication

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewContextWithServiceAuth(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctx := context.Background()
		ctx = NewContextWithServiceAuth(ctx)
		assert.True(t, IsServiceAuth(ctx))
	})
}

func TestIsServiceAuth(t *testing.T) {
	t.Run("no key", func(t *testing.T) {
		ctx := context.Background()
		assert.False(t, IsServiceAuth(ctx))
	})

	t.Run("no bool", func(t *testing.T) {
		ctx := context.Background()
		ctx = context.WithValue(ctx, serviceAuthKey, struct{}{})
		assert.False(t, IsServiceAuth(ctx))
	})

	t.Run("false", func(t *testing.T) {
		ctx := context.Background()
		ctx = context.WithValue(ctx, serviceAuthKey, false)
		assert.False(t, IsServiceAuth(ctx))
	})
}
