package pipeline

import "main/models"

type clearAppMessageSender func(msg *models.AppIDsMsg) (err error)

var cams clearAppMessageSender

func AddSendClearAppMessageSender(sender clearAppMessageSender) {
	cams = sender
}
