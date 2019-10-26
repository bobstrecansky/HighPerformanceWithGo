package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	var topic = "go-example"
	var partition = 0
	var connectionType = "tcp"
	var connectionHost = "0.0.0.0"
	var connectionPort = ":9092"

	connection, _ := kafka.DialLeader(context.Background(), connectionType, connectionHost+connectionPort, topic, partition)
	connection.SetWriteDeadline(time.Now().Add(1 * time.Second))

	for i := 0; i < 10; i++ {
		connection.WriteMessages(
			kafka.Message{Value: []byte(fmt.Sprintf("Message : %v", i))},
		)
	}

	connection.Close()
}
