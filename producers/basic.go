package producers

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

func Basic(amount int) {
	w := &kafka.Writer{
		Addr:  addrs,
		Topic: topic,
	}

	for i := 0; i < amount; i++ {
		_ = w.WriteMessages(context.Background(), message)
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
