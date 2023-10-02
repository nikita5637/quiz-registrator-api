package gameplayer

import (
	"errors"
	"testing"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/gameplayers"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/users"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	gameplayer "github.com/nikita5637/quiz-registrator-api/pkg/pb/game_player"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestRegistrator_RegisterPlayer(t *testing.T) {
	t.Run("game player is nil", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.implementation.RegisterPlayer(fx.ctx, &gameplayer.RegisterPlayerRequest{})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("validation error", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.implementation.RegisterPlayer(fx.ctx, &gameplayer.RegisterPlayerRequest{
			GamePlayer: &gameplayer.GamePlayer{
				RegisteredBy: 1,
				Degree:       gameplayer.Degree_DEGREE_LIKELY,
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)

		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, invalidGameIDReason, errorInfo.Reason)
		assert.Equal(t, map[string]string{
			"error": "GameID: cannot be blank.",
		}, errorInfo.Metadata)
	})

	t.Run("game not found error while get game", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{}, games.ErrGameNotFound)

		got, err := fx.implementation.RegisterPlayer(fx.ctx, &gameplayer.RegisterPlayerRequest{
			GamePlayer: &gameplayer.GamePlayer{
				GameId: 1,
				UserId: &wrapperspb.Int32Value{
					Value: 1,
				},
				RegisteredBy: 1,
				Degree:       gameplayer.Degree_DEGREE_LIKELY,
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

	t.Run("game has passed", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			HasPassed: true,
		}, nil)

		got, err := fx.implementation.RegisterPlayer(fx.ctx, &gameplayer.RegisterPlayerRequest{
			GamePlayer: &gameplayer.GamePlayer{
				GameId: 1,
				UserId: &wrapperspb.Int32Value{
					Value: 1,
				},
				RegisteredBy: 1,
				Degree:       gameplayer.Degree_DEGREE_LIKELY,
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.FailedPrecondition, st.Code())
		assert.Len(t, st.Details(), 2)

		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, games.ReasonGameHasPassed, errorInfo.Reason)
		assert.Nil(t, errorInfo.Metadata)
	})

	t.Run("internal error while get game", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{}, errors.New("some error"))

		got, err := fx.implementation.RegisterPlayer(fx.ctx, &gameplayer.RegisterPlayerRequest{
			GamePlayer: &gameplayer.GamePlayer{
				GameId: 1,
				UserId: &wrapperspb.Int32Value{
					Value: 1,
				},
				RegisteredBy: 1,
				Degree:       gameplayer.Degree_DEGREE_LIKELY,
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("there are no registration for the game", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			Registered: false,
		}, nil)

		got, err := fx.implementation.RegisterPlayer(fx.ctx, &gameplayer.RegisterPlayerRequest{
			GamePlayer: &gameplayer.GamePlayer{
				GameId: 1,
				UserId: &wrapperspb.Int32Value{
					Value: 1,
				},
				RegisteredBy: 1,
				Degree:       gameplayer.Degree_DEGREE_LIKELY,
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.FailedPrecondition, st.Code())
		assert.Len(t, st.Details(), 2)

		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, reasonThereAreNoRegistrationForTheGame, errorInfo.Reason)
		assert.Nil(t, errorInfo.Metadata)
	})

	t.Run("get game players by game ID internal error", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			Registered: true,
		}, nil)

		fx.gamePlayersFacade.EXPECT().GetGamePlayersByGameID(fx.ctx, int32(1)).Return(nil, errors.New("some error"))

		got, err := fx.implementation.RegisterPlayer(fx.ctx, &gameplayer.RegisterPlayerRequest{
			GamePlayer: &gameplayer.GamePlayer{
				GameId:       1,
				RegisteredBy: 1,
				Degree:       gameplayer.Degree_DEGREE_LIKELY,
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("there are no free slot", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			ID:         1,
			MaxPlayers: 2,
			Registered: true,
		}, nil)

		fx.gamePlayersFacade.EXPECT().GetGamePlayersByGameID(fx.ctx, int32(1)).Return([]model.GamePlayer{
			{
				ID:           1,
				GameID:       1,
				UserID:       maybe.Nothing[int32](),
				RegisteredBy: 1,
				Degree:       model.DegreeLikely,
			},
			{
				ID:           2,
				GameID:       1,
				UserID:       maybe.Nothing[int32](),
				RegisteredBy: 1,
				Degree:       model.DegreeUnlikely,
			},
		}, nil)

		got, err := fx.implementation.RegisterPlayer(fx.ctx, &gameplayer.RegisterPlayerRequest{
			GamePlayer: &gameplayer.GamePlayer{
				GameId: 1,
				UserId: &wrapperspb.Int32Value{
					Value: 1,
				},
				RegisteredBy: 1,
				Degree:       gameplayer.Degree_DEGREE_LIKELY,
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.FailedPrecondition, st.Code())
		assert.Len(t, st.Details(), 2)

		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, reasonThereAreNoFreeSlot, errorInfo.Reason)
		assert.Nil(t, errorInfo.Metadata)
	})

	t.Run("user already registered", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			ID:         1,
			MaxPlayers: 9,
			Registered: true,
		}, nil)

		fx.gamePlayersFacade.EXPECT().GetGamePlayersByGameID(fx.ctx, int32(1)).Return([]model.GamePlayer{
			{
				ID:           1,
				GameID:       1,
				UserID:       maybe.Nothing[int32](),
				RegisteredBy: 1,
				Degree:       model.DegreeLikely,
			},
			{
				ID:           2,
				GameID:       1,
				UserID:       maybe.Nothing[int32](),
				RegisteredBy: 1,
				Degree:       model.DegreeUnlikely,
			},
		}, nil)

		fx.gamePlayersFacade.EXPECT().CreateGamePlayer(fx.ctx, model.GamePlayer{
			GameID:       1,
			UserID:       maybe.Just(int32(1)),
			RegisteredBy: 1,
			Degree:       model.DegreeLikely,
		}).Return(model.GamePlayer{}, gameplayers.ErrGamePlayerAlreadyExists)

		got, err := fx.implementation.RegisterPlayer(fx.ctx, &gameplayer.RegisterPlayerRequest{
			GamePlayer: &gameplayer.GamePlayer{
				GameId: 1,
				UserId: &wrapperspb.Int32Value{
					Value: 1,
				},
				RegisteredBy: 1,
				Degree:       gameplayer.Degree_DEGREE_LIKELY,
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.AlreadyExists, st.Code())
		assert.Len(t, st.Details(), 2)

		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, reasonGamePlayerAlreadyRegistered, errorInfo.Reason)
		assert.Equal(t, map[string]string{
			"error": "game player already exists",
		}, errorInfo.Metadata)
	})

	t.Run("error game not found while create game player", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			ID:         1,
			MaxPlayers: 9,
			Registered: true,
		}, nil)

		fx.gamePlayersFacade.EXPECT().GetGamePlayersByGameID(fx.ctx, int32(1)).Return([]model.GamePlayer{
			{
				ID:           1,
				GameID:       1,
				UserID:       maybe.Nothing[int32](),
				RegisteredBy: 1,
				Degree:       model.DegreeLikely,
			},
			{
				ID:           2,
				GameID:       1,
				UserID:       maybe.Nothing[int32](),
				RegisteredBy: 1,
				Degree:       model.DegreeUnlikely,
			},
		}, nil)

		fx.gamePlayersFacade.EXPECT().CreateGamePlayer(fx.ctx, model.GamePlayer{
			GameID:       1,
			UserID:       maybe.Just(int32(1)),
			RegisteredBy: 1,
			Degree:       model.DegreeLikely,
		}).Return(model.GamePlayer{}, games.ErrGameNotFound)

		got, err := fx.implementation.RegisterPlayer(fx.ctx, &gameplayer.RegisterPlayerRequest{
			GamePlayer: &gameplayer.GamePlayer{
				GameId: 1,
				UserId: &wrapperspb.Int32Value{
					Value: 1,
				},
				RegisteredBy: 1,
				Degree:       gameplayer.Degree_DEGREE_LIKELY,
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
		assert.Equal(t, map[string]string{
			"error": "game not found",
		}, errorInfo.Metadata)
	})

	t.Run("error user not found while create game player", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			ID:         1,
			MaxPlayers: 9,
			Registered: true,
		}, nil)

		fx.gamePlayersFacade.EXPECT().GetGamePlayersByGameID(fx.ctx, int32(1)).Return([]model.GamePlayer{
			{
				ID:           1,
				GameID:       1,
				UserID:       maybe.Nothing[int32](),
				RegisteredBy: 1,
				Degree:       model.DegreeLikely,
			},
			{
				ID:           2,
				GameID:       1,
				UserID:       maybe.Nothing[int32](),
				RegisteredBy: 1,
				Degree:       model.DegreeUnlikely,
			},
		}, nil)

		fx.gamePlayersFacade.EXPECT().CreateGamePlayer(fx.ctx, model.GamePlayer{
			GameID:       1,
			UserID:       maybe.Just(int32(1)),
			RegisteredBy: 1,
			Degree:       model.DegreeLikely,
		}).Return(model.GamePlayer{}, users.ErrUserNotFound)

		got, err := fx.implementation.RegisterPlayer(fx.ctx, &gameplayer.RegisterPlayerRequest{
			GamePlayer: &gameplayer.GamePlayer{
				GameId: 1,
				UserId: &wrapperspb.Int32Value{
					Value: 1,
				},
				RegisteredBy: 1,
				Degree:       gameplayer.Degree_DEGREE_LIKELY,
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.FailedPrecondition, st.Code())
		assert.Len(t, st.Details(), 2)

		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, users.ReasonUserNotFound, errorInfo.Reason)
		assert.Equal(t, map[string]string{
			"error": "user not found",
		}, errorInfo.Metadata)
	})

	t.Run("internal error while create game player", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			ID:         1,
			MaxPlayers: 9,
			Registered: true,
		}, nil)

		fx.gamePlayersFacade.EXPECT().GetGamePlayersByGameID(fx.ctx, int32(1)).Return([]model.GamePlayer{
			{
				ID:           1,
				GameID:       1,
				UserID:       maybe.Nothing[int32](),
				RegisteredBy: 1,
				Degree:       model.DegreeLikely,
			},
			{
				ID:           2,
				GameID:       1,
				UserID:       maybe.Nothing[int32](),
				RegisteredBy: 1,
				Degree:       model.DegreeUnlikely,
			},
		}, nil)

		fx.gamePlayersFacade.EXPECT().CreateGamePlayer(fx.ctx, model.GamePlayer{
			GameID:       1,
			UserID:       maybe.Just(int32(1)),
			RegisteredBy: 1,
			Degree:       model.DegreeLikely,
		}).Return(model.GamePlayer{}, errors.New("some error"))

		got, err := fx.implementation.RegisterPlayer(fx.ctx, &gameplayer.RegisterPlayerRequest{
			GamePlayer: &gameplayer.GamePlayer{
				GameId: 1,
				UserId: &wrapperspb.Int32Value{
					Value: 1,
				},
				RegisteredBy: 1,
				Degree:       gameplayer.Degree_DEGREE_LIKELY,
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("ok main player", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			ID:         1,
			MaxPlayers: 9,
			Registered: true,
		}, nil)

		fx.gamePlayersFacade.EXPECT().GetGamePlayersByGameID(fx.ctx, int32(1)).Return([]model.GamePlayer{
			{
				ID:           1,
				GameID:       1,
				UserID:       maybe.Nothing[int32](),
				RegisteredBy: 1,
				Degree:       model.DegreeLikely,
			},
			{
				ID:           2,
				GameID:       1,
				UserID:       maybe.Nothing[int32](),
				RegisteredBy: 1,
				Degree:       model.DegreeUnlikely,
			},
		}, nil)

		fx.gamePlayersFacade.EXPECT().CreateGamePlayer(fx.ctx, model.GamePlayer{
			GameID:       1,
			UserID:       maybe.Just(int32(1)),
			RegisteredBy: 1,
			Degree:       model.DegreeLikely,
		}).Return(model.GamePlayer{
			ID:           1,
			GameID:       1,
			UserID:       maybe.Just(int32(1)),
			RegisteredBy: 1,
			Degree:       model.DegreeLikely,
		}, nil)

		got, err := fx.implementation.RegisterPlayer(fx.ctx, &gameplayer.RegisterPlayerRequest{
			GamePlayer: &gameplayer.GamePlayer{
				GameId: 1,
				UserId: &wrapperspb.Int32Value{
					Value: 1,
				},
				RegisteredBy: 1,
				Degree:       gameplayer.Degree_DEGREE_LIKELY,
			},
		})
		assert.Equal(t, &emptypb.Empty{}, got)
		assert.NoError(t, err)
	})

	t.Run("ok legioner", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			ID:         1,
			MaxPlayers: 9,
			Registered: true,
		}, nil)

		fx.gamePlayersFacade.EXPECT().GetGamePlayersByGameID(fx.ctx, int32(1)).Return([]model.GamePlayer{
			{
				ID:           1,
				GameID:       1,
				UserID:       maybe.Nothing[int32](),
				RegisteredBy: 1,
				Degree:       model.DegreeLikely,
			},
			{
				ID:           2,
				GameID:       1,
				UserID:       maybe.Nothing[int32](),
				RegisteredBy: 1,
				Degree:       model.DegreeUnlikely,
			},
		}, nil)

		fx.gamePlayersFacade.EXPECT().CreateGamePlayer(fx.ctx, model.GamePlayer{
			GameID:       1,
			UserID:       maybe.Nothing[int32](),
			RegisteredBy: 1,
			Degree:       model.DegreeLikely,
		}).Return(model.GamePlayer{
			ID:           1,
			GameID:       1,
			UserID:       maybe.Nothing[int32](),
			RegisteredBy: 1,
			Degree:       model.DegreeLikely,
		}, nil)

		got, err := fx.implementation.RegisterPlayer(fx.ctx, &gameplayer.RegisterPlayerRequest{
			GamePlayer: &gameplayer.GamePlayer{
				GameId:       1,
				RegisteredBy: 1,
				Degree:       gameplayer.Degree_DEGREE_LIKELY,
			},
		})
		assert.Equal(t, &emptypb.Empty{}, got)
		assert.NoError(t, err)
	})
}
