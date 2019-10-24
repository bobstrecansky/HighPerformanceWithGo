package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	topic := "go-example"
	partition := 0

	connection, _ := kafka.DialLeader(context.Background(), "tcp", "0.0.0.0:9092", topic, partition)
	connection.SetWriteDeadline(time.Now().Add(1 * time.Second))

	for i := 0; i < 10; i++ {
		connection.WriteMessages(
			kafka.Message{Value: []byte(fmt.Sprintf("Message : %v", i))},
		)
	}

	connection.Close()
}
