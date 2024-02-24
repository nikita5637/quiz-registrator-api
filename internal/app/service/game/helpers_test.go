package game

import (
	"context"
	"math"
	"reflect"
	"testing"
	"time"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/app/service/game/mocks"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	gamepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game"
	timeutils "github.com/nikita5637/quiz-registrator-api/utils/time"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type fixture struct {
	ctx              context.Context
	gamesFacade      *mocks.GamesFacade
	rabbitMQProducer *mocks.RabbitMQProducer

	implementation *Implementation
}

func tearUp(t *testing.T) *fixture {
	fx := &fixture{
		ctx:              context.Background(),
		gamesFacade:      mocks.NewGamesFacade(t),
		rabbitMQProducer: mocks.NewRabbitMQProducer(t),
	}

	fx.implementation = New(Config{
		GamesFacade:      fx.gamesFacade,
		RabbitMQProducer: fx.rabbitMQProducer,
	})

	t.Cleanup(func() {})

	return fx
}

func Test_convertModelGameToProtoGame(t *testing.T) {
	paymentCertificate := gamepb.Payment_PAYMENT_CERTIFICATE

	type args struct {
		game model.Game
	}
	tests := []struct {
		name string
		args args
		want *gamepb.Game
	}{
		{
			name: "tc1",
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
					Payment:     maybe.Just(model.PaymentCertificate),
					Registered:  true,
					IsInMaster:  true,
					HasPassed:   true,
					GameLink:    maybe.Just("link"),
				},
			},
			want: &gamepb.Game{
				Id:         1,
				ExternalId: wrapperspb.Int32(1),
				LeagueId:   1,
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
				Registered:  true,
				IsInMaster:  true,
				HasPassed:   true,
				GameLink:    wrapperspb.String("link"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertModelGameToProtoGame(tt.args.game); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertModelGameToProtoGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertProtoGameToModelGame(t *testing.T) {
	paymentCertificate := gamepb.Payment_PAYMENT_CERTIFICATE

	type args struct {
		game *gamepb.Game
	}
	tests := []struct {
		name string
		args args
		want model.Game
	}{
		{
			name: "tc1",
			args: args{
				game: &gamepb.Game{
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
					Registered:  true,
					IsInMaster:  true,
					HasPassed:   true,
					GameLink:    wrapperspb.String("link"),
				},
			},
			want: model.Game{
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
				Registered:  true,
				IsInMaster:  true,
				HasPassed:   true,
				GameLink:    maybe.Just("link"),
			},
		},
		{
			name: "tc2",
			args: args{
				game: &gamepb.Game{
					Id:          1,
					ExternalId:  wrapperspb.Int32(1),
					LeagueId:    model.LeagueQuizPlease,
					Type:        gamepb.GameType_GAME_TYPE_CLASSIC,
					Number:      "1",
					Name:        wrapperspb.String("name"),
					PlaceId:     1,
					Date:        nil,
					Price:       400,
					PaymentType: wrapperspb.String("cash"),
					MaxPlayers:  9,
					Payment:     &paymentCertificate,
					Registered:  true,
					IsInMaster:  true,
					HasPassed:   true,
					GameLink:    wrapperspb.String("link"),
				},
			},
			want: model.Game{
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
				Registered:  true,
				IsInMaster:  true,
				HasPassed:   true,
				GameLink:    maybe.Just("link"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertProtoGameToModelGame(tt.args.game); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertProtoGameToModelGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getErrorDetails(t *testing.T) {
	type args struct {
		keys []string
	}
	tests := []struct {
		name string
		args args
		want *errorDetails
	}{
		{
			name: "keys is nil",
			args: args{
				keys: nil,
			},
			want: nil,
		},
		{
			name: "keys is empty",
			args: args{
				keys: []string{},
			},
			want: nil,
		},
		{
			name: "Date",
			args: args{
				keys: []string{
					"Date",
				},
			},
			want: &errorDetails{
				Reason: reasonInvalidGameDate,
				Lexeme: invalidGameDateLexeme,
			},
		},
		{
			name: "ExternalID",
			args: args{
				keys: []string{
					"ExternalID",
				},
			},
			want: &errorDetails{
				Reason: reasonInvalidExternalID,
				Lexeme: invalidExternalIDLexeme,
			},
		},
		{
			name: "LeagueID",
			args: args{
				keys: []string{
					"LeagueID",
				},
			},
			want: &errorDetails{
				Reason: reasonInvalidLeagueID,
				Lexeme: invalidLeagueIDLexeme,
			},
		},
		{
			name: "MaxPlayers",
			args: args{
				keys: []string{
					"MaxPlayers",
				},
			},
			want: &errorDetails{
				Reason: reasonInvalidMaxPlayers,
				Lexeme: invalidMaxPlayersLexeme,
			},
		},
		{
			name: "Name",
			args: args{
				keys: []string{
					"Name",
				},
			},
			want: &errorDetails{
				Reason: reasonInvalidGameName,
				Lexeme: invalidGameNameLexeme,
			},
		},
		{
			name: "Number",
			args: args{
				keys: []string{
					"Number",
				},
			},
			want: &errorDetails{
				Reason: reasonInvalidGameNumber,
				Lexeme: invalidGameNumberLexeme,
			},
		},
		{
			name: "Payment",
			args: args{
				keys: []string{
					"Payment",
				},
			},
			want: &errorDetails{
				Reason: reasonInvalidPayment,
				Lexeme: invalidPaymentLexeme,
			},
		},
		{
			name: "PaymentType",
			args: args{
				keys: []string{
					"PaymentType",
				},
			},
			want: &errorDetails{
				Reason: reasonInvalidPaymentType,
				Lexeme: invalidPaymentTypeLexeme,
			},
		},
		{
			name: "PlaceID",
			args: args{
				keys: []string{
					"PlaceID",
				},
			},
			want: &errorDetails{
				Reason: reasonInvalidPlaceID,
				Lexeme: invalidPlaceIDLexeme,
			},
		},
		{
			name: "Type",
			args: args{
				keys: []string{
					"Type",
				},
			},
			want: &errorDetails{
				Reason: reasonInvalidGameType,
				Lexeme: invalidGameTypeLexeme,
			},
		},
		{
			name: "not found",
			args: args{
				keys: []string{"not found"},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getErrorDetails(tt.args.keys); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getErrorDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateExternalID(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "not Maybe[int32]",
			args: args{
				value: 1,
			},
			wantErr: true,
		},
		{
			name: "is not present",
			args: args{
				value: maybe.Nothing[int32](),
			},
			wantErr: false,
		},
		{
			name: "lt 0",
			args: args{
				value: maybe.Just(int32(-1)),
			},
			wantErr: true,
		},
		{
			name: "eq 0",
			args: args{
				value: maybe.Just(int32(0)),
			},
			wantErr: true,
		},
		{
			name: "gt 0",
			args: args{
				value: maybe.Just(int32(1)),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateExternalID(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("validateExternalID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validateName(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "is not Maybe[string]",
			args: args{
				value: "not Maybe[string]",
			},
			wantErr: true,
		},
		{
			name: "not present",
			args: args{
				value: maybe.Nothing[string](),
			},
			wantErr: false,
		},
		{
			name: "empty",
			args: args{
				value: maybe.Just(""),
			},
			wantErr: true,
		},
		{
			name: "len lt 128",
			args: args{
				value: maybe.Just("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
			},
			wantErr: true,
		},
		{
			name: "ok",
			args: args{
				value: maybe.Just("name"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateName(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("validateName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validatePayment(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "is not Maybe[model.Payment]",
			args: args{
				value: "not Maybe[model.Payment]",
			},
			wantErr: true,
		},
		{
			name: "not present",
			args: args{
				value: maybe.Nothing[model.Payment](),
			},
			wantErr: false,
		},
		{
			name: "lt 0",
			args: args{
				value: maybe.Just(model.Payment(-1)),
			},
			wantErr: true,
		},
		{
			name: "eq 0",
			args: args{
				value: maybe.Just(model.Payment(0)),
			},
			wantErr: true,
		},
		{
			name: "PaymentCash",
			args: args{
				value: maybe.Just(model.PaymentCash),
			},
			wantErr: false,
		},
		{
			name: "PaymentCertificate",
			args: args{
				value: maybe.Just(model.PaymentCertificate),
			},
			wantErr: false,
		},
		{
			name: "PaymentMixed",
			args: args{
				value: maybe.Just(model.PaymentMixed),
			},
			wantErr: false,
		},
		{
			name: "math.MaxInt32",
			args: args{
				value: maybe.Just(model.Payment(math.MaxInt32)),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validatePayment(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("validatePayment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validatePaymentType(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "is not Maybe[string]",
			args: args{
				value: "not Maybe[string]",
			},
			wantErr: true,
		},
		{
			name: "not present",
			args: args{
				value: maybe.Nothing[string](),
			},
			wantErr: false,
		},
		{
			name: "empty",
			args: args{
				value: maybe.Just(""),
			},
			wantErr: true,
		},
		{
			name: "invalid value",
			args: args{
				value: maybe.Just("invalid value"),
			},
			wantErr: true,
		},
		{
			name: "cash,card",
			args: args{
				value: maybe.Just("cash,card"),
			},
			wantErr: false,
		},
		{
			name: "cash",
			args: args{
				value: maybe.Just("cash"),
			},
			wantErr: false,
		},
		{
			name: "card",
			args: args{
				value: maybe.Just("card"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validatePaymentType(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("validatePaymentType() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validatePaymentType2(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "is not string",
			args: args{
				value: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validatePaymentType2(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("validatePaymentType2() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
