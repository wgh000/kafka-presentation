package consumers

import (
	"github.com/segmentio/kafka-go"
	"time"

	"context"
	"log"
)

func CommitInterval(commit time.Duration, maxRead int) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,
		Topic:          topic,
		GroupID:        "reader",
		CommitInterval: commit,
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
