package croupier

import (
	"errors"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	croupierpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/croupier"
	usersutils "github.com/nikita5637/quiz-registrator-api/utils/users"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestRegistrator_RegistRegisteredByerForLottery(t *testing.T) {
	/*
		t.Run("validate register for lottery request error. gameID lt 1", func(t *testing.T) {
			fx := tearUp(t)

			got, err := fx.implementation.RegisterForLottery(fx.ctx, &croupierpb.RegisterForLotteryRequest{
				GameId: 0,
			})
			assert.Nil(t, got)
			assert.Error(t, err)

			st := status.Convert(err)
			assert.Equal(t, codes.InvalidArgument, st.Code())
			assert.Len(t, st.Details(), 0)
		})

		t.Run("error. empty user.Email", func(t *testing.T) {
			fx := tearUp(t)

			ctx := usersutils.NewContextWithUser(fx.ctx, model.User{
				Name:  "user name",
				Phone: "user phone",
			})

			got, err := fx.implementation.RegisterForLottery(ctx, &croupierpb.RegisterForLotteryRequest{
				GameId: 1,
			})
			assert.Nil(t, got)
			assert.Error(t, err)

			st := status.Convert(err)
			assert.Equal(t, codes.PermissionDenied, st.Code())
			assert.Len(t, st.Details(), 2)
		})

		t.Run("error. empty user.Name", func(t *testing.T) {
			fx := tearUp(t)

			ctx := usersutils.NewContextWithUser(fx.ctx, model.User{
				Email: "user email",
				Phone: "user phone",
			})

			got, err := fx.implementation.RegisterForLottery(ctx, &croupierpb.RegisterForLotteryRequest{
				GameId: 1,
			})
			assert.Nil(t, got)
			assert.Error(t, err)

			st := status.Convert(err)
			assert.Equal(t, codes.PermissionDenied, st.Code())
			assert.Len(t, st.Details(), 2)
		})

		t.Run("error. empty user.Phone", func(t *testing.T) {
			fx := tearUp(t)

			ctx := usersutils.NewContextWithUser(fx.ctx, model.User{
				Email: "user email",
				Name:  "user name",
			})

			got, err := fx.implementation.RegisterForLottery(ctx, &croupierpb.RegisterForLotteryRequest{
				GameId: 1,
			})
			assert.Nil(t, got)
			assert.Error(t, err)

			st := status.Convert(err)
			assert.Equal(t, codes.PermissionDenied, st.Code())
			assert.Len(t, st.Details(), 2)
		})

		t.Run("error. gamePlayersFacade error", func(t *testing.T) {
			fx := tearUp(t)

			ctx := usersutils.NewContextWithUser(fx.ctx, model.User{
				ID:    777,
				Email: "user email",
				Name:  "user name",
				Phone: "user phone",
			})

			fx.gamePlayersFacade.EXPECT().PlayerRegisteredOnGame(ctx, int32(1), int32(777)).Return(false, errors.New("some error"))

			got, err := fx.implementation.RegisterForLottery(ctx, &croupierpb.RegisterForLotteryRequest{
				GameId: 1,
			})
			assert.Nil(t, got)
			assert.Error(t, err)

			st := status.Convert(err)
			assert.Equal(t, codes.Internal, st.Code())
			assert.Len(t, st.Details(), 0)
		})

		t.Run("error. user not registered for a game", func(t *testing.T) {
			fx := tearUp(t)

			ctx := usersutils.NewContextWithUser(fx.ctx, model.User{
				ID:    777,
				Email: "user email",
				Name:  "user name",
				Phone: "user phone",
			})

			fx.gamePlayersFacade.EXPECT().PlayerRegisteredOnGame(ctx, int32(1), int32(777)).Return(false, nil)

			got, err := fx.implementation.RegisterForLottery(ctx, &croupierpb.RegisterForLotteryRequest{
				GameId: 1,
			})
			assert.Nil(t, got)
			assert.Error(t, err)

			st := status.Convert(err)
			assert.Equal(t, codes.PermissionDenied, st.Code())
			assert.Len(t, st.Details(), 2)
		})
	*/

	t.Run("error while get game", func(t *testing.T) {
		fx := tearUp(t)

		ctx := usersutils.NewContextWithUser(fx.ctx, model.User{
			ID:    777,
			Email: "user email",
			Name:  "user name",
			Phone: "user phone",
		})

		fx.gamePlayersFacade.EXPECT().PlayerRegisteredOnGame(ctx, int32(1), int32(777)).Return(true, nil)

		fx.gamesFacade.EXPECT().GetGameByID(ctx, int32(1)).Return(model.Game{}, errors.New("some error"))

		got, err := fx.implementation.RegisterForLottery(ctx, &croupierpb.RegisterForLotteryRequest{
			GameId: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("game not found", func(t *testing.T) {
		fx := tearUp(t)

		ctx := usersutils.NewContextWithUser(fx.ctx, model.User{
			ID:    777,
			Email: "user email",
			Name:  "user name",
			Phone: "user phone",
		})

		fx.gamePlayersFacade.EXPECT().PlayerRegisteredOnGame(ctx, int32(1), int32(777)).Return(true, nil)

		fx.gamesFacade.EXPECT().GetGameByID(ctx, int32(1)).Return(model.Game{}, model.ErrGameNotFound)

		got, err := fx.implementation.RegisterForLottery(ctx, &croupierpb.RegisterForLotteryRequest{
			GameId: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Len(t, st.Details(), 2)
	})

	t.Run("error lottery not available", func(t *testing.T) {
		fx := tearUp(t)

		ctx := usersutils.NewContextWithUser(fx.ctx, model.User{
			ID:    777,
			Email: "user email",
			Name:  "user name",
			Phone: "user phone",
		})

		fx.gamePlayersFacade.EXPECT().PlayerRegisteredOnGame(ctx, int32(1), int32(777)).Return(true, nil)

		fx.gamesFacade.EXPECT().GetGameByID(ctx, int32(1)).Return(model.Game{
			ID: 1,
		}, nil)

		fx.croupier.EXPECT().RegisterForLottery(ctx, model.Game{
			ID: 1,
		}, model.User{
			ID:    777,
			Email: "user email",
			Name:  "user name",
			Phone: "user phone",
		}).Return(0, model.ErrLotteryNotAvailable)

		got, err := fx.implementation.RegisterForLottery(ctx, &croupierpb.RegisterForLotteryRequest{
			GameId: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("error lottery not implemented", func(t *testing.T) {
		fx := tearUp(t)

		ctx := usersutils.NewContextWithUser(fx.ctx, model.User{
			ID:    777,
			Email: "user email",
			Name:  "user name",
			Phone: "user phone",
		})

		fx.gamePlayersFacade.EXPECT().PlayerRegisteredOnGame(ctx, int32(1), int32(777)).Return(true, nil)

		fx.gamesFacade.EXPECT().GetGameByID(ctx, int32(1)).Return(model.Game{
			ID: 1,
		}, nil)

		fx.croupier.EXPECT().RegisterForLottery(ctx, model.Game{
			ID: 1,
		}, model.User{
			ID:    777,
			Email: "user email",
			Name:  "user name",
			Phone: "user phone",
		}).Return(0, model.ErrLotteryNotImplemented)

		got, err := fx.implementation.RegisterForLottery(ctx, &croupierpb.RegisterForLotteryRequest{
			GameId: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Unimplemented, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("other error while registration", func(t *testing.T) {
		fx := tearUp(t)

		ctx := usersutils.NewContextWithUser(fx.ctx, model.User{
			ID:    777,
			Email: "user email",
			Name:  "user name",
			Phone: "user phone",
		})

		fx.gamePlayersFacade.EXPECT().PlayerRegisteredOnGame(ctx, int32(1), int32(777)).Return(true, nil)

		fx.gamesFacade.EXPECT().GetGameByID(ctx, int32(1)).Return(model.Game{
			ID: 1,
		}, nil)

		fx.croupier.EXPECT().RegisterForLottery(ctx, model.Game{
			ID: 1,
		}, model.User{
			ID:    777,
			Email: "user email",
			Name:  "user name",
			Phone: "user phone",
		}).Return(0, errors.New("some error"))

		got, err := fx.implementation.RegisterForLottery(ctx, &croupierpb.RegisterForLotteryRequest{
			GameId: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("ok with number", func(t *testing.T) {
		fx := tearUp(t)

		ctx := usersutils.NewContextWithUser(fx.ctx, model.User{
			ID:    777,
			Email: "user email",
			Name:  "user name",
			Phone: "user phone",
		})

		fx.gamePlayersFacade.EXPECT().PlayerRegisteredOnGame(ctx, int32(1), int32(777)).Return(true, nil)

		fx.gamesFacade.EXPECT().GetGameByID(ctx, int32(1)).Return(model.Game{
			ID: 1,
		}, nil)

		fx.croupier.EXPECT().RegisterForLottery(ctx, model.Game{
			ID: 1,
		}, model.User{
			ID:    777,
			Email: "user email",
			Name:  "user name",
			Phone: "user phone",
		}).Return(100, nil)

		got, err := fx.implementation.RegisterForLottery(ctx, &croupierpb.RegisterForLotteryRequest{
			GameId: 1,
		})
		assert.Equal(t, &croupierpb.RegisterForLotteryResponse{
			Number: 100,
		}, got)
		assert.NoError(t, err)
	})

	t.Run("ok without number", func(t *testing.T) {
		fx := tearUp(t)

		ctx := usersutils.NewContextWithUser(fx.ctx, model.User{
			ID:    777,
			Email: "user email",
			Name:  "user name",
			Phone: "user phone",
		})

		fx.gamePlayersFacade.EXPECT().PlayerRegisteredOnGame(ctx, int32(1), int32(777)).Return(true, nil)

		fx.gamesFacade.EXPECT().GetGameByID(ctx, int32(1)).Return(model.Game{
			ID: 1,
		}, nil)

		fx.croupier.EXPECT().RegisterForLottery(ctx, model.Game{
			ID: 1,
		}, model.User{
			ID:    777,
			Email: "user email",
			Name:  "user name",
			Phone: "user phone",
		}).Return(0, nil)

		got, err := fx.implementation.RegisterForLottery(ctx, &croupierpb.RegisterForLotteryRequest{
			GameId: 1,
		})
		assert.Equal(t, &croupierpb.RegisterForLotteryResponse{
			Number: 0,
		}, got)
		assert.NoError(t, err)
	})
}
