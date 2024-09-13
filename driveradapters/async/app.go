package async

import (
	"encoding/json"
	"main/common"
	"main/infra"
	"main/logics/pipeline"
	"main/models"
	"os"
)

func init() {
	if os.Getenv("MODE") != "test" {
		b := newAppBroker()
		pipeline.AddSendClearAppMessageSender(b.SendClearAppMessage)
		pipeline.AddClearAppMessageListener(b.HandleClearAppMessage)
	}
}

type appBroker struct {
	producer infra.MQProducer
	consumer infra.MQConsumer
	cg       common.ChannelGenerator
}

func newAppBroker() *appBroker {
	return &appBroker{
		producer: infra.NewKafkaProducerClient(),
		consumer: infra.NewKafkaConsumerClient(),
		cg:       common.NewChannelGenerator(),
	}
}

func (b *appBroker) SendClearAppMessage(msg *models.AppIDsMsg) (err error) {
	topic := "clear_app"
	data, err := json.Marshal(msg)
	if err != nil {
		return
	}
	err = b.producer.Publish(topic, data)
	return
}

func (b *appBroker) HandleClearAppMessage(handler func(*models.AppIDsMsg) error) {
	topic := "clear_app"
	channel := b.cg.NextChannel(topic)
	go b.consumer.Subscribe(topic, channel, func(msg []byte) (err error) {
		value := new(models.AppIDsMsg)
		err = json.Unmarshal(msg, value)
		if err != nil {
			return
		}
		err = handler(value)
		return
	})
}
