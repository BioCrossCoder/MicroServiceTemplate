package mq

import (
	"main/common"
	"main/infra"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type UserManagementBroker interface {
	HandleUserDelete(handler func(*ObjectIDMsg) error)
	HandleDepartmentDelete(handler func(*ObjectIDMsg) error)
	HandleUserGroupDelete(handler func(*ObjectIDMsg) error)
}

var (
	umb     UserManagementBroker
	umbOnce sync.Once
)

type userManagementBroker struct {
	consumer *kafka.Consumer
	cg       common.ChannelGenerator
}

func NewUserManagementBroker() UserManagementBroker {
	umbOnce.Do(func() {
		umb = &userManagementBroker{
			consumer: infra.NewKafkaConsumerClient(),
			cg:       common.NewChannelGenerator(),
		}
	})
	return umb
}

func (b *userManagementBroker) HandleUserDelete(handler func(*ObjectIDMsg) error) {
	// topic := "user.delete"
	// channel := b.cg.nextchannel(topic)
	// ...
	return
}

func (b *userManagementBroker) HandleDepartmentDelete(handler func(*ObjectIDMsg) error) {
	// topic := "dept.delete"
	// channel := b.cg.nextchannel(topic)
	// ...
	return
}

func (b *userManagementBroker) HandleUserGroupDelete(handler func(*ObjectIDMsg) error) {
	// topic := "group.delete"
	// channel := b.cg.nextchannel(topic)
	// ...
	return
}

type ObjectIDMsg struct {
	ID string `json:"id"`
}
