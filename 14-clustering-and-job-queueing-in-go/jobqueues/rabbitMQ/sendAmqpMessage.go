package main

import (
	"log"

	"github.com/streadway/amqp"
)

func main() {
	var username = "guest"
	var password = "guest"
	var protocol = "amqp://"
	var host = "0.0.0.0"
	var port = ":5672/"
	var queueName = "go-queue"

	connectionString := protocol + username + ":" + password + "@" + host + port
	connection, err := amqp.Dial(connectionString)
	if err != nil {
		log.Printf("Could not connect to Local RabbitMQ instance on " + host)
	}
	defer connection.Close()

	ch, err := connection.Channel()
	if err != nil {
		log.Printf("Could not connect to channel")
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		log.Printf("Could not declare queue : " + queueName)
	}

	messageBody := "Hello Gophers!"
	err = ch.Publish("", queue.Name, false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(messageBody),
		})
	log.Printf("Message sent on queue %s : %s", queueName, messageBody)
	if err != nil {
		log.Printf("Message not sent successfully on queue %s", queueName, messageBody)
	}
}
