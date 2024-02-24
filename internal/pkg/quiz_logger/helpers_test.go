package quizlogger

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mocks"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
)

type fixture struct {
	ctx context.Context

	logStorage *mocks.LogStorage

	logger *Logger
}

func tearUp(t *testing.T) *fixture {
	fx := &fixture{
		ctx: context.Background(),

		logStorage: mocks.NewLogStorage(t),
	}

	fx.logger = New(Config{
		LogStorage: fx.logStorage,
	})

	t.Cleanup(func() {})

	return fx
}
func Test_convertParamsToDBLog(t *testing.T) {
	type args struct {
		params Params
	}
	tests := []struct {
		name string
		args args
		want database.Log
	}{
		{
			name: "tc1",
			args: args{
				params: Params{
					UserID:     maybe.Just(int32(1)),
					ActionID:   2,
					MessageID:  3,
					ObjectType: maybe.Just(ObjectTypeGame),
					ObjectID:   maybe.Just(int32(777)),
					Metadata: GamePaymentChangedMetadata{
						OldPayment: model.PaymentCash,
						NewPayment: model.PaymentCertificate,
					},
				},
			},
			want: database.Log{
				UserID: sql.NullInt64{
					Int64: 1,
					Valid: true,
				},
				ActionID:  2,
				MessageID: 3,
				ObjectType: sql.NullString{
					String: "game",
					Valid:  true,
				},
				ObjectID: sql.NullInt64{
					Int64: 777,
					Valid: true,
				},
				Metadata: sql.NullString{
					String: "{\"OldPayment\":1,\"NewPayment\":2}",
					Valid:  true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertParamsToDBLog(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertParamsToDBLog() = %v, want %v", got, tt.want)
			}
		})
	}
}
