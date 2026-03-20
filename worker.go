package main

import (
	"context"
	"encoding/json"
	"math/rand"

	"github.com/segmentio/kafka-go"
)

func process(msg kafka.Message, retryWriter, dlqWriter *kafka.Writer) {
	var event Event
	json.Unmarshal(msg.Value, &event)

	if rand.Intn(4) == 0 {
		event.Retry++

		data, _ := json.Marshal(event)

		if event.Retry > 2 {
			dlqWriter.WriteMessages(context.Background(), kafka.Message{
				Key:   msg.Key,
				Value: data,
			})
			return
		}

		retryWriter.WriteMessages(context.Background(), kafka.Message{
			Key:   msg.Key,
			Value: data,
		})
		return
	}
}
