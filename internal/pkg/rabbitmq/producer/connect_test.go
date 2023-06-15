package rabbitmq

import (
	"errors"
	"testing"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
)

func TestProducer_Connect(t *testing.T) {
	t.Run("queue declare error ", func(t *testing.T) {
		fx := tearUp(t)

		fx.producer.queueName = ""

		err := fx.producer.Connect(fx.ctx)
		assert.Error(t, err)
	})

	t.Run("queue declare error ", func(t *testing.T) {
		fx := tearUp(t)

		fx.rabbitMQChannel.EXPECT().QueueDeclare("queue_name",
			false,
			false,
			false,
			false,
			amqp.Table(nil)).Return(amqp.Queue{}, errors.New("some error"))

		err := fx.producer.Connect(fx.ctx)
		assert.Error(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.rabbitMQChannel.EXPECT().QueueDeclare("queue_name",
			false,
			false,
			false,
			false,
			amqp.Table(nil)).Return(amqp.Queue{}, nil)

		err := fx.producer.Connect(fx.ctx)
		assert.NoError(t, err)
	})
}
