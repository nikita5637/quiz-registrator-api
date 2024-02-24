package rabbitmq

import (
	"context"
	"errors"
	"fmt"
)

// Connect ...
func (p *Producer) Connect(_ context.Context) error {
	if p.queueName == "" {
		return errors.New("queue name is empty")
	}

	_, err := p.rabbitMQChannel.QueueDeclare(
		p.queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("queue declare error: %w", err)
	}

	return nil
}
