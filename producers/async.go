package producers

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

func Async(amount int) {
	w := &kafka.Writer{
		Addr:  addrs,
		Topic: topic,
		Async: true,
	}

	for i := 0; i < amount; i++ {
		_ = w.WriteMessages(context.Background(), message)
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}

func AsyncCompletion(amount int) {
	w := &kafka.Writer{
		Addr:  addrs,
		Topic: topic,
		Async: true,
		Completion: func(messages []kafka.Message, err error) {
			log.Println("completed: ", len(messages))
		},
	}

	for i := 0; i < amount; i++ {
		_ = w.WriteMessages(context.Background(), message)
	}

	log.Println("wrote")

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}

	log.Println("closed")
}
