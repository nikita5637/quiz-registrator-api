package game

import (
	"errors"
	"testing"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/ics"
	gamepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestImplementation_UnUnregisterGame(t *testing.T) {
	t.Run("error: game not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{}, games.ErrGameNotFound)

		got, err := fx.implementation.UnregisterGame(fx.ctx, &gamepb.UnregisterGameRequest{
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

	t.Run("error: internal error while getting game", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{}, errors.New("some error"))

		got, err := fx.implementation.UnregisterGame(fx.ctx, &gamepb.UnregisterGameRequest{
			Id: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("error: game has passed", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			HasPassed: true,
		}, nil)

		got, err := fx.implementation.UnregisterGame(fx.ctx, &gamepb.UnregisterGameRequest{
			Id: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.FailedPrecondition, st.Code())
		assert.Len(t, st.Details(), 2)

		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, "GAME_HAS_PASSED", errorInfo.Reason)
		assert.Nil(t, errorInfo.Metadata)
	})

	t.Run("error: internal error while patch game", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			ID:        1,
			HasPassed: false,
		}, nil)

		fx.gamesFacade.EXPECT().PatchGame(fx.ctx, model.Game{
			ID:         1,
			Payment:    maybe.Nothing[model.Payment](),
			Registered: false,
			HasPassed:  false,
		}).Return(model.Game{}, errors.New("some error"))

		got, err := fx.implementation.UnregisterGame(fx.ctx, &gamepb.UnregisterGameRequest{
			Id: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("error: rabbitMQ message sending error", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			ID:        1,
			HasPassed: false,
		}, nil)

		fx.gamesFacade.EXPECT().PatchGame(fx.ctx, model.Game{
			ID:         1,
			Payment:    maybe.Nothing[model.Payment](),
			Registered: false,
			HasPassed:  false,
		}).Return(model.Game{}, nil)

		fx.rabbitMQProducer.EXPECT().Send(fx.ctx, ics.Event{
			GameID: 1,
			Event:  ics.EventUnregistered,
		}).Return(errors.New("some error"))

		got, err := fx.implementation.UnregisterGame(fx.ctx, &gamepb.UnregisterGameRequest{
			Id: 1,
		})
		assert.Equal(t, &emptypb.Empty{}, got)
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			ID:        1,
			HasPassed: false,
		}, nil)

		fx.gamesFacade.EXPECT().PatchGame(fx.ctx, model.Game{
			ID:         1,
			Payment:    maybe.Nothing[model.Payment](),
			Registered: false,
			HasPassed:  false,
		}).Return(model.Game{}, nil)

		fx.rabbitMQProducer.EXPECT().Send(fx.ctx, ics.Event{
			GameID: 1,
			Event:  ics.EventUnregistered,
		}).Return(nil)

		got, err := fx.implementation.UnregisterGame(fx.ctx, &gamepb.UnregisterGameRequest{
			Id: 1,
		})
		assert.Equal(t, &emptypb.Empty{}, got)
		assert.NoError(t, err)
	})
}
