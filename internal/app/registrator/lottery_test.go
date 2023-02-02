package registrator

import (
	"errors"
	"testing"
	"time"

	"github.com/nikita5637/quiz-registrator-api/internal/config"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	registrator "github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"
	time_utils "github.com/nikita5637/quiz-registrator-api/utils/time"

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

		got, err := fx.registrator.GetLotteryStatus(fx.ctx, &registrator.GetLotteryStatusRequest{
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

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{}, model.ErrGameNotFound)

		got, err := fx.registrator.GetLotteryStatus(fx.ctx, &registrator.GetLotteryStatusRequest{
			GameId: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("ok game is not my", func(t *testing.T) {
		time_utils.TimeNow = func() time.Time {
			return time_utils.ConvertTime("2022-02-10 15:31")
		}
		fx := tearUp(t)

		game := model.Game{
			Date:     model.DateTime(time_utils.ConvertTime("2022-02-10 16:30")),
			LeagueID: 1,
		}
		game.My = false

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(game, nil)

		got, err := fx.registrator.GetLotteryStatus(fx.ctx, &registrator.GetLotteryStatusRequest{
			GameId: 1,
		})
		assert.False(t, got.GetActive())
		assert.NoError(t, err)
	})

	t.Run("ok active", func(t *testing.T) {
		time_utils.TimeNow = func() time.Time {
			return time_utils.ConvertTime("2022-02-10 15:31")
		}
		fx := tearUp(t)

		game := model.Game{
			Date:     model.DateTime(time_utils.ConvertTime("2022-02-10 16:30")),
			LeagueID: 1,
		}
		game.My = true

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(game, nil)

		fx.croupier.EXPECT().GetIsLotteryActive(fx.ctx, game).Return(true)

		got, err := fx.registrator.GetLotteryStatus(fx.ctx, &registrator.GetLotteryStatusRequest{
			GameId: 1,
		})
		assert.True(t, got.GetActive())
		assert.NoError(t, err)
	})

	t.Run("ok not active", func(t *testing.T) {
		time_utils.TimeNow = func() time.Time {
			return time_utils.ConvertTime("2022-02-10 15:30")
		}
		fx := tearUp(t)

		game := model.Game{
			Date:     model.DateTime(time_utils.ConvertTime("2022-02-10 16:30")),
			LeagueID: 1,
		}
		game.My = true

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(game, nil)

		fx.croupier.EXPECT().GetIsLotteryActive(fx.ctx, game).Return(false)

		got, err := fx.registrator.GetLotteryStatus(fx.ctx, &registrator.GetLotteryStatusRequest{
			GameId: 1,
		})
		assert.False(t, got.GetActive())
		assert.NoError(t, err)
	})
}

func TestRegistrator_RegisterForLottery(t *testing.T) {
	t.Run("error while get game", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{}, errors.New("some error"))

		got, err := fx.registrator.RegisterForLottery(fx.ctx, &registrator.RegisterForLotteryRequest{
			GameId: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Len(t, st.Details(), 0)
	})

	t.Run("game not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{}, model.ErrGameNotFound)

		got, err := fx.registrator.RegisterForLottery(fx.ctx, &registrator.RegisterForLotteryRequest{
			GameId: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Len(t, st.Details(), 2)
	})

	t.Run("error lottery not available", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			ID: 1,
		}, nil)

		fx.croupier.EXPECT().RegisterForLottery(fx.ctx, model.Game{
			ID: 1,
		}, model.User{}).Return(0, model.ErrLotteryNotAvailable)

		got, err := fx.registrator.RegisterForLottery(fx.ctx, &registrator.RegisterForLotteryRequest{
			GameId: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("error lottery not implemented", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			ID: 1,
		}, nil)

		fx.croupier.EXPECT().RegisterForLottery(fx.ctx, model.Game{
			ID: 1,
		}, model.User{}).Return(0, model.ErrLotteryNotImplemented)

		got, err := fx.registrator.RegisterForLottery(fx.ctx, &registrator.RegisterForLotteryRequest{
			GameId: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Unimplemented, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("error lottery permission denied", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			ID: 1,
		}, nil)

		fx.croupier.EXPECT().RegisterForLottery(fx.ctx, model.Game{
			ID: 1,
		}, model.User{}).Return(0, model.ErrLotteryPermissionDenied)

		got, err := fx.registrator.RegisterForLottery(fx.ctx, &registrator.RegisterForLotteryRequest{
			GameId: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("other error while registration", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			ID: 1,
		}, nil)

		fx.croupier.EXPECT().RegisterForLottery(fx.ctx, model.Game{
			ID: 1,
		}, model.User{}).Return(0, errors.New("some error"))

		got, err := fx.registrator.RegisterForLottery(fx.ctx, &registrator.RegisterForLotteryRequest{
			GameId: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("ok with number", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			ID: 1,
		}, nil)

		fx.croupier.EXPECT().RegisterForLottery(fx.ctx, model.Game{
			ID: 1,
		}, model.User{}).Return(100, nil)

		got, err := fx.registrator.RegisterForLottery(fx.ctx, &registrator.RegisterForLotteryRequest{
			GameId: 1,
		})
		assert.Equal(t, &registrator.RegisterForLotteryResponse{
			Number: 100,
		}, got)
		assert.NoError(t, err)
	})

	t.Run("ok without number", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			ID: 1,
		}, nil)

		fx.croupier.EXPECT().RegisterForLottery(fx.ctx, model.Game{
			ID: 1,
		}, model.User{}).Return(0, nil)

		got, err := fx.registrator.RegisterForLottery(fx.ctx, &registrator.RegisterForLotteryRequest{
			GameId: 1,
		})
		assert.Equal(t, &registrator.RegisterForLotteryResponse{
			Number: 0,
		}, got)
		assert.NoError(t, err)
	})
}
