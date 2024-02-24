package gameplayer

import (
	"errors"
	"math"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/gameplayers"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	gameplayer "github.com/nikita5637/quiz-registrator-api/pkg/pb/game_player"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestImplementation_UpdatePlayerDegree(t *testing.T) {
	t.Run("error: game player not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePlayersFacade.EXPECT().GetGamePlayer(fx.ctx, int32(1)).Return(model.GamePlayer{}, gameplayers.ErrGamePlayerNotFound)

		got, err := fx.implementation.UpdatePlayerDegree(fx.ctx, &gameplayer.UpdatePlayerDegreeRequest{
			Id: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())
		assert.Len(t, st.Details(), 2)

		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, "GAME_PLAYER_NOT_FOUND", errorInfo.Reason)
		assert.Equal(t, map[string]string{
			"error": "game player not found",
		}, errorInfo.Metadata)
	})

	t.Run("error: getting game player internal error", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePlayersFacade.EXPECT().GetGamePlayer(fx.ctx, int32(1)).Return(model.GamePlayer{}, errors.New("some error"))

		got, err := fx.implementation.UpdatePlayerDegree(fx.ctx, &gameplayer.UpdatePlayerDegreeRequest{
			Id: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("error: validation error", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePlayersFacade.EXPECT().GetGamePlayer(fx.ctx, int32(1)).Return(model.GamePlayer{
			ID: 1,
		}, nil)

		got, err := fx.implementation.UpdatePlayerDegree(fx.ctx, &gameplayer.UpdatePlayerDegreeRequest{
			Id:     1,
			Degree: gameplayer.Degree(math.MaxInt32),
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)

		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, "INVALID_DEGREE", errorInfo.Reason)
		assert.Nil(t, errorInfo.Metadata)
	})

	t.Run("error: game not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePlayersFacade.EXPECT().GetGamePlayer(fx.ctx, int32(1)).Return(model.GamePlayer{
			ID:     1,
			GameID: 1,
		}, nil)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.NewGame(), games.ErrGameNotFound)

		got, err := fx.implementation.UpdatePlayerDegree(fx.ctx, &gameplayer.UpdatePlayerDegreeRequest{
			Id:     1,
			Degree: gameplayer.Degree_DEGREE_LIKELY,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.FailedPrecondition, st.Code())
		assert.Len(t, st.Details(), 2)

		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, "GAME_NOT_FOUND", errorInfo.Reason)
		assert.Equal(t, map[string]string{
			"error": "game not found",
		}, errorInfo.Metadata)
	})

	t.Run("error: getting game internal error", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePlayersFacade.EXPECT().GetGamePlayer(fx.ctx, int32(1)).Return(model.GamePlayer{
			ID:     1,
			GameID: 1,
		}, nil)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.NewGame(), errors.New("some error"))

		got, err := fx.implementation.UpdatePlayerDegree(fx.ctx, &gameplayer.UpdatePlayerDegreeRequest{
			Id:     1,
			Degree: gameplayer.Degree_DEGREE_LIKELY,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("error: game has passed", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePlayersFacade.EXPECT().GetGamePlayer(fx.ctx, int32(1)).Return(model.GamePlayer{
			ID:     1,
			GameID: 1,
		}, nil)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			ID:        1,
			HasPassed: true,
		}, nil)

		got, err := fx.implementation.UpdatePlayerDegree(fx.ctx, &gameplayer.UpdatePlayerDegreeRequest{
			Id:     1,
			Degree: gameplayer.Degree_DEGREE_LIKELY,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.FailedPrecondition, st.Code())
		assert.Len(t, st.Details(), 2)

		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, "GAME_HAS_PASSED", errorInfo.Reason)
		assert.Nil(t, errorInfo.Metadata)
	})

	t.Run("error: there are no registered for the game", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePlayersFacade.EXPECT().GetGamePlayer(fx.ctx, int32(1)).Return(model.GamePlayer{
			ID:     1,
			GameID: 1,
		}, nil)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			ID:         1,
			HasPassed:  false,
			Registered: false,
		}, nil)

		got, err := fx.implementation.UpdatePlayerDegree(fx.ctx, &gameplayer.UpdatePlayerDegreeRequest{
			Id:     1,
			Degree: gameplayer.Degree_DEGREE_LIKELY,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.FailedPrecondition, st.Code())
		assert.Len(t, st.Details(), 2)

		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, "THERE_ARE_NO_REGISTRATION_FOR_THE_GAME", errorInfo.Reason)
		assert.Nil(t, errorInfo.Metadata)
	})

	t.Run("error: patching game player internal error", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePlayersFacade.EXPECT().GetGamePlayer(fx.ctx, int32(1)).Return(model.GamePlayer{
			ID:     1,
			GameID: 1,
		}, nil)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			ID:         1,
			HasPassed:  false,
			Registered: true,
		}, nil)

		fx.gamePlayersFacade.EXPECT().PatchGamePlayer(fx.ctx, model.GamePlayer{
			ID:     1,
			GameID: 1,
			Degree: 1,
		}).Return(model.GamePlayer{}, errors.New("some error"))

		got, err := fx.implementation.UpdatePlayerDegree(fx.ctx, &gameplayer.UpdatePlayerDegreeRequest{
			Id:     1,
			Degree: gameplayer.Degree_DEGREE_LIKELY,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePlayersFacade.EXPECT().GetGamePlayer(fx.ctx, int32(1)).Return(model.GamePlayer{
			ID:     1,
			GameID: 1,
		}, nil)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			ID:         1,
			HasPassed:  false,
			Registered: true,
		}, nil)

		fx.gamePlayersFacade.EXPECT().PatchGamePlayer(fx.ctx, model.GamePlayer{
			ID:     1,
			GameID: 1,
			Degree: 1,
		}).Return(model.GamePlayer{
			ID:     1,
			GameID: 1,
			Degree: 1,
		}, nil)

		got, err := fx.implementation.UpdatePlayerDegree(fx.ctx, &gameplayer.UpdatePlayerDegreeRequest{
			Id:     1,
			Degree: gameplayer.Degree_DEGREE_LIKELY,
		})
		assert.NotNil(t, got)
		assert.Equal(t, &emptypb.Empty{}, got)
		assert.NoError(t, err)
	})
}
