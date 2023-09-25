package gameresultmanager

import (
	"errors"
	"testing"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/gameresults"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	gameresultmanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game_result_manager"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGameResultManager_SearchGameResultByGameID(t *testing.T) {
	t.Run("error: game result not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameResultsFacade.EXPECT().SearchGameResultByGameID(fx.ctx, int32(1)).Return(model.GameResult{}, gameresults.ErrGameResultNotFound)

		got, err := fx.gameResultManager.SearchGameResultByGameID(fx.ctx, &gameresultmanagerpb.SearchGameResultByGameIDRequest{
			Id: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())
		assert.Len(t, st.Details(), 2)

		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, "GAME_RESULT_NOT_FOUND", errorInfo.Reason)
		assert.Equal(t, map[string]string{
			"error": "game result not found",
		}, errorInfo.Metadata)
	})

	t.Run("error: internal error while search game result", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameResultsFacade.EXPECT().SearchGameResultByGameID(fx.ctx, int32(1)).Return(model.GameResult{}, errors.New("some error"))

		got, err := fx.gameResultManager.SearchGameResultByGameID(fx.ctx, &gameresultmanagerpb.SearchGameResultByGameIDRequest{
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

		fx.gameResultsFacade.EXPECT().SearchGameResultByGameID(fx.ctx, int32(1)).Return(model.GameResult{
			ID:          1,
			RoundPoints: maybe.Nothing[string](),
		}, nil)

		got, err := fx.gameResultManager.SearchGameResultByGameID(fx.ctx, &gameresultmanagerpb.SearchGameResultByGameIDRequest{
			Id: 1,
		})
		assert.Equal(t, &gameresultmanagerpb.GameResult{
			Id: 1,
		}, got)
		assert.NoError(t, err)
	})
}
