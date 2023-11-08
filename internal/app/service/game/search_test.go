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
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestImplementation_SearchGamesByLeagueID(t *testing.T) {
	t.Run("error: internal error", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().SearchGamesByLeagueID(fx.ctx, int32(1), uint64(0), uint64(10)).Return(nil, 0, errors.New("some error"))

		got, err := fx.implementation.SearchGamesByLeagueID(fx.ctx, &gamepb.SearchGamesByLeagueIDRequest{
			Id:       model.LeagueQuizPlease,
			Page:     1,
			PageSize: 10,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("ok: empty list", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().SearchGamesByLeagueID(fx.ctx, int32(1), uint64(0), uint64(10)).Return([]model.Game{}, 0, nil)

		got, err := fx.implementation.SearchGamesByLeagueID(fx.ctx, &gamepb.SearchGamesByLeagueIDRequest{
			Id:       model.LeagueQuizPlease,
			Page:     1,
			PageSize: 10,
		})
		assert.Equal(t, &gamepb.SearchGamesByLeagueIDResponse{
			Games: []*gamepb.Game{},
		}, got)
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().SearchGamesByLeagueID(fx.ctx, int32(1), uint64(0), uint64(10)).Return([]model.Game{
			{
				ID:          1,
				ExternalID:  maybe.Nothing[int32](),
				Name:        maybe.Nothing[string](),
				Date:        model.DateTime(time.Unix(0, 0)),
				PaymentType: maybe.Nothing[string](),
				Payment:     maybe.Nothing[model.Payment](),
				GameLink:    maybe.Nothing[string](),
			},
			{
				ID:          2,
				ExternalID:  maybe.Nothing[int32](),
				Name:        maybe.Nothing[string](),
				Date:        model.DateTime(time.Unix(0, 0)),
				PaymentType: maybe.Nothing[string](),
				Payment:     maybe.Nothing[model.Payment](),
				GameLink:    maybe.Nothing[string](),
			},
			{
				ID:          3,
				ExternalID:  maybe.Nothing[int32](),
				Name:        maybe.Nothing[string](),
				Date:        model.DateTime(time.Unix(0, 0)),
				PaymentType: maybe.Nothing[string](),
				Payment:     maybe.Nothing[model.Payment](),
				GameLink:    maybe.Nothing[string](),
			},
		}, 3, nil)

		got, err := fx.implementation.SearchGamesByLeagueID(fx.ctx, &gamepb.SearchGamesByLeagueIDRequest{
			Id:       model.LeagueQuizPlease,
			Page:     0,
			PageSize: 10,
		})
		assert.Equal(t, &gamepb.SearchGamesByLeagueIDResponse{
			Games: []*gamepb.Game{
				{
					Id:   1,
					Date: &timestamppb.Timestamp{},
				},
				{
					Id:   2,
					Date: &timestamppb.Timestamp{},
				},
				{
					Id:   3,
					Date: &timestamppb.Timestamp{},
				},
			},
			Total: 3,
		}, got)
		assert.NoError(t, err)
	})
}

func TestImplementation_SearchPassedAndRegisteredGames(t *testing.T) {
	t.Run("error: internal error", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().SearchPassedAndRegisteredGames(fx.ctx, uint64(0), uint64(3)).Return(nil, 0, errors.New("some error"))

		got, err := fx.implementation.SearchPassedAndRegisteredGames(fx.ctx, &gamepb.SearchPassedAndRegisteredGamesRequest{
			Page:     1,
			PageSize: 3,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("ok: empty list", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().SearchPassedAndRegisteredGames(fx.ctx, uint64(0), uint64(3)).Return([]model.Game{}, 0, nil)

		got, err := fx.implementation.SearchPassedAndRegisteredGames(fx.ctx, &gamepb.SearchPassedAndRegisteredGamesRequest{
			Page:     1,
			PageSize: 3,
		})
		assert.Equal(t, &gamepb.SearchPassedAndRegisteredGamesResponse{
			Games: []*gamepb.Game{},
			Total: 0,
		}, got)
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().SearchPassedAndRegisteredGames(fx.ctx, uint64(0), uint64(3)).Return([]model.Game{
			{
				ID:          1,
				ExternalID:  maybe.Nothing[int32](),
				Name:        maybe.Nothing[string](),
				Date:        model.DateTime(time.Unix(0, 0)),
				PaymentType: maybe.Nothing[string](),
				Payment:     maybe.Nothing[model.Payment](),
				GameLink:    maybe.Nothing[string](),
			},
			{
				ID:          2,
				ExternalID:  maybe.Nothing[int32](),
				Name:        maybe.Nothing[string](),
				Date:        model.DateTime(time.Unix(0, 0)),
				PaymentType: maybe.Nothing[string](),
				Payment:     maybe.Nothing[model.Payment](),
				GameLink:    maybe.Nothing[string](),
			},
			{
				ID:          3,
				ExternalID:  maybe.Nothing[int32](),
				Name:        maybe.Nothing[string](),
				Date:        model.DateTime(time.Unix(0, 0)),
				PaymentType: maybe.Nothing[string](),
				Payment:     maybe.Nothing[model.Payment](),
				GameLink:    maybe.Nothing[string](),
			},
		}, 4, nil)

		got, err := fx.implementation.SearchPassedAndRegisteredGames(fx.ctx, &gamepb.SearchPassedAndRegisteredGamesRequest{
			Page:     1,
			PageSize: 3,
		})
		assert.Equal(t, &gamepb.SearchPassedAndRegisteredGamesResponse{
			Games: []*gamepb.Game{
				{
					Id:   1,
					Date: &timestamppb.Timestamp{},
				},
				{
					Id:   2,
					Date: &timestamppb.Timestamp{},
				},
				{
					Id:   3,
					Date: &timestamppb.Timestamp{},
				},
			},
			Total: 4,
		}, got)
		assert.NoError(t, err)
	})
}
