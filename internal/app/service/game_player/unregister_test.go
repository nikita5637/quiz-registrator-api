package gameplayer

import (
	"errors"
	"testing"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	gameplayer "github.com/nikita5637/quiz-registrator-api/pkg/pb/game_player"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestImplementation_UnregisterPlayer(t *testing.T) {
	t.Run("game player is nil", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.implementation.UnregisterPlayer(fx.ctx, &gameplayer.UnregisterPlayerRequest{})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("validation error", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.implementation.UnregisterPlayer(fx.ctx, &gameplayer.UnregisterPlayerRequest{
			GamePlayer: &gameplayer.GamePlayer{
				RegisteredBy: 1,
				Degree:       gameplayer.Degree_DEGREE_LIKELY,
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)

		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, invalidGameIDReason, errorInfo.Reason)
		assert.Equal(t, map[string]string{
			"error": "GameID: cannot be blank.",
		}, errorInfo.Metadata)
	})

	t.Run("game not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{}, games.ErrGameNotFound)

		got, err := fx.implementation.UnregisterPlayer(fx.ctx, &gameplayer.UnregisterPlayerRequest{
			GamePlayer: &gameplayer.GamePlayer{
				GameId:       1,
				RegisteredBy: 1,
				Degree:       gameplayer.Degree_DEGREE_LIKELY,
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.FailedPrecondition, st.Code())
		assert.Len(t, st.Details(), 2)

		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, games.ReasonGameNotFound, errorInfo.Reason)
		assert.Equal(t, map[string]string{
			"error": "game not found",
		}, errorInfo.Metadata)
	})

	t.Run("internal error while getting game", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{}, errors.New("some error"))

		got, err := fx.implementation.UnregisterPlayer(fx.ctx, &gameplayer.UnregisterPlayerRequest{
			GamePlayer: &gameplayer.GamePlayer{
				GameId:       1,
				RegisteredBy: 1,
				Degree:       gameplayer.Degree_DEGREE_LIKELY,
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("game has passed", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			HasPassed: true,
		}, nil)

		got, err := fx.implementation.UnregisterPlayer(fx.ctx, &gameplayer.UnregisterPlayerRequest{
			GamePlayer: &gameplayer.GamePlayer{
				GameId:       1,
				RegisteredBy: 1,
				Degree:       gameplayer.Degree_DEGREE_LIKELY,
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.FailedPrecondition, st.Code())
		assert.Len(t, st.Details(), 2)

		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, games.ReasonGameHasPassed, errorInfo.Reason)
		assert.Nil(t, errorInfo.Metadata)
	})

	t.Run("get game players error", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			HasPassed: false,
		}, nil)

		fx.gamePlayersFacade.EXPECT().GetGamePlayersByFields(fx.ctx, model.GamePlayer{
			GameID:       1,
			UserID:       maybe.Nothing[int32](),
			RegisteredBy: 1,
			Degree:       model.DegreeLikely,
		}).Return(nil, errors.New("some error"))

		got, err := fx.implementation.UnregisterPlayer(fx.ctx, &gameplayer.UnregisterPlayerRequest{
			GamePlayer: &gameplayer.GamePlayer{
				GameId:       1,
				RegisteredBy: 1,
				Degree:       gameplayer.Degree_DEGREE_LIKELY,
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("game players empty list", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			HasPassed: false,
		}, nil)

		fx.gamePlayersFacade.EXPECT().GetGamePlayersByFields(fx.ctx, model.GamePlayer{
			GameID:       1,
			UserID:       maybe.Nothing[int32](),
			RegisteredBy: 1,
			Degree:       model.DegreeLikely,
		}).Return([]model.GamePlayer{}, nil)

		got, err := fx.implementation.UnregisterPlayer(fx.ctx, &gameplayer.UnregisterPlayerRequest{
			GamePlayer: &gameplayer.GamePlayer{
				GameId:       1,
				RegisteredBy: 1,
				Degree:       gameplayer.Degree_DEGREE_LIKELY,
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())
		assert.Len(t, st.Details(), 2)

		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, reasonThereAreNoSuitablePlayers, errorInfo.Reason)
		assert.Nil(t, errorInfo.Metadata)
	})

	t.Run("internal error while deleting game player", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{}, nil)

		fx.gamePlayersFacade.EXPECT().GetGamePlayersByFields(fx.ctx, model.GamePlayer{
			GameID:       1,
			UserID:       maybe.Nothing[int32](),
			RegisteredBy: 1,
			Degree:       model.DegreeLikely,
		}).Return([]model.GamePlayer{
			{
				ID:           1,
				GameID:       1,
				UserID:       maybe.Nothing[int32](),
				RegisteredBy: 1,
				Degree:       model.DegreeLikely,
			},
			{
				ID:           2,
				GameID:       1,
				UserID:       maybe.Nothing[int32](),
				RegisteredBy: 1,
				Degree:       model.DegreeLikely,
			},
		}, nil)

		fx.gamePlayersFacade.EXPECT().DeleteGamePlayer(fx.ctx, int32(1)).Return(errors.New("some error"))

		got, err := fx.implementation.UnregisterPlayer(fx.ctx, &gameplayer.UnregisterPlayerRequest{
			GamePlayer: &gameplayer.GamePlayer{
				GameId:       1,
				RegisteredBy: 1,
				Degree:       gameplayer.Degree_DEGREE_LIKELY,
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{}, nil)

		fx.gamePlayersFacade.EXPECT().GetGamePlayersByFields(fx.ctx, model.GamePlayer{
			GameID:       1,
			UserID:       maybe.Nothing[int32](),
			RegisteredBy: 1,
			Degree:       model.DegreeLikely,
		}).Return([]model.GamePlayer{
			{
				ID:           1,
				GameID:       1,
				UserID:       maybe.Nothing[int32](),
				RegisteredBy: 1,
				Degree:       model.DegreeLikely,
			},
			{
				ID:           2,
				GameID:       1,
				UserID:       maybe.Nothing[int32](),
				RegisteredBy: 1,
				Degree:       model.DegreeLikely,
			},
		}, nil)

		fx.gamePlayersFacade.EXPECT().DeleteGamePlayer(fx.ctx, int32(1)).Return(nil)

		got, err := fx.implementation.UnregisterPlayer(fx.ctx, &gameplayer.UnregisterPlayerRequest{
			GamePlayer: &gameplayer.GamePlayer{
				GameId:       1,
				RegisteredBy: 1,
				Degree:       gameplayer.Degree_DEGREE_LIKELY,
			},
		})
		assert.Equal(t, &emptypb.Empty{}, got)
		assert.NoError(t, err)
	})
}
