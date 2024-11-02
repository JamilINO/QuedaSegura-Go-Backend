package queue

import (
	"context"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"quedasegura.com/m/v2/proto/convert"
)

func Send() {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_STR"))
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"teste2", // name
		true,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := convert.QuedaPayload{
		MacAddr: "da:0c:5e:d5:14:12",
		Time: timestamppb.Now(),
		Intensity: 0.5,
	}

	msg, _ := proto.Marshal(&body)

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", msg)
}