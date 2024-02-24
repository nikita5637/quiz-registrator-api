package game

import (
	"errors"
	"testing"
	"time"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	gamepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestImplementation_BatchGetGames(t *testing.T) {
	t.Run("error: internal error", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGamesByIDs(fx.ctx, []int32{
			1,
			2,
			3,
		}).Return(nil, errors.New("some error"))

		got, err := fx.implementation.BatchGetGames(fx.ctx, &gamepb.BatchGetGamesRequest{
			Ids: []int32{
				1,
				2,
				3,
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("ok: empty list", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGamesByIDs(fx.ctx, []int32{
			1,
			2,
			3,
		}).Return([]model.Game{}, nil)

		got, err := fx.implementation.BatchGetGames(fx.ctx, &gamepb.BatchGetGamesRequest{
			Ids: []int32{
				1,
				2,
				3,
			},
		})
		assert.Equal(t, &gamepb.BatchGetGamesResponse{
			Games: []*gamepb.Game{},
		}, got)
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGamesByIDs(fx.ctx, []int32{
			1,
			2,
			3,
		}).Return([]model.Game{
			{
				ID:          1,
				ExternalID:  maybe.Nothing[int32](),
				Name:        maybe.Nothing[string](),
				Date:        model.DateTime(time.Unix(0, 0)),
				PaymentType: maybe.Nothing[string](),
				Payment:     maybe.Nothing[model.Payment](),
				GameLink:    maybe.Nothing[string](),
			},
		}, nil)

		got, err := fx.implementation.BatchGetGames(fx.ctx, &gamepb.BatchGetGamesRequest{
			Ids: []int32{
				1,
				2,
				3,
			},
		})
		assert.Equal(t, &gamepb.BatchGetGamesResponse{
			Games: []*gamepb.Game{
				{
					Id:   1,
					Date: &timestamppb.Timestamp{},
				},
			},
		}, got)
		assert.NoError(t, err)
	})
}

func TestImplementation_GetGame(t *testing.T) {
	t.Run("error: internal", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{}, errors.New("some error"))

		got, err := fx.implementation.GetGame(fx.ctx, &gamepb.GetGameRequest{
			Id: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("error: game not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{}, games.ErrGameNotFound)

		got, err := fx.implementation.GetGame(fx.ctx, &gamepb.GetGameRequest{
			Id: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())
		assert.Len(t, st.Details(), 2)

		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, "GAME_NOT_FOUND", errorInfo.Reason)
		assert.Equal(t, map[string]string{
			"error": "game not found",
		}, errorInfo.Metadata)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		modelGame := model.NewGame()
		modelGame.ID = 1
		modelGame.Date = model.DateTime(time.Unix(0, 0))
		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(modelGame, nil)

		got, err := fx.implementation.GetGame(fx.ctx, &gamepb.GetGameRequest{
			Id: 1,
		})
		assert.Equal(t, &gamepb.Game{
			Id:   1,
			Date: &timestamppb.Timestamp{},
		}, got)
		assert.NoError(t, err)
	})
}
