package pipeline

import (
	"main/models"
)

var senders = map[string]any{}

func InjectSender(topic string, sender any) {
	senders[topic] = sender
}

type clearAppMessageSender func(msg *models.AppIDsMsg) (err error)
