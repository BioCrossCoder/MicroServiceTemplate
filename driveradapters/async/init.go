package async

import "sync"

type MessageSender interface {
	InjectSenders()
}

type MessageListener interface {
	InitListeners()
}

type MessageGateway interface {
	RegisterSubscribeAPI()
	RegisterPublishAPI()
}

var (
	mg     MessageGateway
	mgOnce sync.Once
)

type messageGateway struct {
	senders   []MessageSender
	listeners []MessageListener
}

func NewMessageGateway() MessageGateway {
	mgOnce.Do(func() {
		mg = &messageGateway{
			senders: []MessageSender{
				newAppBroker(),
			},
			listeners: []MessageListener{
				newAppBroker(),
			},
		}
	})
	return mg
}

func (g *messageGateway) RegisterSubscribeAPI() {
	for _, sender := range g.senders {
		sender.InjectSenders()
	}
}

func (g *messageGateway) RegisterPublishAPI() {
	for _, listener := range g.listeners {
		listener.InitListeners()
	}
}
