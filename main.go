package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

func main() {
	producer := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "events",
		Balancer: &kafka.Hash{},
	}

	retryWriter := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "events-retry",
		Balancer: &kafka.Hash{},
	}

	dlqWriter := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "events-dlq",
		Balancer: &kafka.Hash{},
	}

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		GroupID: "group-1",
		Topic:   "events",
	})

	retryReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		GroupID: "group-retry",
		Topic:   "events-retry",
	})

	go func() {
		for {
			event := Event{
				ID:    uuid.New().String(),
				Data:  "payload",
				Retry: 0,
			}

			data, _ := json.Marshal(event)

			producer.WriteMessages(context.Background(), kafka.Message{
				Key:   []byte(event.ID),
				Value: data,
			})

			time.Sleep(300 * time.Millisecond)
		}
	}()

	go func() {
		for {
			msg, _ := reader.ReadMessage(context.Background())
			process(msg, retryWriter, dlqWriter)
		}
	}()

	for {
		msg, _ := retryReader.ReadMessage(context.Background())
		process(msg, retryWriter, dlqWriter)
	}
}
