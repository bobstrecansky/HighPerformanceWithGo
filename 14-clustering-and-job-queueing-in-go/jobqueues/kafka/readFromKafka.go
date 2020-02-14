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
		log.Fatal("Could not create a Kafka Connection")
	}

	connection.SetReadDeadline(time.Now().Add(10 * time.Second))
	readBatch := connection.ReadBatch(500, 500000)

	byteString := make([]byte, 500)
	for {
		_, err := readBatch.Read(byteString)
		if err != nil {
			break
		}
		fmt.Println(string(byteString))
	}

	readBatch.Close()
	connection.Close()
}
