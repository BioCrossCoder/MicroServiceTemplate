package infra

import (
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var (
	mqProducer   *kafka.Producer
	producerOnce sync.Once
)

func NewKafkaProducerClient() *kafka.Producer {
	producerOnce.Do(func() {
		var err error
		mqProducer, err = kafka.NewProducer(&kafka.ConfigMap{
			"bootstrap.servers": "localhost",
		})
		if err != nil {
			panic(err)
		}
	})
	return mqProducer
}

var (
	mqConsumer   *kafka.Consumer
	consumerOnce sync.Once
)

func NewKafkaConsumerClient() *kafka.Consumer {
	consumerOnce.Do(func() {
		var err error
		mqConsumer, err = kafka.NewConsumer(&kafka.ConfigMap{
			"bootstrap.servers": "localhost:9092",
			"group.id":          "test-group",
			"auto.offset.reset": "earliest",
		})
		if err != nil {
			panic(err)
		}
	})
	return mqConsumer
}
