package consumers

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

func Basic(maxRead int) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
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
