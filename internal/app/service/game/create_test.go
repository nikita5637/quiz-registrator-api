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
	timeutils "github.com/nikita5637/quiz-registrator-api/utils/time"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestImplementation_CreateGame(t *testing.T) {
	t.Run("error: bad request", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.implementation.CreateGame(fx.ctx, &gamepb.CreateGameRequest{})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("error: validation error", func(t *testing.T) {
		fx := tearUp(t)

		paymentCertificate := gamepb.Payment_PAYMENT_CERTIFICATE
		got, err := fx.implementation.CreateGame(fx.ctx, &gamepb.CreateGameRequest{
			Game: &gamepb.Game{
				ExternalId: wrapperspb.Int32(1),
				LeagueId:   model.LeagueQuizPlease,
				Type:       gamepb.GameType_GAME_TYPE_CLASSIC,
				Number:     "1",
				Name:       wrapperspb.String(""), // invalid
				PlaceId:    1,
				Date: &timestamppb.Timestamp{
					Seconds: 1136214240,
				},
				Price:       400,
				PaymentType: wrapperspb.String("cash"),
				MaxPlayers:  9,
				Payment:     &paymentCertificate,
				IsInMaster:  true,
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)

		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, "INVALID_GAME_NAME", errorInfo.Reason)
		assert.Equal(t, map[string]string{
			"error": "Name: cannot be blank.",
		}, errorInfo.Metadata)
	})

	t.Run("error: game already exists", func(t *testing.T) {
		fx := tearUp(t)

		paymentCertificate := gamepb.Payment_PAYMENT_CERTIFICATE

		fx.gamesFacade.EXPECT().CreateGame(fx.ctx, model.Game{
			ExternalID:  maybe.Just(int32(1)),
			LeagueID:    math.MaxInt32,
			Type:        model.GameTypeClassic,
			Number:      "1",
			Name:        maybe.Just("name"),
			PlaceID:     1,
			Date:        model.DateTime(timeutils.ConvertTime("2006-01-02 15:04")),
			Price:       400,
			PaymentType: maybe.Just("cash"),
			MaxPlayers:  9,
			Payment:     maybe.Just(model.PaymentCertificate),
			IsInMaster:  true,
			GameLink:    maybe.Just("link"),
		}).Return(model.Game{}, games.ErrGameAlreadyExists)

		got, err := fx.implementation.CreateGame(fx.ctx, &gamepb.CreateGameRequest{
			Game: &gamepb.Game{
				ExternalId: wrapperspb.Int32(1),
				LeagueId:   math.MaxInt32,
				Type:       gamepb.GameType_GAME_TYPE_CLASSIC,
				Number:     "1",
				Name:       wrapperspb.String("name"),
				PlaceId:    1,
				Date: &timestamppb.Timestamp{
					Seconds: 1136214240,
				},
				Price:       400,
				PaymentType: wrapperspb.String("cash"),
				MaxPlayers:  9,
				Payment:     &paymentCertificate,
				IsInMaster:  true,
				GameLink:    wrapperspb.String("link"),
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

		paymentCertificate := gamepb.Payment_PAYMENT_CERTIFICATE

		fx.gamesFacade.EXPECT().CreateGame(fx.ctx, model.Game{
			ExternalID:  maybe.Just(int32(1)),
			LeagueID:    math.MaxInt32,
			Type:        model.GameTypeClassic,
			Number:      "1",
			Name:        maybe.Just("name"),
			PlaceID:     1,
			Date:        model.DateTime(timeutils.ConvertTime("2006-01-02 15:04")),
			Price:       400,
			PaymentType: maybe.Just("cash"),
			MaxPlayers:  9,
			Payment:     maybe.Just(model.PaymentCertificate),
			IsInMaster:  true,
			GameLink:    maybe.Just("link"),
		}).Return(model.Game{}, leagues.ErrLeagueNotFound)

		got, err := fx.implementation.CreateGame(fx.ctx, &gamepb.CreateGameRequest{
			Game: &gamepb.Game{
				ExternalId: wrapperspb.Int32(1),
				LeagueId:   math.MaxInt32,
				Type:       gamepb.GameType_GAME_TYPE_CLASSIC,
				Number:     "1",
				Name:       wrapperspb.String("name"),
				PlaceId:    1,
				Date: &timestamppb.Timestamp{
					Seconds: 1136214240,
				},
				Price:       400,
				PaymentType: wrapperspb.String("cash"),
				MaxPlayers:  9,
				Payment:     &paymentCertificate,
				IsInMaster:  true,
				GameLink:    wrapperspb.String("link"),
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

		paymentCertificate := gamepb.Payment_PAYMENT_CERTIFICATE

		fx.gamesFacade.EXPECT().CreateGame(fx.ctx, model.Game{
			ExternalID:  maybe.Just(int32(1)),
			LeagueID:    1,
			Type:        model.GameTypeClassic,
			Number:      "1",
			Name:        maybe.Just("name"),
			PlaceID:     1,
			Date:        model.DateTime(timeutils.ConvertTime("2006-01-02 15:04")),
			Price:       400,
			PaymentType: maybe.Just("cash"),
			MaxPlayers:  9,
			Payment:     maybe.Just(model.PaymentCertificate),
			IsInMaster:  true,
			GameLink:    maybe.Just("link"),
		}).Return(model.Game{}, places.ErrPlaceNotFound)

		got, err := fx.implementation.CreateGame(fx.ctx, &gamepb.CreateGameRequest{
			Game: &gamepb.Game{
				ExternalId: wrapperspb.Int32(1),
				LeagueId:   model.LeagueQuizPlease,
				Type:       gamepb.GameType_GAME_TYPE_CLASSIC,
				Number:     "1",
				Name:       wrapperspb.String("name"),
				PlaceId:    1,
				Date: &timestamppb.Timestamp{
					Seconds: 1136214240,
				},
				Price:       400,
				PaymentType: wrapperspb.String("cash"),
				MaxPlayers:  9,
				Payment:     &paymentCertificate,
				IsInMaster:  true,
				GameLink:    wrapperspb.String("link"),
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

	t.Run("error: internal error", func(t *testing.T) {
		fx := tearUp(t)

		paymentCertificate := gamepb.Payment_PAYMENT_CERTIFICATE

		fx.gamesFacade.EXPECT().CreateGame(fx.ctx, model.Game{
			ExternalID:  maybe.Just(int32(1)),
			LeagueID:    1,
			Type:        model.GameTypeClassic,
			Number:      "1",
			Name:        maybe.Just("name"),
			PlaceID:     1,
			Date:        model.DateTime(timeutils.ConvertTime("2006-01-02 15:04")),
			Price:       400,
			PaymentType: maybe.Just("cash"),
			MaxPlayers:  9,
			Payment:     maybe.Just(model.PaymentCertificate),
			IsInMaster:  true,
			GameLink:    maybe.Just("link"),
		}).Return(model.Game{}, errors.New("some error"))

		got, err := fx.implementation.CreateGame(fx.ctx, &gamepb.CreateGameRequest{
			Game: &gamepb.Game{
				ExternalId: wrapperspb.Int32(1),
				LeagueId:   model.LeagueQuizPlease,
				Type:       gamepb.GameType_GAME_TYPE_CLASSIC,
				Number:     "1",
				Name:       wrapperspb.String("name"),
				PlaceId:    1,
				Date: &timestamppb.Timestamp{
					Seconds: 1136214240,
				},
				Price:       400,
				PaymentType: wrapperspb.String("cash"),
				MaxPlayers:  9,
				Payment:     &paymentCertificate,
				IsInMaster:  true,
				GameLink:    wrapperspb.String("link"),
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

		paymentCertificate := gamepb.Payment_PAYMENT_CERTIFICATE

		fx.gamesFacade.EXPECT().CreateGame(fx.ctx, model.Game{
			ExternalID:  maybe.Just(int32(1)),
			LeagueID:    1,
			Type:        model.GameTypeClassic,
			Number:      "1",
			Name:        maybe.Just("name"),
			PlaceID:     1,
			Date:        model.DateTime(timeutils.ConvertTime("2006-01-02 15:04")),
			Price:       400,
			PaymentType: maybe.Just("cash"),
			MaxPlayers:  9,
			Payment:     maybe.Just(model.PaymentCertificate),
			IsInMaster:  true,
			GameLink:    maybe.Just("link"),
		}).Return(model.Game{
			ID:          1,
			ExternalID:  maybe.Just(int32(1)),
			LeagueID:    1,
			Type:        model.GameTypeClassic,
			Number:      "1",
			Name:        maybe.Just("name"),
			PlaceID:     1,
			Date:        model.DateTime(timeutils.ConvertTime("2006-01-02 15:04")),
			Price:       400,
			PaymentType: maybe.Just("cash"),
			MaxPlayers:  9,
			Payment:     maybe.Just(model.PaymentCertificate),
			IsInMaster:  true,
			GameLink:    maybe.Just("link"),
		}, nil)

		got, err := fx.implementation.CreateGame(fx.ctx, &gamepb.CreateGameRequest{
			Game: &gamepb.Game{
				ExternalId: wrapperspb.Int32(1),
				LeagueId:   model.LeagueQuizPlease,
				Type:       gamepb.GameType_GAME_TYPE_CLASSIC,
				Number:     "1",
				Name:       wrapperspb.String("name"),
				PlaceId:    1,
				Date: &timestamppb.Timestamp{
					Seconds: 1136214240,
				},
				Price:       400,
				PaymentType: wrapperspb.String("cash"),
				MaxPlayers:  9,
				Payment:     &paymentCertificate,
				IsInMaster:  true,
				GameLink:    wrapperspb.String("link"),
			},
		})
		assert.Equal(t, &gamepb.Game{
			Id:         1,
			ExternalId: wrapperspb.Int32(1),
			LeagueId:   model.LeagueQuizPlease,
			Type:       gamepb.GameType_GAME_TYPE_CLASSIC,
			Number:     "1",
			Name:       wrapperspb.String("name"),
			PlaceId:    1,
			Date: &timestamppb.Timestamp{
				Seconds: 1136214240,
			},
			Price:       400,
			PaymentType: wrapperspb.String("cash"),
			MaxPlayers:  9,
			Payment:     &paymentCertificate,
			IsInMaster:  true,
			GameLink:    wrapperspb.String("link"),
		}, got)
		assert.NoError(t, err)
	})
}

func Test_validateCreatedGame(t *testing.T) {
	type args struct {
		game model.Game
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "error: invalid ExternalID",
			args: args{
				game: model.Game{
					ID:          1,
					ExternalID:  maybe.Just(int32(-1)),
					LeagueID:    1,
					Type:        model.GameTypeClassic,
					Number:      "1",
					Name:        maybe.Just("name"),
					PlaceID:     1,
					Date:        model.DateTime(timeutils.ConvertTime("2006-01-02 15:04")),
					Price:       400,
					PaymentType: maybe.Just("cash"),
					MaxPlayers:  9,
					Payment:     maybe.Just(model.PaymentCertificate),
					IsInMaster:  true,
					Registered:  true,
				},
			},
			wantErr: true,
		},
		{
			name: "error: invalid ExternalID",
			args: args{
				game: model.Game{
					ID:          1,
					ExternalID:  maybe.Just(int32(0)),
					LeagueID:    1,
					Type:        model.GameTypeClassic,
					Number:      "1",
					Name:        maybe.Just("name"),
					PlaceID:     1,
					Date:        model.DateTime(timeutils.ConvertTime("2006-01-02 15:04")),
					Price:       400,
					PaymentType: maybe.Just("cash"),
					MaxPlayers:  9,
					Payment:     maybe.Just(model.PaymentCertificate),
					IsInMaster:  true,
					Registered:  true,
				},
			},
			wantErr: true,
		},
		{
			name: "error: invalid LeagueID",
			args: args{
				game: model.Game{
					ID:          1,
					ExternalID:  maybe.Just(int32(1)),
					LeagueID:    -1,
					Type:        model.GameTypeClassic,
					Number:      "1",
					Name:        maybe.Just("name"),
					PlaceID:     1,
					Date:        model.DateTime(timeutils.ConvertTime("2006-01-02 15:04")),
					Price:       400,
					PaymentType: maybe.Just("cash"),
					MaxPlayers:  9,
					Payment:     maybe.Just(model.PaymentCertificate),
					IsInMaster:  true,
					Registered:  true,
				},
			},
			wantErr: true,
		},
		{
			name: "error: invalid LeagueID",
			args: args{
				game: model.Game{
					ID:          1,
					ExternalID:  maybe.Just(int32(1)),
					LeagueID:    0,
					Type:        model.GameTypeClassic,
					Number:      "1",
					Name:        maybe.Just("name"),
					PlaceID:     1,
					Date:        model.DateTime(timeutils.ConvertTime("2006-01-02 15:04")),
					Price:       400,
					PaymentType: maybe.Just("cash"),
					MaxPlayers:  9,
					Payment:     maybe.Just(model.PaymentCertificate),
					IsInMaster:  true,
					Registered:  true,
				},
			},
			wantErr: true,
		},
		{
			name: "error: invalid Type",
			args: args{
				game: model.Game{
					ID:          1,
					ExternalID:  maybe.Just(int32(1)),
					LeagueID:    1,
					Type:        -1,
					Number:      "1",
					Name:        maybe.Just("name"),
					PlaceID:     1,
					Date:        model.DateTime(timeutils.ConvertTime("2006-01-02 15:04")),
					Price:       400,
					PaymentType: maybe.Just("cash"),
					MaxPlayers:  9,
					Payment:     maybe.Just(model.PaymentCertificate),
					IsInMaster:  true,
					Registered:  true,
				},
			},
			wantErr: true,
		},
		{
			name: "error: invalid Type",
			args: args{
				game: model.Game{
					ID:          1,
					ExternalID:  maybe.Just(int32(1)),
					LeagueID:    1,
					Type:        0,
					Number:      "1",
					Name:        maybe.Just("name"),
					PlaceID:     1,
					Date:        model.DateTime(timeutils.ConvertTime("2006-01-02 15:04")),
					Price:       400,
					PaymentType: maybe.Just("cash"),
					MaxPlayers:  9,
					Payment:     maybe.Just(model.PaymentCertificate),
					IsInMaster:  true,
					Registered:  true,
				},
			},
			wantErr: true,
		},
		{
			name: "error: invalid Number",
			args: args{
				game: model.Game{
					ID:          1,
					ExternalID:  maybe.Just(int32(1)),
					LeagueID:    1,
					Type:        model.GameTypeClassic,
					Number:      "",
					Name:        maybe.Just("name"),
					PlaceID:     1,
					Date:        model.DateTime(timeutils.ConvertTime("2006-01-02 15:04")),
					Price:       400,
					PaymentType: maybe.Just("cash"),
					MaxPlayers:  9,
					Payment:     maybe.Just(model.PaymentCertificate),
					IsInMaster:  true,
					Registered:  true,
				},
			},
			wantErr: true,
		},
		{
			name: "error: invalid Number",
			args: args{
				game: model.Game{
					ID:          1,
					ExternalID:  maybe.Just(int32(1)),
					LeagueID:    1,
					Type:        model.GameTypeClassic,
					Number:      "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
					Name:        maybe.Just("name"),
					PlaceID:     1,
					Date:        model.DateTime(timeutils.ConvertTime("2006-01-02 15:04")),
					Price:       400,
					PaymentType: maybe.Just("cash"),
					MaxPlayers:  9,
					Payment:     maybe.Just(model.PaymentCertificate),
					IsInMaster:  true,
					Registered:  true,
				},
			},
			wantErr: true,
		},
		{
			name: "error: invalid Number",
			args: args{
				game: model.Game{
					ID:          1,
					ExternalID:  maybe.Just(int32(1)),
					LeagueID:    1,
					Type:        model.GameTypeClosed,
					Number:      "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
					Name:        maybe.Just("name"),
					PlaceID:     1,
					Date:        model.DateTime(timeutils.ConvertTime("2006-01-02 15:04")),
					Price:       400,
					PaymentType: maybe.Just("cash"),
					MaxPlayers:  9,
					Payment:     maybe.Just(model.PaymentCertificate),
					IsInMaster:  true,
					Registered:  true,
				},
			},
			wantErr: true,
		},
		{
			name: "error: invalid Name",
			args: args{
				game: model.Game{
					ID:          1,
					ExternalID:  maybe.Just(int32(1)),
					LeagueID:    1,
					Type:        model.GameTypeClassic,
					Number:      "1",
					Name:        maybe.Just(""),
					PlaceID:     1,
					Date:        model.DateTime(timeutils.ConvertTime("2006-01-02 15:04")),
					Price:       400,
					PaymentType: maybe.Just("cash"),
					MaxPlayers:  9,
					Payment:     maybe.Just(model.PaymentCertificate),
					IsInMaster:  true,
					Registered:  true,
				},
			},
			wantErr: true,
		},
		{
			name: "error: invalid Name",
			args: args{
				game: model.Game{
					ID:          1,
					ExternalID:  maybe.Just(int32(1)),
					LeagueID:    1,
					Type:        model.GameTypeClassic,
					Number:      "1",
					Name:        maybe.Just("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
					PlaceID:     1,
					Date:        model.DateTime(timeutils.ConvertTime("2006-01-02 15:04")),
					Price:       400,
					PaymentType: maybe.Just("cash"),
					MaxPlayers:  9,
					Payment:     maybe.Just(model.PaymentCertificate),
					IsInMaster:  true,
					Registered:  true,
				},
			},
			wantErr: true,
		},
		{
			name: "error: invalid PlaceID",
			args: args{
				game: model.Game{
					ID:          1,
					ExternalID:  maybe.Just(int32(1)),
					LeagueID:    1,
					Type:        model.GameTypeClassic,
					Number:      "1",
					Name:        maybe.Just("name"),
					PlaceID:     -1,
					Date:        model.DateTime(timeutils.ConvertTime("2006-01-02 15:04")),
					Price:       400,
					PaymentType: maybe.Just("cash"),
					MaxPlayers:  9,
					Payment:     maybe.Just(model.PaymentCertificate),
					IsInMaster:  true,
					Registered:  true,
				},
			},
			wantErr: true,
		},
		{
			name: "error: invalid PlaceID",
			args: args{
				game: model.Game{
					ID:          1,
					ExternalID:  maybe.Just(int32(1)),
					LeagueID:    1,
					Type:        model.GameTypeClassic,
					Number:      "1",
					Name:        maybe.Just("name"),
					PlaceID:     0,
					Date:        model.DateTime(timeutils.ConvertTime("2006-01-02 15:04")),
					Price:       400,
					PaymentType: maybe.Just("cash"),
					MaxPlayers:  9,
					Payment:     maybe.Just(model.PaymentCertificate),
					IsInMaster:  true,
					Registered:  true,
				},
			},
			wantErr: true,
		},
		{
			name: "error: invalid Date",
			args: args{
				game: model.Game{
					ID:          1,
					ExternalID:  maybe.Just(int32(1)),
					LeagueID:    1,
					Type:        model.GameTypeClassic,
					Number:      "1",
					Name:        maybe.Just("name"),
					PlaceID:     1,
					Date:        model.DateTime(time.Unix(0, 0).UTC()),
					Price:       400,
					PaymentType: maybe.Just("cash"),
					MaxPlayers:  9,
					Payment:     maybe.Just(model.PaymentCertificate),
					IsInMaster:  true,
					Registered:  true,
				},
			},
			wantErr: true,
		},
		{
			name: "error: invalid PaymentType",
			args: args{
				game: model.Game{
					ID:          1,
					ExternalID:  maybe.Just(int32(1)),
					LeagueID:    1,
					Type:        model.GameTypeClassic,
					Number:      "1",
					Name:        maybe.Just("name"),
					PlaceID:     1,
					Date:        model.DateTime(timeutils.ConvertTime("2006-01-02 15:04")),
					Price:       400,
					PaymentType: maybe.Just("invalid"),
					MaxPlayers:  9,
					Payment:     maybe.Just(model.PaymentCertificate),
					IsInMaster:  true,
					Registered:  true,
				},
			},
			wantErr: true,
		},
		{
			name: "error: invalid PaymentType",
			args: args{
				game: model.Game{
					ID:          1,
					ExternalID:  maybe.Just(int32(1)),
					LeagueID:    1,
					Type:        model.GameTypeClassic,
					Number:      "1",
					Name:        maybe.Just("name"),
					PlaceID:     1,
					Date:        model.DateTime(timeutils.ConvertTime("2006-01-02 15:04")),
					Price:       400,
					PaymentType: maybe.Just(""),
					MaxPlayers:  9,
					Payment:     maybe.Just(model.PaymentCertificate),
					IsInMaster:  true,
					Registered:  true,
				},
			},
			wantErr: true,
		},
		{
			name: "error: invalid MaxPlayers",
			args: args{
				game: model.Game{
					ID:          1,
					ExternalID:  maybe.Just(int32(1)),
					LeagueID:    1,
					Type:        model.GameTypeClassic,
					Number:      "1",
					Name:        maybe.Just("name"),
					PlaceID:     1,
					Date:        model.DateTime(timeutils.ConvertTime("2006-01-02 15:04")),
					Price:       400,
					PaymentType: maybe.Just("cash"),
					MaxPlayers:  0,
					Payment:     maybe.Just(model.PaymentCertificate),
					IsInMaster:  true,
					Registered:  true,
				},
			},
			wantErr: true,
		},
		{
			name: "error: invalid Payment",
			args: args{
				game: model.Game{
					ID:          1,
					ExternalID:  maybe.Just(int32(1)),
					LeagueID:    1,
					Type:        model.GameTypeClassic,
					Number:      "1",
					Name:        maybe.Just("name"),
					PlaceID:     1,
					Date:        model.DateTime(timeutils.ConvertTime("2006-01-02 15:04")),
					Price:       400,
					PaymentType: maybe.Just("cash"),
					MaxPlayers:  9,
					Payment:     maybe.Just(model.Payment(-1)),
					IsInMaster:  true,
					Registered:  true,
				},
			},
			wantErr: true,
		},
		{
			name: "error: invalid Payment",
			args: args{
				game: model.Game{
					ID:          1,
					ExternalID:  maybe.Just(int32(1)),
					LeagueID:    1,
					Type:        model.GameTypeClassic,
					Number:      "1",
					Name:        maybe.Just("name"),
					PlaceID:     1,
					Date:        model.DateTime(timeutils.ConvertTime("2006-01-02 15:04")),
					Price:       400,
					PaymentType: maybe.Just("cash"),
					MaxPlayers:  9,
					Payment:     maybe.Just(model.Payment(0)),
					IsInMaster:  true,
					Registered:  true,
				},
			},
			wantErr: true,
		},
		{
			name: "error: invalid Payment",
			args: args{
				game: model.Game{
					ID:          1,
					ExternalID:  maybe.Just(int32(1)),
					LeagueID:    1,
					Type:        model.GameTypeClassic,
					Number:      "1",
					Name:        maybe.Just("name"),
					PlaceID:     1,
					Date:        model.DateTime(timeutils.ConvertTime("2006-01-02 15:04")),
					Price:       400,
					PaymentType: maybe.Just("cash"),
					MaxPlayers:  9,
					Payment:     maybe.Just(model.Payment(math.MaxInt32)),
					IsInMaster:  true,
					Registered:  true,
				},
			},
			wantErr: true,
		},
		{
			name: "ok",
			args: args{
				game: model.Game{
					ID:          1,
					ExternalID:  maybe.Just(int32(1)),
					LeagueID:    1,
					Type:        model.GameTypeClosed,
					Number:      "",
					Name:        maybe.Just("name"),
					PlaceID:     1,
					Date:        model.DateTime(timeutils.ConvertTime("2006-01-02 15:04")),
					Price:       400,
					PaymentType: maybe.Just("cash"),
					MaxPlayers:  9,
					Payment:     maybe.Just(model.PaymentCertificate),
					IsInMaster:  true,
					Registered:  true,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateCreatedGame(tt.args.game); (err != nil) != tt.wantErr {
				t.Errorf("validateCreatedGame() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
