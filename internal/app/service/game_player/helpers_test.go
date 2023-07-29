package gameplayer

import (
	"context"
	"reflect"
	"testing"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/app/service/game_player/mocks"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	gameplayerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game_player"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type fixture struct {
	ctx context.Context

	gamesFacade       *mocks.GamesFacade
	gamePlayersFacade *mocks.GamePlayersFacade

	implementation *Implementation
}

func tearUp(t *testing.T) *fixture {
	fx := &fixture{
		ctx: context.Background(),

		gamesFacade:       mocks.NewGamesFacade(t),
		gamePlayersFacade: mocks.NewGamePlayersFacade(t),
	}

	fx.implementation = New(Config{
		GamesFacade:       fx.gamesFacade,
		GamePlayersFacade: fx.gamePlayersFacade,
	})

	t.Cleanup(func() {})

	return fx
}

func Test_convertModelGamePlayerToProtoGamePlayer(t *testing.T) {
	type args struct {
		gamePlayer model.GamePlayer
	}
	tests := []struct {
		name string
		args args
		want *gameplayerpb.GamePlayer
	}{
		{
			name: "tc1",
			args: args{
				gamePlayer: model.GamePlayer{
					ID:           1,
					GameID:       1,
					UserID:       maybe.Nothing[int32](),
					RegisteredBy: 1,
					Degree:       model.DegreeLikely,
				},
			},
			want: &gameplayerpb.GamePlayer{
				Id:           1,
				GameId:       1,
				RegisteredBy: 1,
				Degree:       gameplayerpb.Degree_DEGREE_LIKELY,
			},
		},
		{
			name: "tc2",
			args: args{
				gamePlayer: model.GamePlayer{
					ID:           1,
					GameID:       1,
					UserID:       maybe.Just(int32(1)),
					RegisteredBy: 1,
					Degree:       model.DegreeLikely,
				},
			},
			want: &gameplayerpb.GamePlayer{
				Id:     1,
				GameId: 1,
				UserId: &wrapperspb.Int32Value{
					Value: 1,
				},
				RegisteredBy: 1,
				Degree:       gameplayerpb.Degree_DEGREE_LIKELY,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertModelGamePlayerToProtoGamePlayer(tt.args.gamePlayer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertModelGamePlayerToProtoGamePlayer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertProtoGamePlayerToModelGamePlayer(t *testing.T) {
	type args struct {
		gamePlayer *gameplayerpb.GamePlayer
	}
	tests := []struct {
		name string
		args args
		want model.GamePlayer
	}{
		{
			name: "tc1",
			args: args{
				gamePlayer: &gameplayerpb.GamePlayer{
					Id:           1,
					GameId:       1,
					RegisteredBy: 1,
					Degree:       gameplayerpb.Degree_DEGREE_LIKELY,
				},
			},
			want: model.GamePlayer{
				ID:           1,
				GameID:       1,
				UserID:       maybe.Nothing[int32](),
				RegisteredBy: 1,
				Degree:       model.DegreeLikely,
			},
		},
		{
			name: "tc2",
			args: args{
				gamePlayer: &gameplayerpb.GamePlayer{
					Id:     1,
					GameId: 1,
					UserId: &wrapperspb.Int32Value{
						Value: 1,
					},
					RegisteredBy: 1,
					Degree:       gameplayerpb.Degree_DEGREE_LIKELY,
				},
			},
			want: model.GamePlayer{
				ID:           1,
				GameID:       1,
				UserID:       maybe.Just(int32(1)),
				RegisteredBy: 1,
				Degree:       model.DegreeLikely,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertProtoGamePlayerToModelGamePlayer(tt.args.gamePlayer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertProtoGamePlayerToModelGamePlayer() = %v, want %v", got, tt.want)
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
			name: "GameID",
			args: args{
				keys: []string{"GameID"},
			},
			want: &errorDetails{
				Reason: invalidGameIDReason,
				Lexeme: invalidGameIDLexeme,
			},
		},
		{
			name: "UserID",
			args: args{
				keys: []string{"UserID"},
			},
			want: &errorDetails{
				Reason: invalidUserIDReason,
				Lexeme: invalidUserIDLexeme,
			},
		},
		{
			name: "RegisteredBy",
			args: args{
				keys: []string{"RegisteredBy"},
			},
			want: &errorDetails{
				Reason: invalidRegisteredByReason,
				Lexeme: invalidRegisteredByLexeme,
			},
		},
		{
			name: "Degree",
			args: args{
				keys: []string{"Degree"},
			},
			want: &errorDetails{
				Reason: invalidDegreeReason,
				Lexeme: invalidDegreeLexeme,
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

func Test_validateUserID(t *testing.T) {
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
				value: "not Maybe[int32]",
			},
			wantErr: true,
		},
		{
			name: "nothing",
			args: args{
				value: maybe.Nothing[int32](),
			},
			wantErr: false,
		},
		{
			name: "eq 0",
			args: args{
				value: maybe.Just(int32(0)),
			},
			wantErr: true,
		},
		{
			name: "ok",
			args: args{
				value: maybe.Just(int32(1)),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateUserID(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("validateUserID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
