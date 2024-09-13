package common

import (
	"context"
	"errors"
	"sync"

	"github.com/biocrosscoder/flex/typed/collections/dict"
)

func init() {
	eventLoops = make(dict.Dict[string, EventLoop])
	elm = &sync.Mutex{}
}

var (
	eventLoops dict.Dict[string, EventLoop]
	elm        sync.Locker
)

func GetEventLoop(key string) EventLoop {
	elm.Lock()
	defer elm.Unlock()
	if eventLoops.Has(key) {
		return eventLoops.Get(key)
	}
	return eventLoops.Set(key, newEventLoop()).Get(key)
}

type EventLoop interface {
	Start() (err error)
	Stop() (err error)
	Trigger(event string, payload any)
	AddListener(event string, handler func(any) error) (key uint64)
	RemoveListener(event string, key uint64) (err error)
}

type eventLoop struct {
	queue    chan eventEntity
	context  context.Context
	cancel   context.CancelFunc
	m        sync.Locker
	handlers dict.Dict[string, dict.Dict[uint64, func(any) error]]
	hm       sync.Locker
}

type eventEntity struct {
	name    string
	payload any
}

func newEventLoop() EventLoop {
	return &eventLoop{
		queue:    make(chan eventEntity, 100),
		m:        &sync.Mutex{},
		handlers: make(dict.Dict[string, dict.Dict[uint64, func(any) error]]),
		hm:       &sync.Mutex{},
	}
}

func (l *eventLoop) Start() (err error) {
	l.m.Lock()
	defer l.m.Unlock()
	if l.context != nil {
		err = errors.New("EventLoop is already running")
		return
	}
	l.context = context.TODO()
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			case e := <-l.queue:
				l.hm.Lock()
				handlers := l.handlers.Get(e.name)
				l.hm.Unlock()
				for _, handler := range handlers {
					handler(e.payload)
				}
			}
		}
	}(l.context)
	return
}

func (l *eventLoop) Stop() (err error) {
	l.m.Lock()
	defer l.m.Unlock()
	if l.cancel == nil {
		err = errors.New("EventLoop is not running")
		return
	}
	l.cancel()
	l.cancel = nil
	l.context = nil
	return
}

func (l *eventLoop) Trigger(event string, payload any) {
	go func() {
		l.queue <- eventEntity{name: event, payload: payload}
	}()
}

func (l *eventLoop) AddListener(event string, handler func(any) error) (key uint64) {
	l.hm.Lock()
	defer l.hm.Unlock()
	if !l.handlers.Has(event) {
		l.handlers.Set(event, make(dict.Dict[uint64, func(any) error]))
	}
	handlers := l.handlers.Get(event)
	key = ShortID()
	handlers.Set(key, handler)
	return
}

func (l *eventLoop) RemoveListener(event string, key uint64) (err error) {
	l.hm.Lock()
	defer l.hm.Unlock()
	if !l.handlers.Has(event) {
		err = errors.New("No such event")
		return
	}
	handlers := l.handlers.Get(event)
	if !handlers.Has(key) {
		err = errors.New("No such listener")
		return
	}
	handlers.Delete(key)
	if handlers.Empty() {
		l.handlers.Delete(event)
	}
	return
}
