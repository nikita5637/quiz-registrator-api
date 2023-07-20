package gameresultmanager

import (
	"errors"
	"testing"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	gameresultmanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game_result_manager"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestRegistrator_ListGameResults(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameResultsFacade.EXPECT().ListGameResults(fx.ctx).Return(nil, errors.New("some error"))

		got, err := fx.gameResultManager.ListGameResults(fx.ctx, nil)

		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameResultsFacade.EXPECT().ListGameResults(fx.ctx).Return(
			[]model.GameResult{
				{
					ID:          1,
					FkGameID:    2,
					ResultPlace: 1,
					RoundPoints: maybe.Just("{}"),
				},
				{
					ID:          2,
					FkGameID:    3,
					ResultPlace: 2,
					RoundPoints: maybe.Just("{}"),
				},
			},
			nil)

		got, err := fx.gameResultManager.ListGameResults(fx.ctx, nil)

		assert.Equal(t, got, &gameresultmanagerpb.ListGameResultsResponse{
			GameResults: []*gameresultmanagerpb.GameResult{
				{
					Id:          1,
					GameId:      2,
					ResultPlace: 1,
					RoundPoints: "{}",
				},
				{
					Id:          2,
					GameId:      3,
					ResultPlace: 2,
					RoundPoints: "{}",
				},
			},
		})
		assert.NoError(t, err)
	})
}
