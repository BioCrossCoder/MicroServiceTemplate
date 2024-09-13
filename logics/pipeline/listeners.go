package pipeline

import "main/models"

type clearAppMessageListener func(func(*models.AppIDsMsg) error)

var caml clearAppMessageListener

func AddClearAppMessageListener(listener clearAppMessageListener) {
	caml = listener
}
