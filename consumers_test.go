package kafka_go_presentation

import (
	"kafka-go-presentation/consumers"
	"time"

	"testing"
)

func BenchmarkBasic(b *testing.B) {
	b.Run("Basic", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			consumers.Basic(1000)
		}
	})
}

func BenchmarkBatchSize(b *testing.B) {
	b.Run("Batch 1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			consumers.Batch(1, 1000)
		}
	})

	b.Run("Batch 10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			consumers.Batch(10, 1000)
		}
	})

	b.Run("Batch 100", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			consumers.Batch(100, 1000)
		}
	})
}

func BenchmarkCommitInterval(b *testing.B) {
	b.Run("CommitInterval 0", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			consumers.CommitInterval(0, 1000)
		}
	})

	b.Run("CommitInterval 1ns", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			consumers.CommitInterval(time.Nanosecond, 1000)
		}
	})

	b.Run("CommitInterval 1us", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			consumers.CommitInterval(time.Microsecond, 1000)
		}
	})

	b.Run("CommitInterval 1ms", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			consumers.CommitInterval(time.Millisecond, 1000)
		}
	})

	b.Run("CommitInterval 1s", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			consumers.CommitInterval(time.Second, 1000)
		}
	})
}
