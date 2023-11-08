package croupier

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/croupier/mocks"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	timeutils "github.com/nikita5637/quiz-registrator-api/utils/time"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestCroupier_GetGamesWithActiveLottery(t *testing.T) {
	timeutils.TimeNow = func() time.Time {
		return timeutils.ConvertTime("2022-02-10 15:22")
	}

	viper.Set("croupier.lottery_starts_before", 3600)

	t.Run("error while get games", func(t *testing.T) {
		ctx := context.Background()
		gamesFacadeMock := mocks.NewGamesFacade(t)
		c := New(Config{
			GamesFacade: gamesFacadeMock,
		})

		gamesFacadeMock.EXPECT().ListGames(ctx).Return(nil, errors.New("some error"))

		got, err := c.GetGamesWithActiveLottery(ctx)
		assert.Nil(t, got)
		assert.Error(t, err)
	})

	t.Run("league not implemented", func(t *testing.T) {
		ctx := context.Background()
		gamesFacadeMock := mocks.NewGamesFacade(t)
		c := New(Config{
			GamesFacade: gamesFacadeMock,
		})

		gamesFacadeMock.EXPECT().ListGames(ctx).Return([]model.Game{
			{
				ID:       1,
				Date:     model.DateTime(timeutils.ConvertTime("2022-02-10 13:00")),
				LeagueID: 4,
			},
			{
				ID:       2,
				Date:     model.DateTime(timeutils.ConvertTime("2022-02-10 16:00")),
				LeagueID: 4,
			},
			{
				ID:       3,
				Date:     model.DateTime(timeutils.ConvertTime("2022-02-10 14:30")),
				LeagueID: 4,
			},
			{
				ID:       4,
				Date:     model.DateTime(timeutils.ConvertTime("2022-02-10 16:30")),
				LeagueID: 4,
			},
		}, nil)

		got, err := c.GetGamesWithActiveLottery(ctx)
		assert.Equal(t, []model.Game{}, got)
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		ctx := context.Background()
		gamesFacadeMock := mocks.NewGamesFacade(t)
		c := New(Config{
			GamesFacade: gamesFacadeMock,
		})

		gamesFacadeMock.EXPECT().ListGames(ctx).Return([]model.Game{
			{
				ID:         1,
				Date:       model.DateTime(timeutils.ConvertTime("2022-02-10 13:00")),
				LeagueID:   model.LeagueSquiz,
				Registered: true,
				HasPassed:  true,
			},
			{
				ID:         2,
				Date:       model.DateTime(timeutils.ConvertTime("2022-02-10 16:00")),
				LeagueID:   model.LeagueQuizPlease,
				Registered: true,
				HasPassed:  false,
			},
			{
				ID:         3,
				Date:       model.DateTime(timeutils.ConvertTime("2022-02-10 14:30")),
				LeagueID:   model.LeagueSquiz,
				Registered: true,
				HasPassed:  false,
			},
			{
				ID:         4,
				Date:       model.DateTime(timeutils.ConvertTime("2022-02-10 14:30")),
				LeagueID:   model.LeagueSquiz,
				Registered: false,
				HasPassed:  false,
			},
			{
				ID:         5,
				Date:       model.DateTime(timeutils.ConvertTime("2022-02-10 16:30")),
				LeagueID:   model.LeagueQuizPlease,
				Registered: true,
				HasPassed:  false,
			},
		}, nil)

		got, err := c.GetGamesWithActiveLottery(ctx)
		assert.Equal(t, []model.Game{
			{
				ID:         2,
				Date:       model.DateTime(timeutils.ConvertTime("2022-02-10 16:00")),
				LeagueID:   model.LeagueQuizPlease,
				Registered: true,
			},
			{
				ID:         3,
				Date:       model.DateTime(timeutils.ConvertTime("2022-02-10 14:30")),
				LeagueID:   model.LeagueSquiz,
				Registered: true,
			},
		}, got)
		assert.NoError(t, err)
	})
}
