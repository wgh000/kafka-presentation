package producers

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

func BatchFillGoroutine(amount int) {
	w := &kafka.Writer{
		Addr:  addrs,
		Topic: topic,
	}

	batch := make([]kafka.Message, 0, 100)

	for i := 0; i < amount; i++ {
		batch = append(batch, kafka.Message{
			Value: []byte("Hello World!"),
		})

		if len(batch) == 100 {
			_ = w.WriteMessages(context.Background(), batch...)
			batch = batch[:0]
		}
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
