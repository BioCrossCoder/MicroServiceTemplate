package pipeline

import (
	"main/models"
	"sync"
)

type AppDeploymentPipeline interface {
	SendClearAppMessage(msg *models.AppIDsMsg) (err error)
}

var (
	adp     AppDeploymentPipeline
	adpOnce sync.Once
)

type appDeploymentPipeline struct{}

func NewAppDeploymentPipeline() AppDeploymentPipeline {
	adpOnce.Do(func() {
		adp = &appDeploymentPipeline{}
	})
	return adp
}

func (p *appDeploymentPipeline) SendClearAppMessage(msg *models.AppIDsMsg) (err error) {
	return cams(msg)
}