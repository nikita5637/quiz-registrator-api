package registrator

import (
	"errors"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"
	users_utils "github.com/nikita5637/quiz-registrator-api/utils/users"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestRegistrator_RegisterPlayer(t *testing.T) {
	t.Run("unauthenticated request", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.registrator.RegisterPlayer(fx.ctx, &registrator.RegisterPlayerRequest{
			GameId:     1,
			PlayerType: registrator.PlayerType_PLAYER_TYPE_MAIN,
			Degree:     registrator.Degree_DEGREE_LIKELY,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Unauthenticated, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("internal error while register player", func(t *testing.T) {
		fx := tearUp(t)
		ctx := users_utils.NewContextWithUser(fx.ctx, model.User{
			ID: 1,
		})

		fx.gamesFacade.EXPECT().RegisterPlayer(ctx, int32(1), int32(1), int32(1), int32(registrator.Degree_DEGREE_LIKELY)).Return(model.RegisterPlayerStatusInvalid, errors.New("some error"))

		got, err := fx.registrator.RegisterPlayer(ctx, &registrator.RegisterPlayerRequest{
			GameId:     1,
			PlayerType: registrator.PlayerType_PLAYER_TYPE_MAIN,
			Degree:     registrator.Degree_DEGREE_LIKELY,
		})
		assert.Nil(t, got)
		assert.Error(t, err)
	})

	t.Run("error game not found while register player", func(t *testing.T) {
		fx := tearUp(t)
		ctx := users_utils.NewContextWithUser(fx.ctx, model.User{
			ID: 1,
		})

		fx.gamesFacade.EXPECT().RegisterPlayer(ctx, int32(1), int32(1), int32(1), int32(registrator.Degree_DEGREE_LIKELY)).Return(model.RegisterPlayerStatusInvalid, model.ErrGameNotFound)

		got, err := fx.registrator.RegisterPlayer(ctx, &registrator.RegisterPlayerRequest{
			GameId:     1,
			PlayerType: registrator.PlayerType_PLAYER_TYPE_MAIN,
			Degree:     registrator.Degree_DEGREE_LIKELY,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("error no free slots while register player", func(t *testing.T) {
		fx := tearUp(t)
		ctx := users_utils.NewContextWithUser(fx.ctx, model.User{
			ID: 1,
		})

		fx.gamesFacade.EXPECT().RegisterPlayer(ctx, int32(1), int32(1), int32(1), int32(registrator.Degree_DEGREE_LIKELY)).Return(model.RegisterPlayerStatusInvalid, model.ErrGameNoFreeSlots)

		got, err := fx.registrator.RegisterPlayer(ctx, &registrator.RegisterPlayerRequest{
			GameId:     1,
			PlayerType: registrator.PlayerType_PLAYER_TYPE_MAIN,
			Degree:     registrator.Degree_DEGREE_LIKELY,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.AlreadyExists, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("ok main player", func(t *testing.T) {
		fx := tearUp(t)
		ctx := users_utils.NewContextWithUser(fx.ctx, model.User{
			ID: 1,
		})

		fx.gamesFacade.EXPECT().RegisterPlayer(ctx, int32(1), int32(1), int32(1), int32(registrator.Degree_DEGREE_LIKELY)).Return(model.RegisterPlayerStatusOK, nil)

		got, err := fx.registrator.RegisterPlayer(ctx, &registrator.RegisterPlayerRequest{
			GameId:     1,
			PlayerType: registrator.PlayerType_PLAYER_TYPE_MAIN,
			Degree:     registrator.Degree_DEGREE_LIKELY,
		})
		assert.NotNil(t, got)
		assert.Equal(t, &registrator.RegisterPlayerResponse{
			Status: registrator.RegisterPlayerStatus_REGISTER_PLAYER_STATUS_OK,
		}, got)
		assert.NoError(t, err)
	})

	t.Run("ok legioner", func(t *testing.T) {
		fx := tearUp(t)
		ctx := users_utils.NewContextWithUser(fx.ctx, model.User{
			ID: 1,
		})

		fx.gamesFacade.EXPECT().RegisterPlayer(ctx, int32(1), int32(0), int32(1), int32(registrator.Degree_DEGREE_LIKELY)).Return(model.RegisterPlayerStatusOK, nil)

		got, err := fx.registrator.RegisterPlayer(ctx, &registrator.RegisterPlayerRequest{
			GameId:     1,
			PlayerType: registrator.PlayerType_PLAYER_TYPE_LEGIONER,
			Degree:     registrator.Degree_DEGREE_LIKELY,
		})
		assert.NotNil(t, got)
		assert.Equal(t, &registrator.RegisterPlayerResponse{
			Status: registrator.RegisterPlayerStatus_REGISTER_PLAYER_STATUS_OK,
		}, got)
		assert.NoError(t, err)
	})
}
