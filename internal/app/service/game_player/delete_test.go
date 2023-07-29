package gameplayer

import (
	"errors"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/gameplayers"
	gameplayer "github.com/nikita5637/quiz-registrator-api/pkg/pb/game_player"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestImplementation_DeleteGamePlayer(t *testing.T) {
	t.Run("game player not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePlayersFacade.EXPECT().DeleteGamePlayer(fx.ctx, int32(1)).Return(gameplayers.ErrGamePlayerNotFound)

		_, err := fx.implementation.DeleteGamePlayer(fx.ctx, &gameplayer.DeleteGamePlayerRequest{
			Id: 1,
		})
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())
		assert.Len(t, st.Details(), 2)

		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, gameplayers.ReasonGamePlayerNotFound, errorInfo.Reason)
		assert.Nil(t, errorInfo.Metadata)
	})

	t.Run("internal error", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePlayersFacade.EXPECT().DeleteGamePlayer(fx.ctx, int32(1)).Return(errors.New("some error"))

		_, err := fx.implementation.DeleteGamePlayer(fx.ctx, &gameplayer.DeleteGamePlayerRequest{
			Id: 1,
		})
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePlayersFacade.EXPECT().DeleteGamePlayer(fx.ctx, int32(1)).Return(nil)

		_, err := fx.implementation.DeleteGamePlayer(fx.ctx, &gameplayer.DeleteGamePlayerRequest{
			Id: 1,
		})
		assert.NoError(t, err)
	})
}
