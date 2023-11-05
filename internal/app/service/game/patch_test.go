package game

import (
	"errors"
	"math"
	"testing"
	"time"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/leagues"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/places"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	gamepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game"
	leaguepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/league"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestImplementation_PatchGame(t *testing.T) {
	t.Run("error: bad request", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.implementation.PatchGame(fx.ctx, &gamepb.PatchGameRequest{})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("error: original game not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{}, games.ErrGameNotFound)

		got, err := fx.implementation.PatchGame(fx.ctx, &gamepb.PatchGameRequest{
			Game: &gamepb.Game{
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
		assert.Equal(t, "GAME_NOT_FOUND", errorInfo.Reason)
		assert.Equal(t, map[string]string{
			"error": "game not found",
		}, errorInfo.Metadata)
	})

	t.Run("error: internal error while getting original game", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{}, errors.New("some error"))

		got, err := fx.implementation.PatchGame(fx.ctx, &gamepb.PatchGameRequest{
			Game: &gamepb.Game{
				Id: 1,
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("", func(t *testing.T) {})

	t.Run("error: validation error", func(t *testing.T) {
		fx := tearUp(t)

		originalGame := model.NewGame()
		originalGame.ID = 1
		originalGame.Type = model.GameTypeClassic
		originalGame.Number = "1"
		originalGame.PlaceID = 1
		originalGame.Date = model.DateTime(time.Unix(1, 0))
		originalGame.MaxPlayers = 9
		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(originalGame, nil)

		got, err := fx.implementation.PatchGame(fx.ctx, &gamepb.PatchGameRequest{
			Game: &gamepb.Game{
				Id:       1,
				LeagueId: leaguepb.LeagueID_INVALID,
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"league_id",
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
		assert.Equal(t, "INVALID_LEAGUE_ID", errorInfo.Reason)
		assert.Equal(t, map[string]string{
			"error": "LeagueID: cannot be blank.",
		}, errorInfo.Metadata)
	})

	t.Run("error: game already exists", func(t *testing.T) {
		fx := tearUp(t)

		originalGame := model.NewGame()
		originalGame.ID = 1
		originalGame.Type = model.GameTypeClassic
		originalGame.Number = "1"
		originalGame.PlaceID = 1
		originalGame.Date = model.DateTime(time.Unix(1, 0))
		originalGame.MaxPlayers = 9
		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(originalGame, nil)

		fx.gamesFacade.EXPECT().PatchGame(fx.ctx, model.Game{
			ID:          1,
			ExternalID:  maybe.Just(int32(1)),
			LeagueID:    1,
			Type:        model.GameTypeClassic,
			Number:      "1",
			Name:        maybe.Just("name"),
			PlaceID:     1,
			Date:        model.DateTime(time.Unix(1, 0).UTC()),
			Price:       400,
			PaymentType: maybe.Just("cash"),
			MaxPlayers:  9,
			Payment:     maybe.Just(model.PaymentCertificate),
			Registered:  true,
			IsInMaster:  true,
			GameLink:    maybe.Nothing[string](),
		}).Return(model.Game{}, games.ErrGameAlreadyExists)

		paymentCertificate := gamepb.Payment_PAYMENT_CERTIFICATE
		got, err := fx.implementation.PatchGame(fx.ctx, &gamepb.PatchGameRequest{
			Game: &gamepb.Game{
				Id:         1,
				ExternalId: wrapperspb.Int32(1),
				LeagueId:   leaguepb.LeagueID_QUIZ_PLEASE,
				Type:       gamepb.GameType_GAME_TYPE_CLASSIC,
				Number:     "1",
				Name:       wrapperspb.String("name"),
				PlaceId:    1,
				Date: &timestamppb.Timestamp{
					Seconds: 1,
				},
				Price:       400,
				PaymentType: wrapperspb.String("cash"),
				MaxPlayers:  9,
				Payment:     &paymentCertificate,
				Registered:  true,
				IsInMaster:  true,
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"external_id",
					"league_id",
					"type",
					"number",
					"name",
					"place_id",
					"date",
					"price",
					"payment_type",
					"max_players",
					"payment",
					"registered",
					"is_in_master",
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
		assert.Equal(t, "GAME_ALREADY_EXISTS", errorInfo.Reason)
		assert.Equal(t, map[string]string{
			"error": "game already exists",
		}, errorInfo.Metadata)
	})

	t.Run("error: league not found", func(t *testing.T) {
		fx := tearUp(t)

		originalGame := model.NewGame()
		originalGame.ID = 1
		originalGame.Type = model.GameTypeClassic
		originalGame.Number = "1"
		originalGame.PlaceID = 1
		originalGame.Date = model.DateTime(time.Unix(1, 0))
		originalGame.MaxPlayers = 9
		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(originalGame, nil)

		patchedGame := model.NewGame()
		patchedGame.ID = 1
		patchedGame.LeagueID = math.MaxInt32
		patchedGame.Type = model.GameTypeClassic
		patchedGame.Number = "1"
		patchedGame.PlaceID = 1
		patchedGame.Date = model.DateTime(time.Unix(1, 0).UTC())
		patchedGame.Price = 400
		patchedGame.MaxPlayers = 9
		patchedGame.Registered = true
		patchedGame.IsInMaster = true
		fx.gamesFacade.EXPECT().PatchGame(fx.ctx, patchedGame).Return(model.Game{}, leagues.ErrLeagueNotFound)

		got, err := fx.implementation.PatchGame(fx.ctx, &gamepb.PatchGameRequest{
			Game: &gamepb.Game{
				Id:       1,
				LeagueId: leaguepb.LeagueID(math.MaxInt32),
				Type:     gamepb.GameType_GAME_TYPE_CLASSIC,
				Number:   "1",
				PlaceId:  1,
				Date: &timestamppb.Timestamp{
					Seconds: 1,
				},
				Price:      400,
				MaxPlayers: 9,
				Registered: true,
				IsInMaster: true,
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"external_id",
					"league_id",
					"type",
					"number",
					"name",
					"place_id",
					"date",
					"price",
					"payment_type",
					"max_players",
					"payment",
					"registered",
					"is_in_master",
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
		assert.Equal(t, "LEAGUE_NOT_FOUND", errorInfo.Reason)
		assert.Equal(t, map[string]string{
			"error": "league not found",
		}, errorInfo.Metadata)
	})

	t.Run("error: place not found", func(t *testing.T) {
		fx := tearUp(t)

		originalGame := model.NewGame()
		originalGame.ID = 1
		originalGame.LeagueID = 1
		originalGame.Type = model.GameTypeClassic
		originalGame.Number = "1"
		originalGame.Date = model.DateTime(time.Unix(1, 0))
		originalGame.MaxPlayers = 9
		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(originalGame, nil)

		patchedGame := originalGame
		patchedGame.PlaceID = math.MaxInt32
		fx.gamesFacade.EXPECT().PatchGame(fx.ctx, patchedGame).Return(model.Game{}, places.ErrPlaceNotFound)

		got, err := fx.implementation.PatchGame(fx.ctx, &gamepb.PatchGameRequest{
			Game: &gamepb.Game{
				Id:      1,
				PlaceId: math.MaxInt32,
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"place_id",
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
		assert.Equal(t, "PLACE_NOT_FOUND", errorInfo.Reason)
		assert.Equal(t, map[string]string{
			"error": "place not found",
		}, errorInfo.Metadata)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		originalGame := model.NewGame()
		originalGame.ID = 1
		originalGame.LeagueID = 1
		originalGame.Type = model.GameTypeClassic
		originalGame.Number = "1"
		originalGame.Date = model.DateTime(time.Unix(1, 0))
		originalGame.MaxPlayers = 9
		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(originalGame, nil)

		patchedGame := originalGame
		patchedGame.PlaceID = math.MaxInt32
		fx.gamesFacade.EXPECT().PatchGame(fx.ctx, patchedGame).Return(patchedGame, nil)

		got, err := fx.implementation.PatchGame(fx.ctx, &gamepb.PatchGameRequest{
			Game: &gamepb.Game{
				Id:      1,
				PlaceId: math.MaxInt32,
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"place_id",
				},
			},
		})
		assert.Equal(t, &gamepb.Game{
			Id:       1,
			LeagueId: leaguepb.LeagueID_QUIZ_PLEASE,
			Type:     gamepb.GameType_GAME_TYPE_CLASSIC,
			Number:   "1",
			PlaceId:  math.MaxInt32,
			Date: &timestamppb.Timestamp{
				Seconds: 1,
			},
			MaxPlayers: 9,
		}, got)
		assert.NoError(t, err)
	})
}
