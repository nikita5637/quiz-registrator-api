package gameresultmanager

import (
	"context"
	"reflect"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/app/game_result_manager/mocks"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	gameresultmanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game_result_manager"
)

type fixture struct {
	ctx context.Context

	gameResultsFacade *mocks.GameResultsFacade

	gameResultManager *GameResultManager
}

func tearUp(t *testing.T) *fixture {
	fx := &fixture{
		ctx: context.Background(),

		gameResultsFacade: mocks.NewGameResultsFacade(t),
	}

	fx.gameResultManager = New(Config{
		GameResultsFacade: fx.gameResultsFacade,
	})

	t.Cleanup(func() {})

	return fx
}

func Test_convertModelGameResultToProtoGameResult(t *testing.T) {
	type args struct {
		gameResult model.GameResult
	}
	tests := []struct {
		name string
		args args
		want *gameresultmanagerpb.GameResult
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
			want: &gameresultmanagerpb.GameResult{
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
