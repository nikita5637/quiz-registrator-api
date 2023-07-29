package gameresultmanager

import (
	"context"
	"testing"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	gameresultmanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game_result_manager"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestRegistrator_CreateGameResult(t *testing.T) {
	t.Run("validation error. invalid game results round points JSON value", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.gameResultManager.CreateGameResult(fx.ctx, &gameresultmanagerpb.CreateGameResultRequest{
			GameResult: &gameresultmanagerpb.GameResult{
				GameId:      1,
				ResultPlace: 1,
				RoundPoints: "invalid JSON value",
			},
		})

		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("validation error. invalid game results result place", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.gameResultManager.CreateGameResult(fx.ctx, &gameresultmanagerpb.CreateGameResultRequest{
			GameResult: &gameresultmanagerpb.GameResult{
				GameId:      1,
				ResultPlace: 0,
				RoundPoints: "{}",
			},
		})

		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("create game result error. game not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameResultsFacade.EXPECT().CreateGameResult(fx.ctx, model.GameResult{
			FkGameID:    1,
			ResultPlace: 1,
			RoundPoints: maybe.Just("{}"),
		}).Return(model.GameResult{}, games.ErrGameNotFound)

		got, err := fx.gameResultManager.CreateGameResult(fx.ctx, &gameresultmanagerpb.CreateGameResultRequest{
			GameResult: &gameresultmanagerpb.GameResult{
				GameId:      1,
				ResultPlace: 1,
				RoundPoints: "{}",
			},
		})

		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("create game result error. game result already exists", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameResultsFacade.EXPECT().CreateGameResult(fx.ctx, model.GameResult{
			FkGameID:    1,
			ResultPlace: 1,
			RoundPoints: maybe.Just("{}"),
		}).Return(model.GameResult{}, model.ErrGameResultAlreadyExists)

		got, err := fx.gameResultManager.CreateGameResult(fx.ctx, &gameresultmanagerpb.CreateGameResultRequest{
			GameResult: &gameresultmanagerpb.GameResult{
				GameId:      1,
				ResultPlace: 1,
				RoundPoints: "{}",
			},
		})

		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.AlreadyExists, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameResultsFacade.EXPECT().CreateGameResult(fx.ctx, model.GameResult{
			FkGameID:    1,
			ResultPlace: 1,
			RoundPoints: maybe.Just("{}"),
		}).Return(model.GameResult{
			ID:          1,
			FkGameID:    1,
			ResultPlace: 1,
			RoundPoints: maybe.Just("{}"),
		}, nil)

		got, err := fx.gameResultManager.CreateGameResult(fx.ctx, &gameresultmanagerpb.CreateGameResultRequest{
			GameResult: &gameresultmanagerpb.GameResult{
				GameId:      1,
				ResultPlace: 1,
				RoundPoints: "{}",
			},
		})

		assert.Equal(t, &gameresultmanagerpb.GameResult{
			Id:          1,
			GameId:      1,
			ResultPlace: 1,
			RoundPoints: "{}",
		}, got)
		assert.NoError(t, err)
	})
}

func Test_validateCreateGameResultRequest(t *testing.T) {
	type args struct {
		ctx context.Context
		req *gameresultmanagerpb.CreateGameResultRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "invalid JSON",
			args: args{
				ctx: context.Background(),
				req: &gameresultmanagerpb.CreateGameResultRequest{
					GameResult: &gameresultmanagerpb.GameResult{
						GameId:      1,
						ResultPlace: 1,
						RoundPoints: "invalid JSON",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "result place lt 1",
			args: args{
				ctx: context.Background(),
				req: &gameresultmanagerpb.CreateGameResultRequest{
					GameResult: &gameresultmanagerpb.GameResult{
						GameId:      1,
						ResultPlace: 0,
						RoundPoints: "{}",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "len of JSON string gt 256",
			args: args{
				ctx: context.Background(),
				req: &gameresultmanagerpb.CreateGameResultRequest{
					GameResult: &gameresultmanagerpb.GameResult{
						GameId:      1,
						ResultPlace: 1,
						RoundPoints: "{\"a\":\"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\"}",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "ok",
			args: args{
				ctx: context.Background(),
				req: &gameresultmanagerpb.CreateGameResultRequest{
					GameResult: &gameresultmanagerpb.GameResult{
						GameId:      1,
						ResultPlace: 1,
						RoundPoints: "{}",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateCreateGameResultRequest(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("validateCreateGameResultRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
