package producers

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"sync"
)

func BatchFillBagWithAck(amount int, requiredAcks kafka.RequiredAcks) {
	w := &kafka.Writer{
		Addr:         addrs,
		Topic:        topic,
		RequiredAcks: requiredAcks,
	}

	wg := sync.WaitGroup{}
	for i := 0; i < amount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = w.WriteMessages(context.Background(), message)
		}()
	}

	wg.Wait()

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
