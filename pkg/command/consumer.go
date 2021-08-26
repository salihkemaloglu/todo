package command

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/salihkemaloglu/todo/pkg/config"
	"github.com/salihkemaloglu/todo/pkg/model"
	"github.com/salihkemaloglu/todo/pkg/service"
	"github.com/spf13/cobra"
)

// Consumer Service
func NewConsumerRun(config *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run-consumer",
		Short: "Run consumer service",
		Long:  `Run consumer backend service.`,
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			runConsumer(config)
		},
	}

	return cmd
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s", msg, err)
	}
}

func runConsumer(config *config.Config) {
	conn, err := amqp.Dial(config.Queue.URL)
	failOnError(err, "failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"callback", // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	failOnError(err, "failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			o := model.Object{}
			err = json.Unmarshal(d.Body, &o)
			failOnError(err, "failed to unmarshal object")
			err = service.Receive(o, config)
			failOnError(err, "failed to receive object from queue")
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
