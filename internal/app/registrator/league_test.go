package registrator

import (
	"errors"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestRegistrator_GetLeagueByID(t *testing.T) {
	t.Run("internal error while get league by ID", func(t *testing.T) {
		fx := tearUp(t)

		fx.leaguesFacade.EXPECT().GetLeagueByID(fx.ctx, int32(1)).Return(model.League{}, errors.New("some error"))

		got, err := fx.registrator.GetLeagueByID(fx.ctx, &registrator.GetLeagueByIDRequest{
			Id: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("league not found error while get league by ID", func(t *testing.T) {
		fx := tearUp(t)

		fx.leaguesFacade.EXPECT().GetLeagueByID(fx.ctx, int32(1)).Return(model.League{}, model.ErrLeagueNotFound)

		got, err := fx.registrator.GetLeagueByID(fx.ctx, &registrator.GetLeagueByIDRequest{
			Id: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.leaguesFacade.EXPECT().GetLeagueByID(fx.ctx, int32(1)).Return(model.League{
			ID:   1,
			Name: "name",
		}, nil)

		got, err := fx.registrator.GetLeagueByID(fx.ctx, &registrator.GetLeagueByIDRequest{
			Id: 1,
		})
		assert.NotNil(t, got)
		assert.Equal(t, &registrator.GetLeagueByIDResponse{
			League: &registrator.League{
				Id:   1,
				Name: "name",
			},
		}, got)
		assert.NoError(t, err)
	})
}
