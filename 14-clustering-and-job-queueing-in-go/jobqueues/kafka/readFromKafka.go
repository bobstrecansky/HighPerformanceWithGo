package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {

	topic := "go-example"
	partition := 0

	connection, err := kafka.DialLeader(context.Background(), "tcp", "0.0.0.0:9092", topic, partition)
	if err != nil {
		log.Printf("Could not create a Kafka Connection")
	}

	connection.SetReadDeadline(time.Now().Add(1 * time.Second))
	readBatch := connection.ReadBatch(500, 500000)

	byteString := make([]byte, 500)
	for {
		_, err := readBatch.Read(byteString)
		if err != nil {
			log.Printf("Could not read byteString")
		}
		fmt.Println(string(byteString))
	}

	readBatch.Close()
	connection.Close()
}
