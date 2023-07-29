package photomanager

import (
	"errors"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	photomanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/photo_manager"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestImplementation_GetPhotosByGameID(t *testing.T) {
	t.Run("internal error while get photos by game id", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePhotosFacade.EXPECT().GetPhotosByGameID(fx.ctx, int32(1)).Return(nil, errors.New("some error"))

		got, err := fx.implementation.GetPhotosByGameID(fx.ctx, &photomanagerpb.GetPhotosByGameIDRequest{
			GameId: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("error \"game not found\" while get photos by game id", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePhotosFacade.EXPECT().GetPhotosByGameID(fx.ctx, int32(1)).Return(nil, games.ErrGameNotFound)

		got, err := fx.implementation.GetPhotosByGameID(fx.ctx, &photomanagerpb.GetPhotosByGameIDRequest{
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

		fx.gamePhotosFacade.EXPECT().GetPhotosByGameID(fx.ctx, int32(1)).Return([]string{"url1", "url2"}, nil)

		got, err := fx.implementation.GetPhotosByGameID(fx.ctx, &photomanagerpb.GetPhotosByGameIDRequest{
			GameId: 1,
		})
		assert.NotNil(t, got)
		assert.NoError(t, err)

		assert.Equal(t, &photomanagerpb.GetPhotosByGameIDResponse{
			Urls: []string{
				"url1",
				"url2",
			},
		}, got)
	})
}
