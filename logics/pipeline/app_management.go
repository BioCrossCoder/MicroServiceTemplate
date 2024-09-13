package pipeline

import (
	"main/models"
	"sync"
)

type AppManagementPipeline interface {
	HandleClearAppMessage(handler func(*models.AppIDsMsg) error)
}

var (
	amp     AppManagementPipeline
	ampOnce sync.Once
)

type appManagementPipeline struct{}

func NewAppManagementPipeline() AppManagementPipeline {
	ampOnce.Do(func() {
		amp = &appManagementPipeline{}
	})
	return amp
}

func (p *appManagementPipeline) HandleClearAppMessage(handler func(*models.AppIDsMsg) error) {
	caml(handler)
}
