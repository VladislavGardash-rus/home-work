package rabbit_mq

import (
	"context"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/logger"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

type Manager struct {
	connection *amqp.Connection
}

func NewManager(address string) (*Manager, error) {
	connection, err := amqp.Dial(address)
	if err != nil {
		return nil, err
	}

	manager := new(Manager)
	manager.connection = connection
	return manager, nil
}

func (c *Manager) Send(message []byte, queueName, contentType string) error {
	ch, err := c.connection.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: contentType,
			Body:        message,
		})
	if err != nil {
		return err
	}

	logger.UseLogger().Info("Send message in rabbitMQ: " + string(message))

	return nil
}

func (c *Manager) Consume(ctx context.Context, queueName string, do func(message []byte) error) error {
	ch, err := c.connection.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			err := do(d.Body)
			if err != nil {
				logger.UseLogger().Error(err)
			}
		}
	}()

	<-ctx.Done()

	return nil
}

func (c *Manager) Disconnect() error {
	return c.connection.Close()
}
