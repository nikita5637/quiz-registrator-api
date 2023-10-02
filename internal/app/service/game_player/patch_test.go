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
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestImplementation_PatchGamePlayer(t *testing.T) {
	t.Run("game paleyr is nil", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.implementation.PatchGamePlayer(fx.ctx, &gameplayer.PatchGamePlayerRequest{})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("game player not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePlayersFacade.EXPECT().GetGamePlayer(fx.ctx, int32(1)).Return(model.GamePlayer{}, gameplayers.ErrGamePlayerNotFound)

		got, err := fx.implementation.PatchGamePlayer(fx.ctx, &gameplayer.PatchGamePlayerRequest{
			GamePlayer: &gameplayer.GamePlayer{
				Id: 1,
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())
		assert.Len(t, st.Details(), 2)

		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, gameplayers.ReasonGamePlayerNotFound, errorInfo.Reason)
		assert.Nil(t, errorInfo.Metadata)
	})

	t.Run("internal error while getting game player", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePlayersFacade.EXPECT().GetGamePlayer(fx.ctx, int32(1)).Return(model.GamePlayer{}, errors.New("some error"))

		got, err := fx.implementation.PatchGamePlayer(fx.ctx, &gameplayer.PatchGamePlayerRequest{
			GamePlayer: &gameplayer.GamePlayer{
				Id: 1,
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("validation error", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePlayersFacade.EXPECT().GetGamePlayer(fx.ctx, int32(1)).Return(model.GamePlayer{
			ID:           1,
			GameID:       1,
			UserID:       maybe.Just(int32(1)),
			RegisteredBy: 1,
			Degree:       model.DegreeLikely,
		}, nil)

		got, err := fx.implementation.PatchGamePlayer(fx.ctx, &gameplayer.PatchGamePlayerRequest{
			GamePlayer: &gameplayer.GamePlayer{
				Id:     1,
				GameId: 0,
				UserId: &wrapperspb.Int32Value{
					Value: 1,
				},
				RegisteredBy: 1,
				Degree:       gameplayer.Degree_DEGREE_LIKELY,
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{"game_id"},
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

	t.Run("user not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePlayersFacade.EXPECT().GetGamePlayer(fx.ctx, int32(1)).Return(model.GamePlayer{
			ID:           1,
			GameID:       1,
			UserID:       maybe.Just(int32(1)),
			RegisteredBy: 1,
			Degree:       model.DegreeLikely,
		}, nil)

		fx.gamePlayersFacade.EXPECT().PatchGamePlayer(fx.ctx, model.GamePlayer{
			ID:           1,
			GameID:       1,
			UserID:       maybe.Just(int32(1)),
			RegisteredBy: 1,
			Degree:       model.DegreeUnlikely,
		}).Return(model.GamePlayer{}, users.ErrUserNotFound)

		got, err := fx.implementation.PatchGamePlayer(fx.ctx, &gameplayer.PatchGamePlayerRequest{
			GamePlayer: &gameplayer.GamePlayer{
				Id:     1,
				GameId: 1,
				UserId: &wrapperspb.Int32Value{
					Value: 1,
				},
				RegisteredBy: 1,
				Degree:       gameplayer.Degree_DEGREE_UNLIKELY,
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{"degree"},
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
		assert.Nil(t, errorInfo.Metadata)
	})

	t.Run("game not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePlayersFacade.EXPECT().GetGamePlayer(fx.ctx, int32(1)).Return(model.GamePlayer{
			ID:           1,
			GameID:       1,
			UserID:       maybe.Just(int32(1)),
			RegisteredBy: 1,
			Degree:       model.DegreeLikely,
		}, nil)

		fx.gamePlayersFacade.EXPECT().PatchGamePlayer(fx.ctx, model.GamePlayer{
			ID:           1,
			GameID:       1,
			UserID:       maybe.Just(int32(1)),
			RegisteredBy: 1,
			Degree:       model.DegreeUnlikely,
		}).Return(model.GamePlayer{}, games.ErrGameNotFound)

		got, err := fx.implementation.PatchGamePlayer(fx.ctx, &gameplayer.PatchGamePlayerRequest{
			GamePlayer: &gameplayer.GamePlayer{
				Id:     1,
				GameId: 1,
				UserId: &wrapperspb.Int32Value{
					Value: 1,
				},
				RegisteredBy: 1,
				Degree:       gameplayer.Degree_DEGREE_UNLIKELY,
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{"degree"},
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

	t.Run("game player already exists", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePlayersFacade.EXPECT().GetGamePlayer(fx.ctx, int32(1)).Return(model.GamePlayer{
			ID:           1,
			GameID:       1,
			UserID:       maybe.Just(int32(1)),
			RegisteredBy: 1,
			Degree:       model.DegreeLikely,
		}, nil)

		fx.gamePlayersFacade.EXPECT().PatchGamePlayer(fx.ctx, model.GamePlayer{
			ID:           1,
			GameID:       1,
			UserID:       maybe.Just(int32(1)),
			RegisteredBy: 1,
			Degree:       model.DegreeUnlikely,
		}).Return(model.GamePlayer{}, gameplayers.ErrGamePlayerAlreadyExists)

		got, err := fx.implementation.PatchGamePlayer(fx.ctx, &gameplayer.PatchGamePlayerRequest{
			GamePlayer: &gameplayer.GamePlayer{
				Id:     1,
				GameId: 1,
				UserId: &wrapperspb.Int32Value{
					Value: 1,
				},
				RegisteredBy: 1,
				Degree:       gameplayer.Degree_DEGREE_UNLIKELY,
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{"degree"},
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.AlreadyExists, st.Code())
		assert.Len(t, st.Details(), 2)

		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, gameplayers.ReasonGamePlayerAlreadyExists, errorInfo.Reason)
		assert.Nil(t, errorInfo.Metadata)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePlayersFacade.EXPECT().GetGamePlayer(fx.ctx, int32(1)).Return(model.GamePlayer{
			ID:           1,
			GameID:       1,
			UserID:       maybe.Nothing[int32](),
			RegisteredBy: 1,
			Degree:       model.DegreeLikely,
		}, nil)

		fx.gamePlayersFacade.EXPECT().PatchGamePlayer(fx.ctx, model.GamePlayer{
			ID:           1,
			GameID:       2,
			UserID:       maybe.Just(int32(2)),
			RegisteredBy: 2,
			Degree:       model.DegreeUnlikely,
		}).Return(model.GamePlayer{
			ID:           1,
			GameID:       2,
			UserID:       maybe.Just(int32(2)),
			RegisteredBy: 2,
			Degree:       model.DegreeUnlikely,
		}, nil)

		got, err := fx.implementation.PatchGamePlayer(fx.ctx, &gameplayer.PatchGamePlayerRequest{
			GamePlayer: &gameplayer.GamePlayer{
				Id:     1,
				GameId: 2,
				UserId: &wrapperspb.Int32Value{
					Value: 2,
				},
				RegisteredBy: 2,
				Degree:       gameplayer.Degree_DEGREE_UNLIKELY,
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{"game_id", "user_id", "registered_by", "degree"},
			},
		})
		assert.Equal(t, &gameplayerpb.GamePlayer{
			Id:     1,
			GameId: 2,
			UserId: &wrapperspb.Int32Value{
				Value: 2,
			},
			RegisteredBy: 2,
			Degree:       gameplayer.Degree_DEGREE_UNLIKELY,
		}, got)
		assert.NoError(t, err)
	})
}

func Test_validatePatchedGamePlayer(t *testing.T) {
	type args struct {
		gamePlayer model.GamePlayer
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "no ID",
			args: args{
				gamePlayer: model.GamePlayer{
					UserID:       maybe.Nothing[int32](),
					GameID:       1,
					RegisteredBy: 1,
					Degree:       model.DegreeLikely,
				},
			},
			wantErr: true,
		},
		{
			name: "game ID lt 0",
			args: args{
				gamePlayer: model.GamePlayer{
					ID:           1,
					GameID:       -1,
					UserID:       maybe.Nothing[int32](),
					RegisteredBy: 1,
					Degree:       model.DegreeLikely,
				},
			},
			wantErr: true,
		},
		{
			name: "game ID lt 0",
			args: args{
				gamePlayer: model.GamePlayer{
					ID:           1,
					GameID:       -1,
					UserID:       maybe.Nothing[int32](),
					RegisteredBy: 1,
					Degree:       model.DegreeLikely,
				},
			},
			wantErr: true,
		},
		{
			name: "game ID eq 0",
			args: args{
				gamePlayer: model.GamePlayer{
					ID:           1,
					GameID:       0,
					UserID:       maybe.Nothing[int32](),
					RegisteredBy: 1,
					Degree:       model.DegreeLikely,
				},
			},
			wantErr: true,
		},
		{
			name: "user ID lt 0",
			args: args{
				gamePlayer: model.GamePlayer{
					ID:           1,
					GameID:       1,
					UserID:       maybe.Just(int32(-1)),
					RegisteredBy: 1,
					Degree:       model.DegreeLikely,
				},
			},
			wantErr: true,
		},
		{
			name: "user ID eq 0",
			args: args{
				gamePlayer: model.GamePlayer{
					ID:           1,
					GameID:       1,
					UserID:       maybe.Just(int32(0)),
					RegisteredBy: 1,
					Degree:       model.DegreeLikely,
				},
			},
			wantErr: true,
		},
		{
			name: "registered by lt 0",
			args: args{
				gamePlayer: model.GamePlayer{
					ID:           1,
					GameID:       1,
					UserID:       maybe.Just(int32(1)),
					RegisteredBy: -1,
					Degree:       model.DegreeLikely,
				},
			},
			wantErr: true,
		},
		{
			name: "registered by eq 0",
			args: args{
				gamePlayer: model.GamePlayer{
					ID:           1,
					GameID:       1,
					UserID:       maybe.Just(int32(1)),
					RegisteredBy: 0,
					Degree:       model.DegreeLikely,
				},
			},
			wantErr: true,
		},
		{
			name: "degree eq 0",
			args: args{
				gamePlayer: model.GamePlayer{
					ID:           1,
					GameID:       1,
					UserID:       maybe.Just(int32(1)),
					RegisteredBy: 1,
					Degree:       0,
				},
			},
			wantErr: true,
		},
		{
			name: "ok legioner",
			args: args{
				gamePlayer: model.GamePlayer{
					ID:           1,
					GameID:       1,
					UserID:       maybe.Nothing[int32](),
					RegisteredBy: 1,
					Degree:       model.DegreeLikely,
				},
			},
			wantErr: false,
		},
		{
			name: "ok main player",
			args: args{
				gamePlayer: model.GamePlayer{
					ID:           1,
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
			if err := validatePatchedGamePlayer(tt.args.gamePlayer); (err != nil) != tt.wantErr {
				t.Errorf("validatePatchedGamePlayer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
