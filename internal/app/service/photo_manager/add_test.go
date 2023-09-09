package photomanager

import (
	"errors"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	photomanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/photo_manager"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestRegistrator_AddGamePhotos(t *testing.T) {
	t.Run("internal error while add game photos", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePhotosFacade.EXPECT().AddGamePhotos(fx.ctx, int32(1), []string{"url1", "url2"}).Return(errors.New("some error"))

		got, err := fx.implementation.AddGamePhotos(fx.ctx, &photomanagerpb.AddGamePhotosRequest{
			GameId: 1,
			Urls: []string{
				"url1",
				"url2",
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("error \"game not found\" while add game photos", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePhotosFacade.EXPECT().AddGamePhotos(fx.ctx, int32(1), []string{"url1", "url2"}).Return(games.ErrGameNotFound)

		got, err := fx.implementation.AddGamePhotos(fx.ctx, &photomanagerpb.AddGamePhotosRequest{
			GameId: 1,
			Urls: []string{
				"url1",
				"url2",
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.FailedPrecondition, st.Code())
		assert.Len(t, st.Details(), 2)

		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, games.ReasonGameNotFound, errorInfo.Reason)
		assert.Nil(t, errorInfo.Metadata)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePhotosFacade.EXPECT().AddGamePhotos(fx.ctx, int32(1), []string{"url1", "url2"}).Return(nil)

		got, err := fx.implementation.AddGamePhotos(fx.ctx, &photomanagerpb.AddGamePhotosRequest{
			GameId: 1,
			Urls: []string{
				"url1",
				"url2",
			},
		})
		assert.NotNil(t, got)
		assert.NoError(t, err)
	})
}