package croupier

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/nikita5637/quiz-registrator-api/internal/config"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/croupier/mocks"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	time_utils "github.com/nikita5637/quiz-registrator-api/utils/time"
	"github.com/stretchr/testify/assert"
)

func TestCroupier_GetGamesWithActiveLottery(t *testing.T) {
	time_utils.TimeNow = func() time.Time {
		return time_utils.ConvertTime("2022-02-10 15:22")
	}

	activeGameLag := uint16(3600)
	assert.Greater(t, activeGameLag, uint16(1))

	lotteryStartsBefore := uint16(3600)
	assert.Greater(t, lotteryStartsBefore, uint16(1))

	cfg := config.GlobalConfig{}
	cfg.ActiveGameLag = activeGameLag
	cfg.LotteryStartsBefore = lotteryStartsBefore

	config.UpdateGlobalConfig(cfg)

	t.Run("error while get games", func(t *testing.T) {
		ctx := context.Background()
		gamesFacadeMock := mocks.NewGamesFacade(t)
		c := New(Config{
			GamesFacade: gamesFacadeMock,
		})

		gamesFacadeMock.EXPECT().GetGames(ctx).Return(nil, errors.New("some error"))

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

		gamesFacadeMock.EXPECT().GetGames(ctx).Return([]model.Game{
			{
				ID:       1,
				Date:     model.DateTime(time_utils.ConvertTime("2022-02-10 13:00")),
				LeagueID: model.LeagueShaker,
			},
			{
				ID:       2,
				Date:     model.DateTime(time_utils.ConvertTime("2022-02-10 16:00")),
				LeagueID: model.LeagueShaker,
			},
			{
				ID:       3,
				Date:     model.DateTime(time_utils.ConvertTime("2022-02-10 14:30")),
				LeagueID: model.LeagueShaker,
			},
			{
				ID:       4,
				Date:     model.DateTime(time_utils.ConvertTime("2022-02-10 16:30")),
				LeagueID: model.LeagueShaker,
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

		gamesFacadeMock.EXPECT().GetGames(ctx).Return([]model.Game{
			{
				ID:         1,
				Date:       model.DateTime(time_utils.ConvertTime("2022-02-10 13:00")),
				LeagueID:   model.LeagueSquiz,
				Registered: true,
			},
			{
				ID:         2,
				Date:       model.DateTime(time_utils.ConvertTime("2022-02-10 16:00")),
				LeagueID:   model.LeagueQuizPlease,
				Registered: true,
			},
			{
				ID:         3,
				Date:       model.DateTime(time_utils.ConvertTime("2022-02-10 14:30")),
				LeagueID:   model.LeagueSquiz,
				Registered: true,
			},
			{
				ID:         4,
				Date:       model.DateTime(time_utils.ConvertTime("2022-02-10 14:30")),
				LeagueID:   model.LeagueSquiz,
				Registered: false,
			},
			{
				ID:         5,
				Date:       model.DateTime(time_utils.ConvertTime("2022-02-10 16:30")),
				LeagueID:   model.LeagueQuizPlease,
				Registered: true,
			},
		}, nil)

		got, err := c.GetGamesWithActiveLottery(ctx)
		assert.Equal(t, []model.Game{
			{
				ID:         2,
				Date:       model.DateTime(time_utils.ConvertTime("2022-02-10 16:00")),
				LeagueID:   model.LeagueQuizPlease,
				Registered: true,
			},
			{
				ID:         3,
				Date:       model.DateTime(time_utils.ConvertTime("2022-02-10 14:30")),
				LeagueID:   model.LeagueSquiz,
				Registered: true,
			},
		}, got)
		assert.NoError(t, err)
	})
}
