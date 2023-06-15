package rabbitmq

import (
	"context"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/rabbitmq/producer/mocks"
)

type fixture struct {
	ctx context.Context

	producer        *Producer
	rabbitMQChannel *mocks.Channel
}

func tearUp(t *testing.T) *fixture {
	fx := &fixture{
		ctx: context.Background(),

		rabbitMQChannel: mocks.NewChannel(t),
	}

	fx.producer = New(Config{
		QueueName:       "queue_name",
		RabbitMQChannel: fx.rabbitMQChannel,
	})

	t.Cleanup(func() {})

	return fx
}
