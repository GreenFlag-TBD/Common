package broker

import (
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	ch *amqp.Channel
	*amqp.Queue
}

func (q *Publisher) Publish(msg amqp.Publishing) error {
	return q.ch.Publish("", q.Name, false, false, msg)
}

func (q *Publisher) PublishWithTimeout(msg amqp.Publishing, timeOut time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()
	return q.ch.PublishWithContext(ctx, "", q.Name, false, false, msg)
}
