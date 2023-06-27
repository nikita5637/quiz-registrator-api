package place

import (
	"errors"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/places"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	placepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/place"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestImplementation_GetPlace(t *testing.T) {
	t.Run("internal error while get place by ID", func(t *testing.T) {
		fx := tearUp(t)

		fx.placesFacade.EXPECT().GetPlaceByID(fx.ctx, int32(1)).Return(model.Place{}, errors.New("some error"))

		got, err := fx.implementation.GetPlace(fx.ctx, &placepb.GetPlaceRequest{
			Id: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("place not found error while get place by ID", func(t *testing.T) {
		fx := tearUp(t)

		fx.placesFacade.EXPECT().GetPlaceByID(fx.ctx, int32(1)).Return(model.Place{}, places.ErrPlaceNotFound)

		got, err := fx.implementation.GetPlace(fx.ctx, &placepb.GetPlaceRequest{
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

		fx.placesFacade.EXPECT().GetPlaceByID(fx.ctx, int32(1)).Return(model.Place{
			ID:        1,
			Name:      "name",
			Address:   "address",
			ShortName: "short name",
			Latitude:  1.1,
			Longitude: 1.1,
			MenuLink:  "menu link",
		}, nil)

		got, err := fx.implementation.GetPlace(fx.ctx, &placepb.GetPlaceRequest{
			Id: 1,
		})
		assert.NotNil(t, got)
		assert.Equal(t, &placepb.Place{
			Id:        1,
			Name:      "name",
			Address:   "address",
			ShortName: "short name",
			Latitude:  1.1,
			Longitude: 1.1,
			MenuLink:  "menu link",
		}, got)
		assert.NoError(t, err)
	})
}
