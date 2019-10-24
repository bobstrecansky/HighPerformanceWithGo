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

	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	conn.SetWriteDeadline(time.Now().Add(1 * time.Second))

	for i := 0; i < 10; i++ {
		conn.WriteMessages(
			kafka.Message{Value: []byte(fmt.Sprintf("Message : %v", i))},
		)
	}

	conn.Close()
}
