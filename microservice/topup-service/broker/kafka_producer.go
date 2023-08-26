package broker

import (
	"context"
	"errors"
	"log"

	"github.com/segmentio/kafka-go"
)

func PublishMessage(msg []byte) error {
	w := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "topup",
		Balancer: &kafka.LeastBytes{},
	}

	err := w.WriteMessages(context.Background(), kafka.Message{
		Value: msg,
	})
	if err != nil {
		log.Println("failed to write messages:", err)
		return errors.New("failed to write messages")
	}

	if err := w.Close(); err != nil {
		log.Println("failed to close writer:", err)
		return errors.New("failed to close writer")
	}

	return nil
}
