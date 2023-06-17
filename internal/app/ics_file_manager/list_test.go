package icsfilemanager

import (
	"errors"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	icsfilemanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/ics_file_manager"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestImplementation_ListICSFiles(t *testing.T) {
	t.Run("error while list certificates", func(t *testing.T) {
		fx := tearUp(t)

		fx.icsFilesFacade.EXPECT().ListICSFiles(fx.ctx).Return(nil, errors.New("some error"))

		got, err := fx.implementation.ListICSFiles(fx.ctx, &emptypb.Empty{})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.icsFilesFacade.EXPECT().ListICSFiles(fx.ctx).Return([]model.ICSFile{
			{
				ID:     1,
				GameID: 1,
				Name:   "some uuid 1",
			},
			{
				ID:     2,
				GameID: 2,
				Name:   "some uuid 2",
			},
		}, nil)

		got, err := fx.implementation.ListICSFiles(fx.ctx, &emptypb.Empty{})
		assert.ElementsMatch(t,
			[]*icsfilemanagerpb.ICSFile{
				{
					Id:     1,
					GameId: 1,
					Name:   "some uuid 1",
				},
				{
					Id:     2,
					GameId: 2,
					Name:   "some uuid 2",
				},
			},
			got.GetIcsFiles())
		assert.NoError(t, err)
	})
}
