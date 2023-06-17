package icsfilemanager

import (
	"errors"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	icsfilemanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/ics_file_manager"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestImplementation_GetICSFile(t *testing.T) {
	t.Run("error ICS file not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.icsFilesFacade.EXPECT().GetICSFile(fx.ctx, int32(1)).Return(model.ICSFile{}, model.ErrICSFileNotFound)

		got, err := fx.implementation.GetICSFile(fx.ctx, &icsfilemanagerpb.GetICSFileRequest{
			Id: 1,
		})

		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("some internal error", func(t *testing.T) {
		fx := tearUp(t)

		fx.icsFilesFacade.EXPECT().GetICSFile(fx.ctx, int32(1)).Return(model.ICSFile{}, errors.New("some error"))

		got, err := fx.implementation.GetICSFile(fx.ctx, &icsfilemanagerpb.GetICSFileRequest{
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

		fx.icsFilesFacade.EXPECT().GetICSFile(fx.ctx, int32(1)).Return(model.ICSFile{
			ID: 1,
		}, nil)

		got, err := fx.implementation.GetICSFile(fx.ctx, &icsfilemanagerpb.GetICSFileRequest{
			Id: 1,
		})

		assert.Equal(t, &icsfilemanagerpb.ICSFile{
			Id: 1,
		}, got)
		assert.NoError(t, err)
	})
}

func TestImplementation_GetICSFileByGameID(t *testing.T) {
	t.Run("error ICS file not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.icsFilesFacade.EXPECT().GetICSFileByGameID(fx.ctx, int32(1)).Return(model.ICSFile{}, model.ErrICSFileNotFound)

		got, err := fx.implementation.GetICSFileByGameID(fx.ctx, &icsfilemanagerpb.GetICSFileByGameIDRequest{
			GameId: 1,
		})

		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("some internal error", func(t *testing.T) {
		fx := tearUp(t)

		fx.icsFilesFacade.EXPECT().GetICSFileByGameID(fx.ctx, int32(1)).Return(model.ICSFile{}, errors.New("some error"))

		got, err := fx.implementation.GetICSFileByGameID(fx.ctx, &icsfilemanagerpb.GetICSFileByGameIDRequest{
			GameId: 1,
		})

		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.icsFilesFacade.EXPECT().GetICSFileByGameID(fx.ctx, int32(1)).Return(model.ICSFile{
			ID:     1,
			GameID: 1,
		}, nil)

		got, err := fx.implementation.GetICSFileByGameID(fx.ctx, &icsfilemanagerpb.GetICSFileByGameIDRequest{
			GameId: 1,
		})

		assert.Equal(t, &icsfilemanagerpb.ICSFile{
			Id:     1,
			GameId: 1,
		}, got)
		assert.NoError(t, err)
	})
}
