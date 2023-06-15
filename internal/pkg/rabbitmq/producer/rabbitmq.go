//go:generate mockery --case underscore --name Channel --with-expecter

package rabbitmq

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Channel ...
type Channel interface {
	PublishWithContext(ctx context.Context, exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error
	QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) (amqp.Queue, error)
}

// Producer ...
type Producer struct {
	queueName       string
	rabbitMQChannel Channel
}

// Config ...
type Config struct {
	QueueName       string
	RabbitMQChannel Channel
}

// New ...
func New(cfg Config) *Producer {
	return &Producer{
		queueName:       cfg.QueueName,
		rabbitMQChannel: cfg.RabbitMQChannel,
	}
}
