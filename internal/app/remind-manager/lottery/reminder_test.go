package lottery

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/logger"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	leaguepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/league"
	"github.com/nikita5637/quiz-registrator-api/pkg/reminder"
	timeutils "github.com/nikita5637/quiz-registrator-api/utils/time"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestNew(t *testing.T) {
	type args struct {
		cfg Config
	}
	tests := []struct {
		name string
		args args
		want *Reminder
	}{
		{
			name: "test case 1",
			args: args{
				cfg: Config{},
			},
			want: &Reminder{
				alreadyRemindedGames: map[int32]struct{}{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReminder_Run(t *testing.T) {
	t.Run("get games with active lottery error", func(t *testing.T) {
		timeutils.TimeNow = func() time.Time {
			return timeutils.ConvertTime("2022-02-10 10:00")
		}

		fx := tearUp(t)

		ctx := logger.ToContext(fx.ctx, logger.FromContext(fx.ctx).WithOptions(zap.Fields(
			zap.String("reminder_name", "lottery reminder"),
		)))

		fx.croupier.EXPECT().GetGamesWithActiveLottery(ctx).Return(nil, errors.New("some error"))

		err := fx.reminder.Run(fx.ctx)
		assert.Error(t, err)
	})

	t.Run("get players by game ID error", func(t *testing.T) {
		timeutils.TimeNow = func() time.Time {
			return timeutils.ConvertTime("2022-02-10 10:00")
		}

		fx := tearUp(t)

		ctx := logger.ToContext(fx.ctx, logger.FromContext(fx.ctx).WithOptions(zap.Fields(
			zap.String("reminder_name", "lottery reminder"),
		)))

		fx.croupier.EXPECT().GetGamesWithActiveLottery(ctx).Return([]model.Game{
			{
				ID: 1,
			},
			{
				ID: 2,
			},
		}, nil)

		fx.gamePlayersFacade.EXPECT().GetGamePlayersByGameID(ctx, int32(1)).Return([]model.GamePlayer{
			{
				UserID: maybe.Just(int32(1)),
			},
			{
				UserID: maybe.Just(int32(2)),
			},
		}, nil)

		fx.gamePlayersFacade.EXPECT().GetGamePlayersByGameID(ctx, int32(2)).Return([]model.GamePlayer{}, errors.New("some error"))

		fx.rabbitMQProducer.EXPECT().Send(ctx, reminder.Lottery{
			GameID:    1,
			PlayerIDs: []int32{1, 2},
		}).Return(nil)

		err := fx.reminder.Run(fx.ctx)
		assert.NoError(t, err)
	})

	t.Run("there are not players to remind", func(t *testing.T) {
		timeutils.TimeNow = func() time.Time {
			return timeutils.ConvertTime("2022-02-10 10:00")
		}

		fx := tearUp(t)

		ctx := logger.ToContext(fx.ctx, logger.FromContext(fx.ctx).WithOptions(zap.Fields(
			zap.String("reminder_name", "lottery reminder"),
		)))

		fx.croupier.EXPECT().GetGamesWithActiveLottery(ctx).Return([]model.Game{
			{
				ID: 1,
			},
		}, nil)

		fx.gamePlayersFacade.EXPECT().GetGamePlayersByGameID(ctx, int32(1)).Return([]model.GamePlayer{
			{
				ID:           1,
				UserID:       maybe.Nothing[int32](),
				RegisteredBy: 1,
			},
			{
				ID:           2,
				UserID:       maybe.Nothing[int32](),
				RegisteredBy: 1,
			},
		}, nil)

		err := fx.reminder.Run(fx.ctx)
		assert.NoError(t, err)
	})

	t.Run("publish with context error", func(t *testing.T) {
		timeutils.TimeNow = func() time.Time {
			return timeutils.ConvertTime("2022-02-10 10:00")
		}

		fx := tearUp(t)

		ctx := logger.ToContext(fx.ctx, logger.FromContext(fx.ctx).WithOptions(zap.Fields(
			zap.String("reminder_name", "lottery reminder"),
		)))

		fx.croupier.EXPECT().GetGamesWithActiveLottery(ctx).Return([]model.Game{
			{
				ID: 1,
			},
			{
				ID: 2,
			},
		}, nil)

		fx.gamePlayersFacade.EXPECT().GetGamePlayersByGameID(ctx, int32(1)).Return([]model.GamePlayer{
			{
				UserID: maybe.Just(int32(1)),
			},
			{
				UserID: maybe.Just(int32(2)),
			},
		}, nil)

		fx.gamePlayersFacade.EXPECT().GetGamePlayersByGameID(ctx, int32(2)).Return([]model.GamePlayer{
			{
				UserID: maybe.Just(int32(3)),
			},
			{
				UserID: maybe.Just(int32(4)),
			},
		}, nil)

		fx.rabbitMQProducer.EXPECT().Send(ctx, reminder.Lottery{
			GameID:    1,
			PlayerIDs: []int32{1, 2},
		}).Return(nil)

		fx.rabbitMQProducer.EXPECT().Send(ctx, reminder.Lottery{
			GameID:    2,
			PlayerIDs: []int32{3, 4},
		}).Return(errors.New("some error"))

		err := fx.reminder.Run(fx.ctx)
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		timeutils.TimeNow = func() time.Time {
			return timeutils.ConvertTime("2022-02-10 10:00")
		}

		fx := tearUp(t)

		ctx := logger.ToContext(fx.ctx, logger.FromContext(fx.ctx).WithOptions(zap.Fields(
			zap.String("reminder_name", "lottery reminder"),
		)))

		fx.croupier.EXPECT().GetGamesWithActiveLottery(ctx).Return([]model.Game{
			{
				ID:       1,
				LeagueID: int32(leaguepb.LeagueID_QUIZ_PLEASE),
			},
			{
				ID:       2,
				LeagueID: int32(leaguepb.LeagueID_SQUIZ),
			},
		}, nil)

		fx.gamePlayersFacade.EXPECT().GetGamePlayersByGameID(ctx, int32(1)).Return([]model.GamePlayer{
			{
				UserID: maybe.Just(int32(1)),
			},
			{
				UserID: maybe.Just(int32(2)),
			},
		}, nil)

		fx.gamePlayersFacade.EXPECT().GetGamePlayersByGameID(ctx, int32(2)).Return([]model.GamePlayer{
			{
				UserID: maybe.Just(int32(3)),
			},
			{
				UserID: maybe.Just(int32(4)),
			},
		}, nil)

		fx.rabbitMQProducer.EXPECT().Send(ctx, reminder.Lottery{
			GameID:    1,
			LeagueID:  1,
			PlayerIDs: []int32{1, 2},
		}).Return(nil)

		fx.rabbitMQProducer.EXPECT().Send(ctx, reminder.Lottery{
			GameID:    2,
			LeagueID:  2,
			PlayerIDs: []int32{3, 4},
		}).Return(nil)

		err := fx.reminder.Run(fx.ctx)
		assert.NoError(t, err)
	})
}
