package league

import (
	"context"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/app/service/league/mocks"
)

type fixture struct {
	ctx context.Context

	leaguesFacade *mocks.LeaguesFacade

	implementation *Implementation
}

func tearUp(t *testing.T) *fixture {
	fx := &fixture{
		ctx: context.Background(),

		leaguesFacade: mocks.NewLeaguesFacade(t),
	}

	fx.implementation = &Implementation{
		leaguesFacade: fx.leaguesFacade,
	}

	t.Cleanup(func() {})

	return fx
}
