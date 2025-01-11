package broker

import amqp "github.com/rabbitmq/amqp091-go"

type Consumer struct {
	*amqp.Channel
}

func (c *Consumer) Consume(
	/// The name of the queue
	queue string,
	/// The name of the consumer
	consumer string,
	/// AutoAck will automatically acknowledge the message
	autoAck bool,
	/// Exclusive queues are only accessible by the connection that created them
	exclusive bool,
	/// NoLocal will not receive messages published by this connection
	noLocal bool,
	/// NoWait will not wait for a confirmation from the server
	noWait bool,
	/// Arguments for declaring the queue
	args amqp.Table,
) (<-chan amqp.Delivery, error) {
	return c.Channel.Consume(queue, consumer, autoAck, exclusive, noLocal, noWait, args)
}

func (c *Consumer) ConsumeMapper(
	/// The name of the queue
	queue string,
	/// AutoAck will automatically acknowledge the message
	autoAck bool,
	/// Exclusive queues are only accessible by the connection that created them
	exclusive bool,
	/// NoLocal will not receive messages published by this connection
	noLocal bool,
	/// NoWait will not wait for a confirmation from the server
	noWait bool,
	/// Arguments for declaring the queue
	args amqp.Table,
	/// The function to map the delivery to a function
	f func(amqp.Delivery),
) error {
	msgs, err := c.Channel.Consume(queue, "", autoAck, exclusive, noLocal, noWait, args)
	if err != nil {
		return err
	}
	go func() {
		for d := range msgs {
			f(d)
		}
	}()
	return nil
}
