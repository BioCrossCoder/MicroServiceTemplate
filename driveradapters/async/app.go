package async

import (
	"encoding/json"
	"main/common"
	"main/infra"
	"main/logics"
	"main/logics/pipeline"
	"main/models"
)

type appBroker struct {
	producer infra.MQProducer
	consumer infra.MQConsumer
	cg       common.ChannelGenerator
	amSvc    logics.AppManagementService
}

func newAppBroker() *appBroker {
	return &appBroker{
		producer: infra.NewKafkaProducerClient(),
		consumer: infra.NewKafkaConsumerClient(),
		cg:       common.NewChannelGenerator(),
		amSvc:    logics.NewAppManagementService(),
	}
}

func (b *appBroker) InjectSenders() {
	pipeline.InjectSender(common.TOPIC_CLEAR_APP, b.sendClearAppMessage)
}

func (b *appBroker) InitListeners() {
	b.handleClearAppMessage()
}

func (b *appBroker) sendClearAppMessage(msg *models.AppIDsMsg) (err error) {
	data, err := json.Marshal(msg)
	if err != nil {
		return
	}
	err = b.producer.Publish(common.TOPIC_CLEAR_APP, data)
	return
}

func (b *appBroker) handleClearAppMessage() {
	channel := b.cg.NextChannel(common.TOPIC_CLEAR_APP)
	handler := func(msg *models.AppIDsMsg) (err error) {
		for _, appID := range msg.AppIDs {
			err = b.amSvc.CancelApp(&models.DeleteAppByIdReqVO{ID: int(appID)})
			if err != nil {
				return
			}
		}
		return
	}
	go b.consumer.Subscribe(common.TOPIC_CLEAR_APP, channel, func(msg []byte) (err error) {
		value := new(models.AppIDsMsg)
		err = json.Unmarshal(msg, value)
		if err != nil {
			return
		}
		err = handler(value)
		return
	})
}
