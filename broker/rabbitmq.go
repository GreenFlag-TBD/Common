package broker

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
)

type RabbitMQ struct {
	*amqp.Connection
	chs []*amqp.Channel
}

func NewRabbitMQ(dial string) (*RabbitMQ, error) {
	log.Info().Msg("Connecting to RabbitMQ")
	conn, err := amqp.Dial(dial)
	if err != nil {
		log.Error().Err(err).Msg("Failed to connect to RabbitMQ, exiting with error message:")
		return nil, err
	}
	log.Info().Msg("Connected to RabbitMQ")
	return &RabbitMQ{conn, make([]*amqp.Channel, 0)}, nil
}

func (r *RabbitMQ) Disconnect() error {
	log.Info().Msg("Disconnecting from RabbitMQ")
	for _, ch := range r.chs {
		log.Info().Msg("Closing channel to RabbitMQ")
		if ch != nil {
			ch.Close()
		}
	}
	return r.Close()
}

func (r *RabbitMQ) QueueDeclare(
	/// The name of the queue
	name string,
	/// Durable queues will survive server restarts
	durable bool,
	/// Auto-deleted queues will be deleted when there are no remaining consumers
	autoDelete bool,
	/// Exclusive queues are only accessible by the connection that created them
	exclusive bool,
	/// NoWait will not wait for a confirmation from the server
	noWait bool,
	/// Arguments for declaring the queue
	args amqp.Table,
) (*Publisher, error) {
	log.Info().Msg("Declaring queue")
	ch, err := r.Channel()
	if err != nil {
		log.Err(err).Msg("Failed to create channel, exiting with error message:")
		return nil, err
	}
	r.chs = append(r.chs, ch)
	q, err := ch.QueueDeclare(name, durable, autoDelete, exclusive, noWait, args)
	if err != nil {
		log.Err(err).Msg("Failed to declare queue, exiting with error message:")
		return nil, err
	}
	log.Info().Msg("Queue declared with name: " + name)
	return &Publisher{Queue: &q, ch: ch}, nil
}

func (r *RabbitMQ) ConsumerDeclare() (*Consumer, error) {
	log.Info().Msg("Declaring queue")
	ch, err := r.Channel()
	if err != nil {
		log.Err(err).Msg("Failed to create channel, exiting with error message:")
		return nil, err
	}
	r.chs = append(r.chs, ch)
	return &Consumer{Channel: ch}, nil
}
