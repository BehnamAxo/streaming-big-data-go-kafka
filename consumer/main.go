package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     "logs",
		GroupID:   "log-consumers",
		Partition: 0,
		MinBytes:  10e3,
		MaxBytes:  10e6,
	})
	defer r.Close()

	count := 0
	start := time.Now()

	for {
		_, err := r.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}

		count++
		if time.Since(start) >= 1*time.Second {
			fmt.Printf("Processed %d messages/sec\n", count)
			count = 0
			start = time.Now()
		}
	}
}
