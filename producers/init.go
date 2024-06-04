package producers

import "github.com/segmentio/kafka-go"

var (
	addrs   = kafka.TCP("localhost:9092", "localhost:9093")
	topic   = "Topic1"
	message = kafka.Message{
		Value: []byte("Hello World!"),
	}
)
