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
	gameplayerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game_player"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestImplementation_CreateGamePlayer(t *testing.T) {
	t.Run("game player is nil", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.implementation.CreateGamePlayer(fx.ctx, &gameplayer.CreateGamePlayerRequest{})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("validation error", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.implementation.CreateGamePlayer(fx.ctx, &gameplayer.CreateGamePlayerRequest{
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

	t.Run("user already registered", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePlayersFacade.EXPECT().CreateGamePlayer(fx.ctx, model.GamePlayer{
			GameID:       1,
			UserID:       maybe.Just(int32(1)),
			RegisteredBy: 1,
			Degree:       model.DegreeLikely,
		}).Return(model.GamePlayer{}, gameplayers.ErrGamePlayerAlreadyRegistered)

		got, err := fx.implementation.CreateGamePlayer(fx.ctx, &gameplayer.CreateGamePlayerRequest{
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
		assert.Equal(t, gameplayers.ReasonGamePlayerAlreadyRegistered, errorInfo.Reason)
		assert.Nil(t, errorInfo.Metadata)
	})

	t.Run("error game not found while create game player", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePlayersFacade.EXPECT().CreateGamePlayer(fx.ctx, model.GamePlayer{
			GameID:       1,
			UserID:       maybe.Just(int32(1)),
			RegisteredBy: 1,
			Degree:       model.DegreeLikely,
		}).Return(model.GamePlayer{}, games.ErrGameNotFound)

		got, err := fx.implementation.CreateGamePlayer(fx.ctx, &gameplayer.CreateGamePlayerRequest{
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
		assert.Equal(t, codes.InvalidArgument, st.Code())
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

		fx.gamePlayersFacade.EXPECT().CreateGamePlayer(fx.ctx, model.GamePlayer{
			GameID:       1,
			UserID:       maybe.Just(int32(1)),
			RegisteredBy: 1,
			Degree:       model.DegreeLikely,
		}).Return(model.GamePlayer{}, users.ErrUserNotFound)

		got, err := fx.implementation.CreateGamePlayer(fx.ctx, &gameplayer.CreateGamePlayerRequest{
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
		assert.Equal(t, codes.InvalidArgument, st.Code())
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

		fx.gamePlayersFacade.EXPECT().CreateGamePlayer(fx.ctx, model.GamePlayer{
			GameID:       1,
			UserID:       maybe.Just(int32(1)),
			RegisteredBy: 1,
			Degree:       model.DegreeLikely,
		}).Return(model.GamePlayer{}, errors.New("some error"))

		got, err := fx.implementation.CreateGamePlayer(fx.ctx, &gameplayer.CreateGamePlayerRequest{
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

		got, err := fx.implementation.CreateGamePlayer(fx.ctx, &gameplayer.CreateGamePlayerRequest{
			GamePlayer: &gameplayer.GamePlayer{
				GameId: 1,
				UserId: &wrapperspb.Int32Value{
					Value: 1,
				},
				RegisteredBy: 1,
				Degree:       gameplayer.Degree_DEGREE_LIKELY,
			},
		})
		assert.Equal(t, &gameplayerpb.GamePlayer{
			Id:     1,
			GameId: 1,
			UserId: &wrapperspb.Int32Value{
				Value: 1,
			},
			RegisteredBy: 1,
			Degree:       gameplayer.Degree_DEGREE_LIKELY,
		}, got)
		assert.NoError(t, err)
	})

	t.Run("ok legioner", func(t *testing.T) {
		fx := tearUp(t)

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

		got, err := fx.implementation.CreateGamePlayer(fx.ctx, &gameplayer.CreateGamePlayerRequest{
			GamePlayer: &gameplayer.GamePlayer{
				GameId:       1,
				RegisteredBy: 1,
				Degree:       gameplayer.Degree_DEGREE_LIKELY,
			},
		})
		assert.Equal(t, &gameplayerpb.GamePlayer{
			Id:           1,
			GameId:       1,
			RegisteredBy: 1,
			Degree:       gameplayer.Degree_DEGREE_LIKELY,
		}, got)
		assert.NoError(t, err)
	})
}

func Test_validateCreatedGamePlayer(t *testing.T) {
	type args struct {
		gamePlayer model.GamePlayer
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "empty game ID",
			args: args{
				gamePlayer: model.GamePlayer{
					UserID:       maybe.Just(int32(1)),
					RegisteredBy: 1,
					Degree:       model.DegreeLikely,
				},
			},
			wantErr: true,
		},
		{
			name: "empty user ID",
			args: args{
				gamePlayer: model.GamePlayer{
					GameID:       1,
					UserID:       maybe.Nothing[int32](),
					RegisteredBy: 1,
					Degree:       model.DegreeLikely,
				},
			},
			wantErr: false,
		},
		{
			name: "user ID eq 0",
			args: args{
				gamePlayer: model.GamePlayer{
					GameID:       1,
					UserID:       maybe.Just(int32(0)),
					RegisteredBy: 1,
					Degree:       model.DegreeLikely,
				},
			},
			wantErr: true,
		},
		{
			name: "registered by eq 0",
			args: args{
				gamePlayer: model.GamePlayer{
					GameID:       1,
					UserID:       maybe.Just(int32(1)),
					RegisteredBy: 0,
					Degree:       model.DegreeLikely,
				},
			},
			wantErr: true,
		},
		{
			name: "invalid degree",
			args: args{
				gamePlayer: model.GamePlayer{
					GameID:       1,
					UserID:       maybe.Just(int32(1)),
					RegisteredBy: 1,
					Degree:       model.DegreeInvalid,
				},
			},
			wantErr: true,
		},
		{
			name: "ok",
			args: args{
				gamePlayer: model.GamePlayer{
					GameID:       1,
					UserID:       maybe.Just(int32(1)),
					RegisteredBy: 1,
					Degree:       model.DegreeLikely,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateCreatedGamePlayer(tt.args.gamePlayer); (err != nil) != tt.wantErr {
				t.Errorf("validateCreatedGamePlayer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
