package errorwrap

import (
	"context"
	"testing"
)

type fixture struct {
	ctx context.Context

	middleware *Middleware
}

func tearUp(t *testing.T) *fixture {
	fx := &fixture{
		ctx: context.Background(),
	}

	fx.middleware = New()

	t.Cleanup(func() {})

	return fx
}
