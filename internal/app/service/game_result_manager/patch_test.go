package gameresultmanager

import (
	"errors"
	"testing"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/gameresults"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	gameresultmanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game_result_manager"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
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

		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, "invalid game result round points JSON value: \"invalid JSON value\"", errorInfo.Reason)
		assert.Nil(t, errorInfo.Metadata)
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

		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, "invalid game result result place: \"0\"", errorInfo.Reason)
		assert.Nil(t, errorInfo.Metadata)
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
		}).Return(model.GameResult{}, gameresults.ErrGameResultNotFound)

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

		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, "GAME_RESULT_NOT_FOUND", errorInfo.Reason)
		assert.Equal(t, map[string]string{
			"error": "game result not found",
		}, errorInfo.Metadata)
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
		}).Return(model.GameResult{}, games.ErrGameNotFound)

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
		assert.Equal(t, codes.FailedPrecondition, st.Code())
		assert.Len(t, st.Details(), 2)

		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, games.ReasonGameNotFound, errorInfo.Reason)
		assert.Equal(t, map[string]string{
			"error": "game not found",
		}, errorInfo.Metadata)
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
		}).Return(model.GameResult{}, gameresults.ErrGameResultAlreadyExists)

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

		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, gameresults.ReasonGameResultAlreadyExists, errorInfo.Reason)
		assert.Equal(t, map[string]string{
			"error": "game result already exists",
		}, errorInfo.Metadata)
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
