package mq

import (
	"main/infra"
	"main/models"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type AuditLogBroker interface {
	SendLog(logBody *models.AuditLog) (err error)
}

var (
	alb     AuditLogBroker
	albOnce sync.Once
)

type auditLogBroker struct {
	producer *kafka.Producer
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
	// topic := "audit_log.log_management"
	return
}
