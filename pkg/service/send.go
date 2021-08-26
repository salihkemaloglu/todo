package service

import (
	"encoding/json"

	"github.com/pkg/errors"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/salihkemaloglu/todo/pkg/config"
	"github.com/salihkemaloglu/todo/pkg/model"
)

// Send sends object to the queue
func Send(o model.Object, config *config.Config) error {
	conn, err := amqp.Dial(config.Queue.URL)
	if err != nil {
		return errors.Wrap(err, "failed to connect to RabbitMQ")
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return errors.Wrap(err, "failed to open a channel")
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"callback", // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		return errors.Wrap(err, "failed to declare a queue")
	}

	body, err := json.Marshal(o)
	if err != nil {
		return errors.Wrap(err, "couldn't marchall object")
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})

	return errors.Wrap(err, "failed to publish a message")
}
