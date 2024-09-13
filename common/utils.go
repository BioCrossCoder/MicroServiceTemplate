package common

import (
	"fmt"
	"math/rand"
	"sync"

	"github.com/biocrosscoder/flex/typed/collections"
	"github.com/yitter/idgenerator-go/idgen"
)

func init() {
	initUniqueIDGenerator()
}

func initUniqueIDGenerator() {
	options := idgen.NewIdGeneratorOptions(uint16(rand.Intn(2048)))
	options.WorkerIdBitLength = 11
	options.SeqBitLength = 5
	options.BaseTime = 1690777000000
	idgen.SetIdGenerator(options)
}

func ShortID() uint64 {
	return uint64(idgen.NextId())
}

type ChannelGenerator interface {
	NextChannel(topic string) (channel string)
}

var (
	cg     ChannelGenerator
	cgOnce sync.Once
)

type channelGenerator struct {
	counter *collections.Counter[string]
	m       sync.Locker
}

func NewChannelGenerator() ChannelGenerator {
	cgOnce.Do(func() {
		cg = &channelGenerator{
			counter: collections.NewCounter([]string{}),
			m:       &sync.Mutex{},
		}
	})
	return cg
}

func (g *channelGenerator) NextChannel(topic string) (channel string) {
	g.m.Lock()
	defer g.m.Unlock()
	order := g.counter.Get(topic)
	g.counter.Increment(topic)
	channel = fmt.Sprintf("%s_%d", Channel, order)
	return
}
