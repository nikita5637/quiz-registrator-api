package croupier

import (
	"errors"
	"testing"
	"time"

	"github.com/nikita5637/quiz-registrator-api/internal/config"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	croupierpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/croupier"
	timeutils "github.com/nikita5637/quiz-registrator-api/utils/time"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestRegistrator_GetLotteryStatus(t *testing.T) {
	globalConfig := config.GlobalConfig{}
	globalConfig.LotteryStartsBefore = 3600
	config.UpdateGlobalConfig(globalConfig)
	t.Run("internal error while get game", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{}, errors.New("some error"))

		got, err := fx.implementation.GetLotteryStatus(fx.ctx, &croupierpb.GetLotteryStatusRequest{
			GameId: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("error game not found while get game", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{}, games.ErrGameNotFound)

		got, err := fx.implementation.GetLotteryStatus(fx.ctx, &croupierpb.GetLotteryStatusRequest{
			GameId: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("error game has passed while get game", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{}, games.ErrGameHasPassed)

		got, err := fx.implementation.GetLotteryStatus(fx.ctx, &croupierpb.GetLotteryStatusRequest{
			GameId: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("ok game is not my", func(t *testing.T) {
		timeutils.TimeNow = func() time.Time {
			return timeutils.ConvertTime("2022-02-10 15:31")
		}
		fx := tearUp(t)

		game := model.Game{
			Date:     model.DateTime(timeutils.ConvertTime("2022-02-10 16:30")),
			LeagueID: 1,
		}
		game.My = false

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(game, nil)

		got, err := fx.implementation.GetLotteryStatus(fx.ctx, &croupierpb.GetLotteryStatusRequest{
			GameId: 1,
		})
		assert.False(t, got.GetActive())
		assert.NoError(t, err)
	})

	t.Run("ok active", func(t *testing.T) {
		timeutils.TimeNow = func() time.Time {
			return timeutils.ConvertTime("2022-02-10 15:31")
		}
		fx := tearUp(t)

		game := model.Game{
			Date:     model.DateTime(timeutils.ConvertTime("2022-02-10 16:30")),
			LeagueID: 1,
		}
		game.My = true

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(game, nil)

		fx.croupier.EXPECT().GetIsLotteryActive(fx.ctx, game).Return(true)

		got, err := fx.implementation.GetLotteryStatus(fx.ctx, &croupierpb.GetLotteryStatusRequest{
			GameId: 1,
		})
		assert.True(t, got.GetActive())
		assert.NoError(t, err)
	})

	t.Run("ok not active", func(t *testing.T) {
		timeutils.TimeNow = func() time.Time {
			return timeutils.ConvertTime("2022-02-10 15:30")
		}
		fx := tearUp(t)

		game := model.Game{
			Date:     model.DateTime(timeutils.ConvertTime("2022-02-10 16:30")),
			LeagueID: 1,
		}
		game.My = true

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(game, nil)

		fx.croupier.EXPECT().GetIsLotteryActive(fx.ctx, game).Return(false)

		got, err := fx.implementation.GetLotteryStatus(fx.ctx, &croupierpb.GetLotteryStatusRequest{
			GameId: 1,
		})
		assert.False(t, got.GetActive())
		assert.NoError(t, err)
	})
}
