package game

import (
	"errors"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	gamepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestImplementation_DeleteGame(t *testing.T) {
	t.Run("error: game not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().DeleteGame(fx.ctx, int32(1)).Return(games.ErrGameNotFound)

		got, err := fx.implementation.DeleteGame(fx.ctx, &gamepb.DeleteGameRequest{
			Id: 1,
		})
		assert.Nil(t, got)

		st := status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())
		assert.Len(t, st.Details(), 2)

		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, "GAME_NOT_FOUND", errorInfo.Reason)
		assert.Nil(t, errorInfo.Metadata)
	})

	t.Run("error: internal", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().DeleteGame(fx.ctx, int32(1)).Return(errors.New("some error"))

		got, err := fx.implementation.DeleteGame(fx.ctx, &gamepb.DeleteGameRequest{
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

		fx.gamesFacade.EXPECT().DeleteGame(fx.ctx, int32(1)).Return(nil)

		got, err := fx.implementation.DeleteGame(fx.ctx, &gamepb.DeleteGameRequest{
			Id: 1,
		})
		assert.NotNil(t, got)
		assert.NoError(t, err)
	})
}
