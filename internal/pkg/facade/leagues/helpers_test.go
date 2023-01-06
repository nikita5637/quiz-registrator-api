package leagues

import (
	"context"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mocks"
)

type fixture struct {
	ctx    context.Context
	facade *Facade

	leagueStorage *mocks.LeagueStorage
}

func tearUp(t *testing.T) *fixture {
	fx := &fixture{
		ctx: context.Background(),

		leagueStorage: mocks.NewLeagueStorage(t),
	}

	fx.facade = &Facade{
		leagueStorage: fx.leagueStorage,
	}

	t.Cleanup(func() {})

	return fx
}
