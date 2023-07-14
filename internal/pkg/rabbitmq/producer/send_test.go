package rabbitmq

import (
	"errors"
	"testing"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
)

func TestProducer_Send(t *testing.T) {
	t.Run("publish with context error", func(t *testing.T) {
		fx := tearUp(t)

		fx.rabbitMQChannel.EXPECT().PublishWithContext(fx.ctx, "",
			"queue_name",
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        []byte("\"{}\""),
			}).Return(errors.New("some errror"))

		err := fx.producer.Send(fx.ctx, "{}")
		assert.Error(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.rabbitMQChannel.EXPECT().PublishWithContext(fx.ctx, "",
			"queue_name",
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        []byte("\"{}\""),
			}).Return(nil)

		err := fx.producer.Send(fx.ctx, "{}")
		assert.NoError(t, err)
	})
}
