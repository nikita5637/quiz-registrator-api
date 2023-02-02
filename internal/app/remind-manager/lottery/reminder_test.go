//go:generate mockery --case underscore --name Croupier --with-expecter
//go:generate mockery --case underscore --name GamesFacade --with-expecter
//go:generate mockery --case underscore --name RabbitMQChannel --with-expecter

package game

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	time_utils "github.com/nikita5637/quiz-registrator-api/utils/time"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
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
				cfg: Config{
					QueueName: "queue",
				},
			},
			want: &Reminder{
				alreadyRemindedGames: map[int32]struct{}{},
				queueName:            "queue",
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
	t.Run("queue declare error", func(t *testing.T) {
		time_utils.TimeNow = func() time.Time {
			return time_utils.ConvertTime("2022-02-10 10:00")
		}

		fx := tearUp(t)

		fx.rabbitMQChannel.EXPECT().QueueDeclare("queue", false, false, false, false, amqp.Table(nil)).Return(amqp.Queue{}, errors.New("some error"))

		err := fx.reminder.Run(fx.ctx)
		assert.Error(t, err)
	})

	t.Run("get games with active lottery error", func(t *testing.T) {
		time_utils.TimeNow = func() time.Time {
			return time_utils.ConvertTime("2022-02-10 10:00")
		}

		fx := tearUp(t)

		fx.rabbitMQChannel.EXPECT().QueueDeclare("queue", false, false, false, false, amqp.Table(nil)).Return(amqp.Queue{}, nil)
		fx.croupier.EXPECT().GetGamesWithActiveLottery(fx.ctx).Return(nil, errors.New("some error"))

		err := fx.reminder.Run(fx.ctx)
		assert.Error(t, err)
	})

	t.Run("get players by game ID error", func(t *testing.T) {
		time_utils.TimeNow = func() time.Time {
			return time_utils.ConvertTime("2022-02-10 10:00")
		}

		fx := tearUp(t)

		fx.rabbitMQChannel.EXPECT().QueueDeclare("queue", false, false, false, false, amqp.Table(nil)).Return(amqp.Queue{}, nil)
		fx.croupier.EXPECT().GetGamesWithActiveLottery(fx.ctx).Return([]model.Game{
			{
				ID: 1,
			},
			{
				ID: 2,
			},
		}, nil)

		fx.gamesFacade.EXPECT().GetPlayersByGameID(fx.ctx, int32(1)).Return([]model.GamePlayer{
			{
				FkUserID: 1,
			},
			{
				FkUserID: 2,
			},
		}, nil)

		fx.gamesFacade.EXPECT().GetPlayersByGameID(fx.ctx, int32(2)).Return([]model.GamePlayer{}, errors.New("some error"))

		fx.rabbitMQChannel.EXPECT().PublishWithContext(fx.ctx, "", "queue", false, false, amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(`{"game_id":1,"player_ids":[1,2]}`),
		}).Return(nil)

		err := fx.reminder.Run(fx.ctx)
		assert.NoError(t, err)
	})

	t.Run("there are not players to remind", func(t *testing.T) {
		time_utils.TimeNow = func() time.Time {
			return time_utils.ConvertTime("2022-02-10 10:00")
		}

		fx := tearUp(t)

		fx.rabbitMQChannel.EXPECT().QueueDeclare("queue", false, false, false, false, amqp.Table(nil)).Return(amqp.Queue{}, nil)
		fx.croupier.EXPECT().GetGamesWithActiveLottery(fx.ctx).Return([]model.Game{
			{
				ID: 1,
			},
		}, nil)

		fx.gamesFacade.EXPECT().GetPlayersByGameID(fx.ctx, int32(1)).Return([]model.GamePlayer{
			{
				ID:           1,
				RegisteredBy: 1,
			},
			{
				ID:           2,
				RegisteredBy: 1,
			},
		}, nil)

		err := fx.reminder.Run(fx.ctx)
		assert.NoError(t, err)
	})

	t.Run("publish with context error", func(t *testing.T) {
		time_utils.TimeNow = func() time.Time {
			return time_utils.ConvertTime("2022-02-10 10:00")
		}

		fx := tearUp(t)

		fx.rabbitMQChannel.EXPECT().QueueDeclare("queue", false, false, false, false, amqp.Table(nil)).Return(amqp.Queue{}, nil)
		fx.croupier.EXPECT().GetGamesWithActiveLottery(fx.ctx).Return([]model.Game{
			{
				ID: 1,
			},
			{
				ID: 2,
			},
		}, nil)

		fx.gamesFacade.EXPECT().GetPlayersByGameID(fx.ctx, int32(1)).Return([]model.GamePlayer{
			{
				FkUserID: 1,
			},
			{
				FkUserID: 2,
			},
		}, nil)

		fx.gamesFacade.EXPECT().GetPlayersByGameID(fx.ctx, int32(2)).Return([]model.GamePlayer{
			{
				FkUserID: 3,
			},
			{
				FkUserID: 4,
			},
		}, nil)

		fx.rabbitMQChannel.EXPECT().PublishWithContext(fx.ctx, "", "queue", false, false, amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(`{"game_id":1,"player_ids":[1,2]}`),
		}).Return(nil)

		fx.rabbitMQChannel.EXPECT().PublishWithContext(fx.ctx, "", "queue", false, false, amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(`{"game_id":2,"player_ids":[3,4]}`),
		}).Return(errors.New("some error"))

		err := fx.reminder.Run(fx.ctx)
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		time_utils.TimeNow = func() time.Time {
			return time_utils.ConvertTime("2022-02-10 10:00")
		}

		fx := tearUp(t)

		fx.rabbitMQChannel.EXPECT().QueueDeclare("queue", false, false, false, false, amqp.Table(nil)).Return(amqp.Queue{}, nil)
		fx.croupier.EXPECT().GetGamesWithActiveLottery(fx.ctx).Return([]model.Game{
			{
				ID: 1,
			},
			{
				ID: 2,
			},
		}, nil)

		fx.gamesFacade.EXPECT().GetPlayersByGameID(fx.ctx, int32(1)).Return([]model.GamePlayer{
			{
				FkUserID: 1,
			},
			{
				FkUserID: 2,
			},
		}, nil)

		fx.gamesFacade.EXPECT().GetPlayersByGameID(fx.ctx, int32(2)).Return([]model.GamePlayer{
			{
				FkUserID: 3,
			},
			{
				FkUserID: 4,
			},
		}, nil)

		fx.rabbitMQChannel.EXPECT().PublishWithContext(fx.ctx, "", "queue", false, false, amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(`{"game_id":1,"player_ids":[1,2]}`),
		}).Return(nil)

		fx.rabbitMQChannel.EXPECT().PublishWithContext(fx.ctx, "", "queue", false, false, amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(`{"game_id":2,"player_ids":[3,4]}`),
		}).Return(nil)

		err := fx.reminder.Run(fx.ctx)
		assert.NoError(t, err)
	})
}
