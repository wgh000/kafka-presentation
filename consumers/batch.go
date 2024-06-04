package consumers

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

func Batch(batchSize, maxRead int) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:       brokers,
		Topic:         topic,
		QueueCapacity: batchSize,
	})

	i := 0

	for {
		_, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}

		i++
		if i == maxRead {
			break
		}
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
