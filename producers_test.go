package kafka_go_presentation

import (
	"github.com/segmentio/kafka-go"
	"kafka-go-presentation/producers"

	"testing"
)

func BenchmarkProducers(b *testing.B) {
	//b.Run("Basic", func(b *testing.B) {
	//	for i := 0; i < b.N; i++ {
	//		producers.Basic(1)
	//	}
	//})

	b.Run("BatchSize", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			producers.BatchSize(100)
		}
	})

	b.Run("BatchTimeout", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			producers.BatchTimeout(100)
		}
	})

	b.Run("Async", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			producers.Async(100)
		}
	})

	b.Run("BatchFillBag", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			producers.BatchFillBag(100)
		}
	})

	b.Run("BatchFillGoroutine", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			producers.BatchFillGoroutine(100)
		}
	})
}

func BenchmarkProducersAcks(b *testing.B) {
	b.Run("RequireNone", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			producers.BatchFillBagWithAck(100, kafka.RequireNone)
		}
	})

	b.Run("RequireOne", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			producers.BatchFillBagWithAck(100, kafka.RequireOne)
		}
	})

	b.Run("RequireAll", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			producers.BatchFillBagWithAck(100, kafka.RequireAll)
		}
	})
}

func TestAsyncCompletion(t *testing.T) {
	producers.AsyncCompletion(3)
}
