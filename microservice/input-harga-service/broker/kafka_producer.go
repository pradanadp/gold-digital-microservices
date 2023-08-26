package broker

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func PublishMessage(msg []byte) error {
	w := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "input-harga",
		Balancer: &kafka.LeastBytes{},
	}

	err := w.WriteMessages(context.Background(), kafka.Message{
		Value: msg,
	})
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}

	return nil
}
