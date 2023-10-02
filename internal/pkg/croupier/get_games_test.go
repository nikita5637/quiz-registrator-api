package croupier

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/nikita5637/quiz-registrator-api/internal/config"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/croupier/mocks"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	leaguepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/league"
	timeutils "github.com/nikita5637/quiz-registrator-api/utils/time"
	"github.com/stretchr/testify/assert"
)

func TestCroupier_GetGamesWithActiveLottery(t *testing.T) {
	timeutils.TimeNow = func() time.Time {
		return timeutils.ConvertTime("2022-02-10 15:22")
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
				LeagueID: int32(leaguepb.LeagueID_SHAKER),
			},
			{
				ID:       2,
				Date:     model.DateTime(timeutils.ConvertTime("2022-02-10 16:00")),
				LeagueID: int32(leaguepb.LeagueID_SHAKER),
			},
			{
				ID:       3,
				Date:     model.DateTime(timeutils.ConvertTime("2022-02-10 14:30")),
				LeagueID: int32(leaguepb.LeagueID_SHAKER),
			},
			{
				ID:       4,
				Date:     model.DateTime(timeutils.ConvertTime("2022-02-10 16:30")),
				LeagueID: int32(leaguepb.LeagueID_SHAKER),
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
				LeagueID:   int32(leaguepb.LeagueID_SQUIZ),
				Registered: true,
				HasPassed:  true,
			},
			{
				ID:         2,
				Date:       model.DateTime(timeutils.ConvertTime("2022-02-10 16:00")),
				LeagueID:   int32(leaguepb.LeagueID_QUIZ_PLEASE),
				Registered: true,
				HasPassed:  false,
			},
			{
				ID:         3,
				Date:       model.DateTime(timeutils.ConvertTime("2022-02-10 14:30")),
				LeagueID:   int32(leaguepb.LeagueID_SQUIZ),
				Registered: true,
				HasPassed:  false,
			},
			{
				ID:         4,
				Date:       model.DateTime(timeutils.ConvertTime("2022-02-10 14:30")),
				LeagueID:   int32(leaguepb.LeagueID_SQUIZ),
				Registered: false,
				HasPassed:  false,
			},
			{
				ID:         5,
				Date:       model.DateTime(timeutils.ConvertTime("2022-02-10 16:30")),
				LeagueID:   int32(leaguepb.LeagueID_QUIZ_PLEASE),
				Registered: true,
				HasPassed:  false,
			},
		}, nil)

		got, err := c.GetGamesWithActiveLottery(ctx)
		assert.Equal(t, []model.Game{
			{
				ID:         2,
				Date:       model.DateTime(timeutils.ConvertTime("2022-02-10 16:00")),
				LeagueID:   int32(leaguepb.LeagueID_QUIZ_PLEASE),
				Registered: true,
			},
			{
				ID:         3,
				Date:       model.DateTime(timeutils.ConvertTime("2022-02-10 14:30")),
				LeagueID:   int32(leaguepb.LeagueID_SQUIZ),
				Registered: true,
			},
		}, got)
		assert.NoError(t, err)
	})
}
