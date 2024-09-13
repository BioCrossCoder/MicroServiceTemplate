package mq

import (
	"encoding/json"
	"main/common"
	"main/infra"
	"sync"
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
	consumer infra.MQConsumer
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

func (b *userManagementBroker) handleObjectIDMsg(topic, channel string, handler func(*ObjectIDMsg) error) {
	b.consumer.Subscribe(topic, channel, func(msg []byte) (err error) {
		value := new(ObjectIDMsg)
		err = json.Unmarshal(msg, value)
		if err != nil {
			return
		}
		err = handler(value)
		return
	})
}

func (b *userManagementBroker) HandleUserDelete(handler func(*ObjectIDMsg) error) {
	topic := "user.delete"
	channel := b.cg.NextChannel(topic)
	go b.handleObjectIDMsg(topic, channel, handler)
	return
}

func (b *userManagementBroker) HandleDepartmentDelete(handler func(*ObjectIDMsg) error) {
	topic := "dept.delete"
	channel := b.cg.NextChannel(topic)
	go b.handleObjectIDMsg(topic, channel, handler)
	return
}

func (b *userManagementBroker) HandleUserGroupDelete(handler func(*ObjectIDMsg) error) {
	topic := "group.delete"
	channel := b.cg.NextChannel(topic)
	go b.handleObjectIDMsg(topic, channel, handler)
	return
}

type ObjectIDMsg struct {
	ID string `json:"id"`
}
