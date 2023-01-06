package registrator

import (
	"errors"
	"testing"
	"time"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestRegistrator_AddGamePhotos(t *testing.T) {
	t.Run("internal error while add game photos", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePhotosFacade.EXPECT().AddGamePhotos(fx.ctx, int32(1), []string{"url1", "url2"}).Return(errors.New("some error"))

		got, err := fx.registrator.AddGamePhotos(fx.ctx, &registrator.AddGamePhotosRequest{
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

		fx.gamePhotosFacade.EXPECT().AddGamePhotos(fx.ctx, int32(1), []string{"url1", "url2"}).Return(model.ErrGameNotFound)

		got, err := fx.registrator.AddGamePhotos(fx.ctx, &registrator.AddGamePhotosRequest{
			GameId: 1,
			Urls: []string{
				"url1",
				"url2",
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePhotosFacade.EXPECT().AddGamePhotos(fx.ctx, int32(1), []string{"url1", "url2"}).Return(nil)

		got, err := fx.registrator.AddGamePhotos(fx.ctx, &registrator.AddGamePhotosRequest{
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

func TestRegistrator_GetGamesWithPhotos(t *testing.T) {
	t.Run("internal error while get games with photos", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePhotosFacade.EXPECT().GetNumberOfGamesWithPhotos(fx.ctx).Return(0, errors.New("some error"))

		got, err := fx.registrator.GetGamesWithPhotos(fx.ctx, &registrator.GetGamesWithPhotosRequest{
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

		got, err := fx.registrator.GetGamesWithPhotos(fx.ctx, &registrator.GetGamesWithPhotosRequest{
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

		got, err := fx.registrator.GetGamesWithPhotos(fx.ctx, &registrator.GetGamesWithPhotosRequest{
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

		got, err := fx.registrator.GetGamesWithPhotos(fx.ctx, &registrator.GetGamesWithPhotosRequest{
			Limit:  1,
			Offset: 2,
		})
		assert.NotNil(t, got)
		assert.NoError(t, err)

		assert.Equal(t, &registrator.GetGamesWithPhotosResponse{
			Games: []*registrator.Game{
				&registrator.Game{
					Date:   timestamppb.New(time.Time{}),
					Id:     1,
					Number: "1",
				},
			},
			Total: 10,
		}, got)
	})
}

func TestRegistrator_GetPhotosByGameID(t *testing.T) {
	t.Run("internal error while get photos by game id", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePhotosFacade.EXPECT().GetPhotosByGameID(fx.ctx, int32(1)).Return(nil, errors.New("some error"))

		got, err := fx.registrator.GetPhotosByGameID(fx.ctx, &registrator.GetPhotosByGameIDRequest{
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

		fx.gamePhotosFacade.EXPECT().GetPhotosByGameID(fx.ctx, int32(1)).Return(nil, model.ErrGameNotFound)

		got, err := fx.registrator.GetPhotosByGameID(fx.ctx, &registrator.GetPhotosByGameIDRequest{
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

		got, err := fx.registrator.GetPhotosByGameID(fx.ctx, &registrator.GetPhotosByGameIDRequest{
			GameId: 1,
		})
		assert.NotNil(t, got)
		assert.NoError(t, err)

		assert.Equal(t, &registrator.GetPhotosByGameIDResponse{
			Urls: []string{
				"url1",
				"url2",
			},
		}, got)
	})
}
