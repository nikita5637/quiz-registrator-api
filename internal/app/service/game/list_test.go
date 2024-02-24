package game

import (
	"errors"
	"testing"
	"time"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	gamepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestImplementation_ListGames(t *testing.T) {
	t.Run("error: internal error", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().ListGames(fx.ctx).Return(nil, errors.New("some error"))

		got, err := fx.implementation.ListGames(fx.ctx, &emptypb.Empty{})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("ok: empty list", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().ListGames(fx.ctx).Return([]model.Game{}, nil)

		got, err := fx.implementation.ListGames(fx.ctx, &emptypb.Empty{})
		assert.Equal(t, &gamepb.ListGamesResponse{
			Games: []*gamepb.Game{},
		}, got)
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().ListGames(fx.ctx).Return([]model.Game{
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

		got, err := fx.implementation.ListGames(fx.ctx, &emptypb.Empty{})
		assert.Equal(t, &gamepb.ListGamesResponse{
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
