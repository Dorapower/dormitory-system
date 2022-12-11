package rabbitmq

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
)

var OrderChannel *amqp.Channel

const OrderQueueName = "order"

func Connect() {
	//connect to rabbitmq
	username := os.Getenv("RABBITMQ_USERNAME")
	password := os.Getenv("RABBITMQ_PASSWORD")
	host := os.Getenv("RABBITMQ_HOST")
	port := os.Getenv("RABBITMQ_PORT")

	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", username, password, host, port)
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalln("Error when trying to connect to the rabbitmq server: " + err.Error())
	}
	//create channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalln("Error when trying to create channel: " + err.Error())
	}
	//create queue
	_, err = ch.QueueDeclare(
		OrderQueueName, // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		log.Fatalln("Error when trying to create queue: " + err.Error())
	}
}

func PublishOrderMessage(message []byte) {
	//publish message
	err := OrderChannel.PublishWithContext(
		context.Background(),
		"", // exchange
		OrderQueueName,
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		})
	if err != nil {
		log.Fatalln("Error when trying to publish message: " + err.Error())
	}
}
