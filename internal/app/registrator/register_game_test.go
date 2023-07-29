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

func TestRegistrator_RegisterGame(t *testing.T) {
	t.Run("internal error while register game", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().RegisterGame(fx.ctx, int32(1)).Return(model.RegisterGameStatusInvalid, errors.New("some error"))

		got, err := fx.registrator.RegisterGame(fx.ctx, &registrator.RegisterGameRequest{
			GameId: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("error game not found while register game", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().RegisterGame(fx.ctx, int32(1)).Return(model.RegisterGameStatusInvalid, games.ErrGameNotFound)

		got, err := fx.registrator.RegisterGame(fx.ctx, &registrator.RegisterGameRequest{
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

		fx.gamesFacade.EXPECT().RegisterGame(fx.ctx, int32(1)).Return(model.RegisterGameStatusOK, nil)

		got, err := fx.registrator.RegisterGame(fx.ctx, &registrator.RegisterGameRequest{
			GameId: 1,
		})
		assert.NotNil(t, got)
		assert.Equal(t, &registrator.RegisterGameResponse{
			Status: registrator.RegisterGameStatus_REGISTER_GAME_STATUS_OK,
		}, got)
		assert.NoError(t, err)
	})
}
