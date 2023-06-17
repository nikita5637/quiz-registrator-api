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

func TestImplementation_CreateICSFile(t *testing.T) {
	t.Run("error while create ICS file. game not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.icsFilesFacade.EXPECT().CreateICSFile(fx.ctx, model.ICSFile{
			GameID: 1,
			Name:   "some uuid",
		}).Return(model.ICSFile{}, model.ErrGameNotFound)

		got, err := fx.implementation.CreateICSFile(fx.ctx, &icsfilemanagerpb.CreateICSFileRequest{
			IcsFile: &icsfilemanagerpb.ICSFile{
				GameId: 1,
				Name:   "some uuid",
			},
		})

		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("error while create ICS file. ICS file for game already exists", func(t *testing.T) {
		fx := tearUp(t)

		fx.icsFilesFacade.EXPECT().CreateICSFile(fx.ctx, model.ICSFile{
			GameID: 1,
			Name:   "some uuid",
		}).Return(model.ICSFile{}, model.ErrICSFileForGameAlreadyExists)

		got, err := fx.implementation.CreateICSFile(fx.ctx, &icsfilemanagerpb.CreateICSFileRequest{
			IcsFile: &icsfilemanagerpb.ICSFile{
				GameId: 1,
				Name:   "some uuid",
			},
		})

		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.AlreadyExists, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("error while create ICS file. ICS file name already exists", func(t *testing.T) {
		fx := tearUp(t)

		fx.icsFilesFacade.EXPECT().CreateICSFile(fx.ctx, model.ICSFile{
			GameID: 1,
			Name:   "some uuid",
		}).Return(model.ICSFile{}, model.ErrICSFileNameAlreadyExists)

		got, err := fx.implementation.CreateICSFile(fx.ctx, &icsfilemanagerpb.CreateICSFileRequest{
			IcsFile: &icsfilemanagerpb.ICSFile{
				GameId: 1,
				Name:   "some uuid",
			},
		})

		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.AlreadyExists, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("error while create ICS file. other error", func(t *testing.T) {
		fx := tearUp(t)

		fx.icsFilesFacade.EXPECT().CreateICSFile(fx.ctx, model.ICSFile{
			GameID: 1,
			Name:   "some uuid",
		}).Return(model.ICSFile{}, errors.New("some error"))

		got, err := fx.implementation.CreateICSFile(fx.ctx, &icsfilemanagerpb.CreateICSFileRequest{
			IcsFile: &icsfilemanagerpb.ICSFile{
				GameId: 1,
				Name:   "some uuid",
			},
		})

		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.icsFilesFacade.EXPECT().CreateICSFile(fx.ctx, model.ICSFile{
			GameID: 1,
			Name:   "some uuid",
		}).Return(model.ICSFile{
			ID:     1,
			GameID: 1,
			Name:   "some uuid",
		}, nil)

		got, err := fx.implementation.CreateICSFile(fx.ctx, &icsfilemanagerpb.CreateICSFileRequest{
			IcsFile: &icsfilemanagerpb.ICSFile{
				GameId: 1,
				Name:   "some uuid",
			},
		})

		assert.Equal(t, &icsfilemanagerpb.ICSFile{
			Id:     1,
			GameId: 1,
			Name:   "some uuid",
		}, got)
		assert.NoError(t, err)
	})
}
