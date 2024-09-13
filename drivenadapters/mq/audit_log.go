package mq

import (
	"encoding/json"
	"main/infra"
	"main/models"
	"sync"
)

type AuditLogBroker interface {
	SendLog(logBody *models.AuditLog) (err error)
}

var (
	alb     AuditLogBroker
	albOnce sync.Once
)

type auditLogBroker struct {
	producer infra.MQProducer
}

func NewAuditLogBroker() AuditLogBroker {
	albOnce.Do(func() {
		alb = &auditLogBroker{
			producer: infra.NewKafkaProducerClient(),
		}
	})
	return alb
}

func (b *auditLogBroker) SendLog(logBody *models.AuditLog) (err error) {
	topic := "audit_log.log_management"
	msg, err := json.Marshal(logBody)
	if err != nil {
		return
	}
	err = b.producer.Publish(topic, msg)
	return
}
