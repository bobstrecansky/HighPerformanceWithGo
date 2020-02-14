package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	var topic = "go-example"
	var partition = 0
	var connectionType = "tcp"
	var connectionHost = "0.0.0.0"
	var connectionPort = ":9092"

	connection, err := kafka.DialLeader(context.Background(), connectionType, connectionHost+connectionPort, topic, partition)
	if err != nil {
		log.Fatal(err)
	}
	connection.SetWriteDeadline(time.Now().Add(10 * time.Second))

	for i := 0; i < 10; i++ {
		connection.WriteMessages(
			kafka.Message{Value: []byte(fmt.Sprintf("Message : %v", i))},
		)
	}

	connection.Close()
}
