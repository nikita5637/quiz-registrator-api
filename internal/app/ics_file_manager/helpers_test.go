package icsfilemanager

import (
	"context"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/app/ics_file_manager/mocks"
)

type fixture struct {
	ctx context.Context

	icsFilesFacade *mocks.ICSFilesFacade

	implementation *Implementation
}

func tearUp(t *testing.T) *fixture {
	fx := &fixture{
		ctx: context.Background(),

		icsFilesFacade: mocks.NewICSFilesFacade(t),
	}

	fx.implementation = New(Config{
		ICSFilesFacade: fx.icsFilesFacade,
	})

	t.Cleanup(func() {})

	return fx
}
