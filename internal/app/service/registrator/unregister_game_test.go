package registrator

import (
	"errors"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestRegistrator_UnregisterGame(t *testing.T) {
	t.Run("internal error while unregister game", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().UnregisterGame(fx.ctx, int32(1)).Return(model.UnregisterGameStatusInvalid, errors.New("some error"))

		got, err := fx.implementation.UnregisterGame(fx.ctx, &registrator.UnregisterGameRequest{
			GameId: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("error game not found while unregister game", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().UnregisterGame(fx.ctx, int32(1)).Return(model.UnregisterGameStatusInvalid, games.ErrGameNotFound)

		got, err := fx.implementation.UnregisterGame(fx.ctx, &registrator.UnregisterGameRequest{
			GameId: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().UnregisterGame(fx.ctx, int32(1)).Return(model.UnregisterGameStatusOK, nil)

		got, err := fx.implementation.UnregisterGame(fx.ctx, &registrator.UnregisterGameRequest{
			GameId: 1,
		})
		assert.NotNil(t, got)
		assert.Equal(t, &registrator.UnregisterGameResponse{
			Status: registrator.UnregisterGameStatus_UNREGISTER_GAME_STATUS_OK,
		}, got)
		assert.NoError(t, err)
	})
}
