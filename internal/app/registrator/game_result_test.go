package registrator

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func TestRegistrator_CreateGameResult(t *testing.T) {
	t.Run("validation error. invalid game results round points JSON value", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.registrator.CreateGameResult(fx.ctx, &registrator.CreateGameResultRequest{
			GameResult: &registrator.GameResult{
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

		got, err := fx.registrator.CreateGameResult(fx.ctx, &registrator.CreateGameResultRequest{
			GameResult: &registrator.GameResult{
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
			RoundPoints: model.NewMaybeString("{}"),
		}).Return(model.GameResult{}, model.ErrGameNotFound)

		got, err := fx.registrator.CreateGameResult(fx.ctx, &registrator.CreateGameResultRequest{
			GameResult: &registrator.GameResult{
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
			RoundPoints: model.NewMaybeString("{}"),
		}).Return(model.GameResult{}, model.ErrGameResultAlreadyExists)

		got, err := fx.registrator.CreateGameResult(fx.ctx, &registrator.CreateGameResultRequest{
			GameResult: &registrator.GameResult{
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
			RoundPoints: model.NewMaybeString("{}"),
		}).Return(model.GameResult{
			ID:          1,
			FkGameID:    1,
			ResultPlace: 1,
			RoundPoints: model.NewMaybeString("{}"),
		}, nil)

		got, err := fx.registrator.CreateGameResult(fx.ctx, &registrator.CreateGameResultRequest{
			GameResult: &registrator.GameResult{
				GameId:      1,
				ResultPlace: 1,
				RoundPoints: "{}",
			},
		})

		assert.Equal(t, &registrator.GameResult{
			Id:          1,
			GameId:      1,
			ResultPlace: 1,
			RoundPoints: "{}",
		}, got)
		assert.NoError(t, err)
	})
}

func TestRegistrator_ListGameResults(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameResultsFacade.EXPECT().ListGameResults(fx.ctx).Return(nil, errors.New("some error"))

		got, err := fx.registrator.ListGameResults(fx.ctx, nil)

		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameResultsFacade.EXPECT().ListGameResults(fx.ctx).Return(
			[]model.GameResult{
				{
					ID:          1,
					FkGameID:    2,
					ResultPlace: 1,
					RoundPoints: model.NewMaybeString("{}"),
				},
				{
					ID:          2,
					FkGameID:    3,
					ResultPlace: 2,
					RoundPoints: model.NewMaybeString("{}"),
				},
			},
			nil)

		got, err := fx.registrator.ListGameResults(fx.ctx, nil)

		assert.Equal(t, got, &registrator.ListGameResultsResponse{
			GameResults: []*registrator.GameResult{
				{
					Id:          1,
					GameId:      2,
					ResultPlace: 1,
					RoundPoints: "{}",
				},
				{
					Id:          2,
					GameId:      3,
					ResultPlace: 2,
					RoundPoints: "{}",
				},
			},
		})
		assert.NoError(t, err)
	})
}

func TestRegistrator_PatchGameResults(t *testing.T) {
	t.Run("validation error. invalid game result round points JSON value", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.registrator.PatchGameResult(fx.ctx, &registrator.PatchGameResultRequest{
			GameResult: &registrator.GameResult{
				Id:          1,
				GameId:      2,
				ResultPlace: 3,
				RoundPoints: "invalid JSON value",
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"game_id",
					"result_place",
					"round_points",
				},
			},
		})

		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("validation error. invalid game result result place", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.registrator.PatchGameResult(fx.ctx, &registrator.PatchGameResultRequest{
			GameResult: &registrator.GameResult{
				Id:          1,
				GameId:      2,
				ResultPlace: 0,
				RoundPoints: "{}",
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"game_id",
					"result_place",
					"round_points",
				},
			},
		})

		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("patch error. game result not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameResultsFacade.EXPECT().PatchGameResult(fx.ctx, model.GameResult{
			ID:          1,
			FkGameID:    2,
			ResultPlace: 3,
			RoundPoints: model.NewMaybeString("{}"),
		}, []string{
			"game_id",
			"result_place",
			"round_points",
		}).Return(model.GameResult{}, model.ErrGameResultNotFound)

		got, err := fx.registrator.PatchGameResult(fx.ctx, &registrator.PatchGameResultRequest{
			GameResult: &registrator.GameResult{
				Id:          1,
				GameId:      2,
				ResultPlace: 3,
				RoundPoints: "{}",
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"game_id",
					"result_place",
					"round_points",
				},
			},
		})

		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("patch error. game not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameResultsFacade.EXPECT().PatchGameResult(fx.ctx, model.GameResult{
			ID:          1,
			FkGameID:    2,
			ResultPlace: 3,
			RoundPoints: model.NewMaybeString("{}"),
		}, []string{
			"game_id",
			"result_place",
			"round_points",
		}).Return(model.GameResult{}, model.ErrGameNotFound)

		got, err := fx.registrator.PatchGameResult(fx.ctx, &registrator.PatchGameResultRequest{
			GameResult: &registrator.GameResult{
				Id:          1,
				GameId:      2,
				ResultPlace: 3,
				RoundPoints: "{}",
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"game_id",
					"result_place",
					"round_points",
				},
			},
		})

		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("patch error. game result already exists", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameResultsFacade.EXPECT().PatchGameResult(fx.ctx, model.GameResult{
			ID:          1,
			FkGameID:    2,
			ResultPlace: 3,
			RoundPoints: model.NewMaybeString("{}"),
		}, []string{
			"game_id",
			"result_place",
			"round_points",
		}).Return(model.GameResult{}, model.ErrGameResultAlreadyExists)

		got, err := fx.registrator.PatchGameResult(fx.ctx, &registrator.PatchGameResultRequest{
			GameResult: &registrator.GameResult{
				Id:          1,
				GameId:      2,
				ResultPlace: 3,
				RoundPoints: "{}",
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"game_id",
					"result_place",
					"round_points",
				},
			},
		})

		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.AlreadyExists, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("patch error. other error", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameResultsFacade.EXPECT().PatchGameResult(fx.ctx, model.GameResult{
			ID:          1,
			FkGameID:    2,
			ResultPlace: 3,
			RoundPoints: model.NewMaybeString("{}"),
		}, []string{
			"game_id",
			"result_place",
			"round_points",
		}).Return(model.GameResult{}, errors.New("some error"))

		got, err := fx.registrator.PatchGameResult(fx.ctx, &registrator.PatchGameResultRequest{
			GameResult: &registrator.GameResult{
				Id:          1,
				GameId:      2,
				ResultPlace: 3,
				RoundPoints: "{}",
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"game_id",
					"result_place",
					"round_points",
				},
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

		fx.gameResultsFacade.EXPECT().PatchGameResult(fx.ctx, model.GameResult{
			ID:          1,
			FkGameID:    2,
			ResultPlace: 3,
			RoundPoints: model.NewMaybeString("{}"),
		}, []string{
			"game_id",
			"result_place",
			"round_points",
		}).Return(model.GameResult{
			ID:          1,
			FkGameID:    2,
			ResultPlace: 3,
			RoundPoints: model.NewMaybeString("{}"),
		}, nil)

		got, err := fx.registrator.PatchGameResult(fx.ctx, &registrator.PatchGameResultRequest{
			GameResult: &registrator.GameResult{
				Id:          1,
				GameId:      2,
				ResultPlace: 3,
				RoundPoints: "{}",
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"game_id",
					"result_place",
					"round_points",
				},
			},
		})

		assert.Equal(t, &registrator.GameResult{
			Id:          1,
			GameId:      2,
			ResultPlace: 3,
			RoundPoints: "{}",
		}, got)
		assert.NoError(t, err)
	})
}

func Test_convertModelGameResultToProtoGameResult(t *testing.T) {
	type args struct {
		gameResult model.GameResult
	}
	tests := []struct {
		name string
		args args
		want *registrator.GameResult
	}{
		{
			name: "tc1",
			args: args{
				gameResult: model.GameResult{
					ID:          1,
					FkGameID:    2,
					ResultPlace: 3,
					RoundPoints: model.NewMaybeString("{\"a\": 1}"),
				},
			},
			want: &registrator.GameResult{
				Id:          1,
				GameId:      2,
				ResultPlace: 3,
				RoundPoints: "{\"a\": 1}",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertModelGameResultToProtoGameResult(tt.args.gameResult); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertModelGameResultToProtoGameResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_validateCreateGameResultRequest(t *testing.T) {
	type args struct {
		ctx context.Context
		req *registrator.CreateGameResultRequest
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
				req: &registrator.CreateGameResultRequest{
					GameResult: &registrator.GameResult{
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
				req: &registrator.CreateGameResultRequest{
					GameResult: &registrator.GameResult{
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
				req: &registrator.CreateGameResultRequest{
					GameResult: &registrator.GameResult{
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
				req: &registrator.CreateGameResultRequest{
					GameResult: &registrator.GameResult{
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
