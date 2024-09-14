package pipeline

import "fmt"

type senderFunc[T any] func(T) error

var senders = map[string]any{}

func InjectSender(topic string, sender any) {
	senders[topic] = sender
}

func sendMessage[T any](topic string, msg T) error {
	f, ok := senders[topic]
	if !ok {
		return fmt.Errorf("sender for topic %s not found", topic)
	}
	sender, ok := f.(senderFunc[T])
	if !ok {
		return fmt.Errorf("sender for topic %s has wrong type", topic)
	}
	return sender(msg)
}
