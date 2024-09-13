package infra

import (
	"fmt"
	"sync"

	"github.com/biocrosscoder/flex/typed/collections/dict"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type MQProducer interface {
	Publish(topic string, msg []byte) (err error)
}

var (
	mqProducer   MQProducer
	producerOnce sync.Once
)

type kafkaProducer struct {
	client *kafka.Producer
}

func NewKafkaProducerClient() MQProducer {
	producerOnce.Do(func() {
		var err error
		p, err := kafka.NewProducer(&kafka.ConfigMap{
			"bootstrap.servers": "localhost",
		})
		if err != nil {
			panic(err)
		}
		mqProducer = &kafkaProducer{client: p}
	})
	return mqProducer
}

func (p *kafkaProducer) Publish(topic string, msg []byte) (err error) {
	ch := make(chan kafka.Event)
	defer close(ch)
	err = p.client.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: msg,
	}, ch)
	if err != nil {
		return
	}
	err = (<-ch).(*kafka.Message).TopicPartition.Error
	return
}

type MQConsumer interface {
	Subscribe(topic, channel string, handler func(msg []byte) error)
}

var (
	mqConsumer   MQConsumer
	consumerOnce sync.Once
)

type kafkaConsumer struct {
	clients dict.Dict[string, *kafka.Consumer]
	m       sync.Locker
}

func NewKafkaConsumerClient() MQConsumer {
	consumerOnce.Do(func() {
		mqConsumer = &kafkaConsumer{
			clients: make(dict.Dict[string, *kafka.Consumer]),
			m:       &sync.Mutex{},
		}
	})
	return mqConsumer
}

func (c *kafkaConsumer) getConsumerClient(channel string) *kafka.Consumer {
	c.m.Lock()
	defer c.m.Unlock()
	if c.clients.Has(channel) {
		return c.clients.Get(channel)
	}
	var err error
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          channel,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}
	c.clients.Set(channel, consumer)
	return consumer
}

func (c *kafkaConsumer) Subscribe(topic, channel string, handler func(msg []byte) error) {
	consumer := c.getConsumerClient(channel)
	err := consumer.Subscribe(topic, nil)
	if err != nil {
		panic(err)
	}
	for {
		event := consumer.Poll(100)
		if event == nil {
			continue
		}
		switch e := event.(type) {
		case *kafka.Message:
			err = handler(e.Value)
			if err != nil {
				fmt.Printf("Error handling message: %v\n", err)
			}
		case kafka.Error:
			fmt.Printf("Error: %v\n", e)
		default:
			fmt.Printf("Ignored event: %v\n", e)
		}
	}
}
