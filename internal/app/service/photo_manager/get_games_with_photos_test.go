package photomanager

import (
	"errors"
	"testing"
	"time"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	commonpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/common"
	photomanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/photo_manager"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestImplementation_GetGamesWithPhotos(t *testing.T) {
	t.Run("internal error while get games with photos", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePhotosFacade.EXPECT().GetNumberOfGamesWithPhotos(fx.ctx).Return(0, errors.New("some error"))

		got, err := fx.implementation.GetGamesWithPhotos(fx.ctx, &photomanagerpb.GetGamesWithPhotosRequest{
			Limit:  1,
			Offset: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("gamesTotal == 0", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePhotosFacade.EXPECT().GetNumberOfGamesWithPhotos(fx.ctx).Return(0, nil)

		got, err := fx.implementation.GetGamesWithPhotos(fx.ctx, &photomanagerpb.GetGamesWithPhotosRequest{
			Limit:  1,
			Offset: 1,
		})
		assert.NotNil(t, got)
		assert.NoError(t, err)
	})

	t.Run("internal error while get games with photos", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePhotosFacade.EXPECT().GetNumberOfGamesWithPhotos(fx.ctx).Return(10, nil)
		fx.gamePhotosFacade.EXPECT().GetGamesWithPhotos(fx.ctx, uint32(1), uint32(2)).Return(nil, errors.New("some error"))

		got, err := fx.implementation.GetGamesWithPhotos(fx.ctx, &photomanagerpb.GetGamesWithPhotosRequest{
			Limit:  1,
			Offset: 2,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePhotosFacade.EXPECT().GetNumberOfGamesWithPhotos(fx.ctx).Return(10, nil)
		fx.gamePhotosFacade.EXPECT().GetGamesWithPhotos(fx.ctx, uint32(1), uint32(2)).Return([]model.Game{
			{
				ID:     1,
				Number: "1",
			},
		}, nil)

		got, err := fx.implementation.GetGamesWithPhotos(fx.ctx, &photomanagerpb.GetGamesWithPhotosRequest{
			Limit:  1,
			Offset: 2,
		})
		assert.NotNil(t, got)
		assert.NoError(t, err)

		assert.Equal(t, &photomanagerpb.GetGamesWithPhotosResponse{
			Games: []*commonpb.Game{
				&commonpb.Game{
					Date:   timestamppb.New(time.Time{}),
					Id:     1,
					Number: "1",
				},
			},
			Total: 10,
		}, got)
	})
}
