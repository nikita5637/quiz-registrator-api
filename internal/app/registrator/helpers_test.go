package registrator

import (
	"context"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/app/registrator/mocks"
)

type fixture struct {
	ctx         context.Context
	registrator *Registrator

	gamesFacade      *mocks.GamesFacade
	gamePhotosFacade *mocks.GamePhotosFacade
	leaguesFacade    *mocks.LeaguesFacade
	placesFacade     *mocks.PlacesFacade
	usersFacade      *mocks.UsersFacade
}

func tearUp(t *testing.T) *fixture {
	fx := &fixture{
		ctx: context.Background(),

		gamesFacade:      mocks.NewGamesFacade(t),
		gamePhotosFacade: mocks.NewGamePhotosFacade(t),
		leaguesFacade:    mocks.NewLeaguesFacade(t),
		placesFacade:     mocks.NewPlacesFacade(t),
		usersFacade:      mocks.NewUsersFacade(t),
	}

	fx.registrator = &Registrator{
		gamesFacade:      fx.gamesFacade,
		gamePhotosFacade: fx.gamePhotosFacade,
		leaguesFacade:    fx.leaguesFacade,
		placesFacade:     fx.placesFacade,
		usersFacade:      fx.usersFacade,
	}

	t.Cleanup(func() {})

	return fx
}
