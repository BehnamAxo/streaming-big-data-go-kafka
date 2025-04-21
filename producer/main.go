package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "logs",
		Balancer: &kafka.LeastBytes{},
	})

	defer writer.Close()

	for {
		message := fmt.Sprintf(`{"level": "INFO", "timestamp": "%s", "message": "User %d clicked button"}`,
			time.Now().Format(time.RFC3339), rand.Intn(1000))

		err := writer.WriteMessages(
			context.Background(),
			kafka.Message{
				Key:   []byte("log"),
				Value: []byte(message),
			},
		)

		if err != nil {
			fmt.Println("Write error:", err)
		} else {
			fmt.Println("Sent:", message)
		}

		time.Sleep(10 * time.Millisecond)
	}
}
