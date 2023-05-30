package registrator

import (
	"errors"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	pkgmodel "github.com/nikita5637/quiz-registrator-api/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"
	time_utils "github.com/nikita5637/quiz-registrator-api/utils/time"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestRegistrator_AddGame(t *testing.T) {
	t.Run("error invalid leagueID", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().AddGame(fx.ctx, model.Game{
			PlaceID: 1,
		}).Return(0, model.ErrInvalidLeagueID)

		got, err := fx.registrator.AddGame(fx.ctx, &registrator.AddGameRequest{
			PlaceId: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("error invalid game type", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().AddGame(fx.ctx, model.Game{
			PlaceID: 1,
		}).Return(0, model.ErrInvalidGameType)

		got, err := fx.registrator.AddGame(fx.ctx, &registrator.AddGameRequest{
			PlaceId: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("error invalid game number", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().AddGame(fx.ctx, model.Game{
			PlaceID: 1,
		}).Return(0, model.ErrInvalidGameNumber)

		got, err := fx.registrator.AddGame(fx.ctx, &registrator.AddGameRequest{
			PlaceId: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("error invalid place id", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().AddGame(fx.ctx, model.Game{
			PlaceID: 1,
		}).Return(0, model.ErrInvalidPlaceID)

		got, err := fx.registrator.AddGame(fx.ctx, &registrator.AddGameRequest{
			PlaceId: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("error invalid date", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().AddGame(fx.ctx, model.Game{
			PlaceID: 1,
		}).Return(0, model.ErrInvalidDate)

		got, err := fx.registrator.AddGame(fx.ctx, &registrator.AddGameRequest{
			PlaceId: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("error invalid price", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().AddGame(fx.ctx, model.Game{
			PlaceID: 1,
		}).Return(0, model.ErrInvalidPrice)

		got, err := fx.registrator.AddGame(fx.ctx, &registrator.AddGameRequest{
			PlaceId: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("error invalid max players", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().AddGame(fx.ctx, model.Game{
			PlaceID: 1,
		}).Return(0, model.ErrInvalidMaxPlayers)

		got, err := fx.registrator.AddGame(fx.ctx, &registrator.AddGameRequest{
			PlaceId: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("internal error while add game", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().AddGame(fx.ctx, model.Game{
			PlaceID: 1,
		}).Return(0, errors.New("some error"))

		got, err := fx.registrator.AddGame(fx.ctx, &registrator.AddGameRequest{
			PlaceId: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().AddGame(fx.ctx, model.Game{
			PlaceID: 2,
		}).Return(1, nil)

		got, err := fx.registrator.AddGame(fx.ctx, &registrator.AddGameRequest{
			PlaceId: 2,
		})
		assert.Equal(t, &registrator.AddGameResponse{
			Id: 1,
		}, got)
		assert.NoError(t, err)
	})
}

func TestRegistrator_AddGames(t *testing.T) {
	t.Run("internal error with skip game while add games", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().AddGames(fx.ctx, []model.Game{
			{
				ExternalID: 2,
				LeagueID:   2,
				Type:       model.GameTypeClassic,
				Number:     "2",
				PlaceID:    1,
				Date:       model.DateTime(time_utils.ConvertTime("2023-01-02 16:30")),
				Price:      400,
				MaxPlayers: 8,
			},
			{
				ExternalID: 4,
				LeagueID:   4,
				Type:       model.GameTypeMoviesAndMusic,
				Number:     "4",
				PlaceID:    1,
				Date:       model.DateTime(time_utils.ConvertTime("2023-01-03 13:00")),
				Price:      400,
				MaxPlayers: 8,
			},
		}).Return(errors.New("some error"))

		got, err := fx.registrator.AddGames(fx.ctx, &registrator.AddGamesRequest{
			Games: []*registrator.AddGamesRequest_Game{
				{
					ExternalId: 1,
					LeagueId:   1,
				},
				{
					ExternalId: 2,
					LeagueId:   2,
					Type:       registrator.GameType(model.GameTypeClassic),
					Number:     "2",
					PlaceId:    1,
					Date:       timestamppb.New(time_utils.ConvertTime("2023-01-02 16:30")),
					Price:      400,
					MaxPlayers: 8,
				},
				{
					ExternalId: 3,
					LeagueId:   3,
				},
				{
					ExternalId: 4,
					LeagueId:   4,
					Type:       registrator.GameType(model.GameTypeMoviesAndMusic),
					Number:     "4",
					PlaceId:    1,
					Date:       timestamppb.New(time_utils.ConvertTime("2023-01-03 13:00")),
					Price:      400,
					MaxPlayers: 8,
				},
			},
		})

		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})
	t.Run("ok with skip game", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().AddGames(fx.ctx, []model.Game{
			{
				ExternalID: 2,
				LeagueID:   2,
				Type:       model.GameTypeClassic,
				Number:     "2",
				PlaceID:    1,
				Date:       model.DateTime(time_utils.ConvertTime("2023-01-02 16:30")),
				Price:      400,
				MaxPlayers: 8,
			},
			{
				ExternalID: 4,
				LeagueID:   4,
				Type:       model.GameTypeMoviesAndMusic,
				Number:     "4",
				PlaceID:    1,
				Date:       model.DateTime(time_utils.ConvertTime("2023-01-03 13:00")),
				Price:      400,
				MaxPlayers: 8,
			},
		}).Return(nil)

		got, err := fx.registrator.AddGames(fx.ctx, &registrator.AddGamesRequest{
			Games: []*registrator.AddGamesRequest_Game{
				{
					ExternalId: 1,
					LeagueId:   1,
				},
				{
					ExternalId: 2,
					LeagueId:   2,
					Type:       registrator.GameType(model.GameTypeClassic),
					Number:     "2",
					PlaceId:    1,
					Date:       timestamppb.New(time_utils.ConvertTime("2023-01-02 16:30")),
					Price:      400,
					MaxPlayers: 8,
				},
				{
					ExternalId: 3,
					LeagueId:   3,
				},
				{
					ExternalId: 4,
					LeagueId:   4,
					Type:       registrator.GameType(model.GameTypeMoviesAndMusic),
					Number:     "4",
					PlaceId:    1,
					Date:       timestamppb.New(time_utils.ConvertTime("2023-01-03 13:00")),
					Price:      400,
					MaxPlayers: 8,
				},
			},
		})

		assert.NotNil(t, got)
		assert.NoError(t, err)
	})

	t.Run("ok without skip game", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().AddGames(fx.ctx, []model.Game{
			{
				ExternalID: 2,
				LeagueID:   2,
				Type:       model.GameTypeClassic,
				Number:     "2",
				PlaceID:    1,
				Date:       model.DateTime(time_utils.ConvertTime("2023-01-02 16:30")),
				Price:      400,
				MaxPlayers: 8,
			},
			{
				ExternalID: 4,
				LeagueID:   4,
				Type:       model.GameTypeMoviesAndMusic,
				Number:     "4",
				PlaceID:    1,
				Date:       model.DateTime(time_utils.ConvertTime("2023-01-03 13:00")),
				Price:      400,
				MaxPlayers: 8,
			},
		}).Return(nil)

		got, err := fx.registrator.AddGames(fx.ctx, &registrator.AddGamesRequest{
			Games: []*registrator.AddGamesRequest_Game{
				{
					ExternalId: 2,
					LeagueId:   2,
					Type:       registrator.GameType(model.GameTypeClassic),
					Number:     "2",
					PlaceId:    1,
					Date:       timestamppb.New(time_utils.ConvertTime("2023-01-02 16:30")),
					Price:      400,
					MaxPlayers: 8,
				},
				{
					ExternalId: 4,
					LeagueId:   4,
					Type:       registrator.GameType(model.GameTypeMoviesAndMusic),
					Number:     "4",
					PlaceId:    1,
					Date:       timestamppb.New(time_utils.ConvertTime("2023-01-03 13:00")),
					Price:      400,
					MaxPlayers: 8,
				},
			},
		})

		assert.NotNil(t, got)
		assert.NoError(t, err)
	})
}

func TestRegistrator_DeleteGame(t *testing.T) {
	t.Run("error \"game not found\" while deleting game", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().DeleteGame(fx.ctx, int32(1)).Return(model.ErrGameNotFound)

		got, err := fx.registrator.DeleteGame(fx.ctx, &registrator.DeleteGameRequest{
			Id: 1,
		})
		assert.Nil(t, got)

		st := status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("internal error while deleting game", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().DeleteGame(fx.ctx, int32(1)).Return(errors.New("some error"))

		got, err := fx.registrator.DeleteGame(fx.ctx, &registrator.DeleteGameRequest{
			Id: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().DeleteGame(fx.ctx, int32(1)).Return(nil)

		got, err := fx.registrator.DeleteGame(fx.ctx, &registrator.DeleteGameRequest{
			Id: 1,
		})
		assert.NotNil(t, got)
		assert.NoError(t, err)
	})
}

func TestRegistrator_GetGameByID(t *testing.T) {
	t.Run("error game not found while get game by id", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{}, model.ErrGameNotFound)

		got, err := fx.registrator.GetGameByID(fx.ctx, &registrator.GetGameByIDRequest{
			GameId: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("internal error while get game by id", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{}, errors.New("some error"))

		got, err := fx.registrator.GetGameByID(fx.ctx, &registrator.GetGameByIDRequest{
			GameId: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		timeNow := time_utils.TimeNow()
		game := model.Game{
			ID:          1,
			ExternalID:  2,
			LeagueID:    pkgmodel.LeagueQuizPlease,
			Type:        model.GameTypeClassic,
			Number:      "number",
			Name:        "name",
			PlaceID:     3,
			Date:        model.DateTime(timeNow),
			Price:       400,
			PaymentType: "cash,card",
			MaxPlayers:  9,
			Payment:     int32(registrator.Payment_PAYMENT_CASH),
			Registered:  true,
		}

		game.My = true
		game.NumberOfMyLegioners = 3
		game.NumberOfLegioners = 6
		game.NumberOfPlayers = 9

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(game, nil)

		got, err := fx.registrator.GetGameByID(fx.ctx, &registrator.GetGameByIDRequest{
			GameId: 1,
		})
		assert.NotNil(t, got)
		assert.Equal(t, &registrator.GetGameByIDResponse{
			Game: &registrator.Game{
				Id:                  1,
				ExternalId:          2,
				LeagueId:            pkgmodel.LeagueQuizPlease,
				Type:                registrator.GameType_GAME_TYPE_CLASSIC,
				Number:              "number",
				Name:                "name",
				PlaceId:             3,
				Date:                timestamppb.New(timeNow),
				Price:               400,
				PaymentType:         "cash,card",
				MaxPlayers:          9,
				Payment:             registrator.Payment_PAYMENT_CASH,
				Registered:          true,
				My:                  true,
				NumberOfMyLegioners: 3,
				NumberOfLegioners:   6,
				NumberOfPlayers:     9,
			},
		}, got)
		assert.NoError(t, err)
	})
}

func TestRegistrator_GetGames(t *testing.T) {
	// TODO tests
}

func TestRegistrator_GetPlayersByGameID(t *testing.T) {
	t.Run("error \"game not found\" while getting player list", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetPlayersByGameID(fx.ctx, int32(1)).Return(nil, model.ErrGameNotFound)

		got, err := fx.registrator.GetPlayersByGameID(fx.ctx, &registrator.GetPlayersByGameIDRequest{
			GameId: 1,
		})
		assert.Nil(t, got)

		st := status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("internal error while getting player list", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetPlayersByGameID(fx.ctx, int32(1)).Return(nil, errors.New("some error"))

		got, err := fx.registrator.GetPlayersByGameID(fx.ctx, &registrator.GetPlayersByGameIDRequest{
			GameId: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)
	})

	t.Run("empty players list", func(t *testing.T) {
		fx := tearUp(t)

		players := []model.GamePlayer{}

		fx.gamesFacade.EXPECT().GetPlayersByGameID(fx.ctx, int32(1)).Return(players, nil)

		got, err := fx.registrator.GetPlayersByGameID(fx.ctx, &registrator.GetPlayersByGameIDRequest{
			GameId: 1,
		})
		assert.ElementsMatch(t, []*registrator.Player{}, got.GetPlayers())
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		players := []model.GamePlayer{
			{
				ID:           1,
				FkGameID:     1,
				FkUserID:     1,
				RegisteredBy: 1,
				Degree:       int32(registrator.Degree_DEGREE_LIKELY),
			},
			{
				ID:           2,
				FkGameID:     1,
				FkUserID:     2,
				RegisteredBy: 2,
				Degree:       int32(registrator.Degree_DEGREE_UNLIKELY),
			},
		}

		fx.gamesFacade.EXPECT().GetPlayersByGameID(fx.ctx, int32(1)).Return(players, nil)

		got, err := fx.registrator.GetPlayersByGameID(fx.ctx, &registrator.GetPlayersByGameIDRequest{
			GameId: 1,
		})
		assert.ElementsMatch(t, []*registrator.Player{
			&registrator.Player{
				UserId:       1,
				RegisteredBy: 1,
				Degree:       registrator.Degree_DEGREE_LIKELY,
			},
			&registrator.Player{
				UserId:       2,
				RegisteredBy: 2,
				Degree:       registrator.Degree_DEGREE_UNLIKELY,
			},
		}, got.GetPlayers())
		assert.NoError(t, err)
	})
}

func TestRegistrator_GetRegisteredGames(t *testing.T) {
	// TODO tests
}

func TestRegistrator_GetUserGames(t *testing.T) {
	// TODO tests
}

func Test_convertModelGameToPBGame(t *testing.T) {
	timeNow := time_utils.TimeNow()
	t.Run("ok", func(t *testing.T) {
		game := model.Game{
			ID:          1,
			ExternalID:  2,
			LeagueID:    pkgmodel.LeagueQuizPlease,
			Type:        1,
			Number:      "1",
			Name:        "name",
			PlaceID:     4,
			Date:        model.DateTime(timeNow),
			Price:       400,
			PaymentType: "cash,card",
			MaxPlayers:  9,
			Payment:     1,
			Registered:  true,
		}

		game.My = true
		game.NumberOfMyLegioners = 3
		game.NumberOfLegioners = 6
		game.NumberOfPlayers = 8
		game.ResultPlace = 1

		got := convertModelGameToPBGame(game)
		assert.Equal(t, &registrator.Game{
			Id:                  1,
			ExternalId:          2,
			LeagueId:            pkgmodel.LeagueQuizPlease,
			Type:                1,
			Number:              "1",
			Name:                "name",
			PlaceId:             4,
			Date:                timestamppb.New(timeNow),
			Price:               400,
			PaymentType:         "cash,card",
			MaxPlayers:          9,
			Payment:             1,
			Registered:          true,
			My:                  true,
			NumberOfMyLegioners: 3,
			NumberOfLegioners:   6,
			NumberOfPlayers:     8,
			ResultPlace:         1,
		}, got)
	})
}
