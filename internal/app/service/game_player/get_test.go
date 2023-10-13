package gameplayer

import (
	"errors"
	"testing"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/gameplayers"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	gameplayer "github.com/nikita5637/quiz-registrator-api/pkg/pb/game_player"
	gameplayerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game_player"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestImplementation_GetGamePlayer(t *testing.T) {
	t.Run("game player not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePlayersFacade.EXPECT().GetGamePlayer(fx.ctx, int32(1)).Return(model.GamePlayer{}, gameplayers.ErrGamePlayerNotFound)

		got, err := fx.implementation.GetGamePlayer(fx.ctx, &gameplayerpb.GetGamePlayerRequest{
			Id: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())
		assert.Len(t, st.Details(), 2)

		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, gameplayers.ReasonGamePlayerNotFound, errorInfo.Reason)
		assert.Equal(t, map[string]string{
			"error": "game player not found",
		}, errorInfo.Metadata)
	})

	t.Run("internal error", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePlayersFacade.EXPECT().GetGamePlayer(fx.ctx, int32(1)).Return(model.GamePlayer{}, errors.New("some error"))

		got, err := fx.implementation.GetGamePlayer(fx.ctx, &gameplayerpb.GetGamePlayerRequest{
			Id: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePlayersFacade.EXPECT().GetGamePlayer(fx.ctx, int32(1)).Return(model.GamePlayer{
			ID:           1,
			GameID:       1,
			UserID:       maybe.Just(int32(1)),
			RegisteredBy: 1,
			Degree:       model.DegreeLikely,
		}, nil)

		got, err := fx.implementation.GetGamePlayer(fx.ctx, &gameplayerpb.GetGamePlayerRequest{
			Id: 1,
		})
		assert.Equal(t, &gameplayerpb.GamePlayer{
			Id:     1,
			GameId: 1,
			UserId: &wrapperspb.Int32Value{
				Value: 1,
			},
			RegisteredBy: 1,
			Degree:       gameplayer.Degree_DEGREE_LIKELY,
		}, got)
		assert.NoError(t, err)
	})
}

func TestImplementation_GetGamePlayersByGameID(t *testing.T) {
	t.Run("internal error", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePlayersFacade.EXPECT().GetGamePlayersByGameID(fx.ctx, int32(1)).Return(nil, errors.New("some error"))

		got, err := fx.implementation.GetGamePlayersByGameID(fx.ctx, &gameplayer.GetGamePlayersByGameIDRequest{
			GameId: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePlayersFacade.EXPECT().GetGamePlayersByGameID(fx.ctx, int32(1)).Return([]model.GamePlayer{
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
				UserID:       maybe.Just(int32(1)),
				RegisteredBy: 1,
				Degree:       model.DegreeUnlikely,
			},
		}, nil)

		got, err := fx.implementation.GetGamePlayersByGameID(fx.ctx, &gameplayer.GetGamePlayersByGameIDRequest{
			GameId: 1,
		})
		assert.Equal(t, &gameplayerpb.GetGamePlayersByGameIDResponse{
			GamePlayers: []*gameplayerpb.GamePlayer{
				{
					Id:           1,
					GameId:       1,
					RegisteredBy: 1,
					Degree:       gameplayer.Degree_DEGREE_LIKELY,
				},
				{
					Id:     2,
					GameId: 1,
					UserId: &wrapperspb.Int32Value{
						Value: 1,
					},
					RegisteredBy: 1,
					Degree:       gameplayer.Degree_DEGREE_UNLIKELY,
				},
			},
		}, got)
		assert.NoError(t, err)
	})
}

func TestImplementation_GetUserGameIDs(t *testing.T) {
	t.Run("error: internal error", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePlayersFacade.EXPECT().GetUserGameIDs(fx.ctx, int32(1)).Return(nil, errors.New("some error"))

		got, err := fx.implementation.GetUserGameIDs(fx.ctx, &gameplayerpb.GetUserGameIDsRequest{
			UserId: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePlayersFacade.EXPECT().GetUserGameIDs(fx.ctx, int32(1)).Return([]int32{
			1,
			2,
			3,
		}, nil)

		got, err := fx.implementation.GetUserGameIDs(fx.ctx, &gameplayerpb.GetUserGameIDsRequest{
			UserId: 1,
		})
		assert.Equal(t, &gameplayerpb.GetUserGameIDsResponse{
			GameIds: []int32{
				1,
				2,
				3,
			},
		}, got)
		assert.NoError(t, err)
	})
}
