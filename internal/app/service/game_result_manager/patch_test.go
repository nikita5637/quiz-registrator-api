package gameresultmanager

import (
	"errors"
	"testing"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	gameresultmanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game_result_manager"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func TestRegistrator_PatchGameResults(t *testing.T) {
	t.Run("validation error. invalid game result round points JSON value", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.gameResultManager.PatchGameResult(fx.ctx, &gameresultmanagerpb.PatchGameResultRequest{
			GameResult: &gameresultmanagerpb.GameResult{
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

		got, err := fx.gameResultManager.PatchGameResult(fx.ctx, &gameresultmanagerpb.PatchGameResultRequest{
			GameResult: &gameresultmanagerpb.GameResult{
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
			RoundPoints: maybe.Just("{}"),
		}, []string{
			"game_id",
			"result_place",
			"round_points",
		}).Return(model.GameResult{}, model.ErrGameResultNotFound)

		got, err := fx.gameResultManager.PatchGameResult(fx.ctx, &gameresultmanagerpb.PatchGameResultRequest{
			GameResult: &gameresultmanagerpb.GameResult{
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
			RoundPoints: maybe.Just("{}"),
		}, []string{
			"game_id",
			"result_place",
			"round_points",
		}).Return(model.GameResult{}, model.ErrGameNotFound)

		got, err := fx.gameResultManager.PatchGameResult(fx.ctx, &gameresultmanagerpb.PatchGameResultRequest{
			GameResult: &gameresultmanagerpb.GameResult{
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
			RoundPoints: maybe.Just("{}"),
		}, []string{
			"game_id",
			"result_place",
			"round_points",
		}).Return(model.GameResult{}, model.ErrGameResultAlreadyExists)

		got, err := fx.gameResultManager.PatchGameResult(fx.ctx, &gameresultmanagerpb.PatchGameResultRequest{
			GameResult: &gameresultmanagerpb.GameResult{
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
			RoundPoints: maybe.Just("{}"),
		}, []string{
			"game_id",
			"result_place",
			"round_points",
		}).Return(model.GameResult{}, errors.New("some error"))

		got, err := fx.gameResultManager.PatchGameResult(fx.ctx, &gameresultmanagerpb.PatchGameResultRequest{
			GameResult: &gameresultmanagerpb.GameResult{
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
			RoundPoints: maybe.Just("{}"),
		}, []string{
			"game_id",
			"result_place",
			"round_points",
		}).Return(model.GameResult{
			ID:          1,
			FkGameID:    2,
			ResultPlace: 3,
			RoundPoints: maybe.Just("{}"),
		}, nil)

		got, err := fx.gameResultManager.PatchGameResult(fx.ctx, &gameresultmanagerpb.PatchGameResultRequest{
			GameResult: &gameresultmanagerpb.GameResult{
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

		assert.Equal(t, &gameresultmanagerpb.GameResult{
			Id:          1,
			GameId:      2,
			ResultPlace: 3,
			RoundPoints: "{}",
		}, got)
		assert.NoError(t, err)
	})
}
