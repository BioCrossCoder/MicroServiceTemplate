package logics

import (
	"main/common"
	"main/logics/dependency"
	"main/logics/pipeline"
	"main/models"
	"sync"
)

type AppDeploymentService interface {
	InstallApp(appName string) (err error)
	UninstallApp(appName string) (err error)
	ClearInvalidApps() (err error)
	ListApps() (appList []string, err error)
}

var (
	adSvc     AppDeploymentService
	adSvcOnce sync.Once
)

type appDeploymentService struct {
	repo      dependency.AppDeploymentRepo
	eventLoop common.EventLoop
	pipeline  pipeline.AppDeploymentPipeline
}

func NewAppDeploymentService() AppDeploymentService {
	adSvcOnce.Do(func() {
		adSvc = &appDeploymentService{
			repo:      dependency.GetAppDeploymentRepo(),
			eventLoop: common.GetEventLoop(common.Channel),
			pipeline:  pipeline.NewAppDeploymentPipeline(),
		}
	})
	return adSvc
}

func (s *appDeploymentService) InstallApp(appName string) (err error) {
	// ...
	return
}

func (s *appDeploymentService) UninstallApp(appName string) (err error) {
	// ...
	return
}

func (s *appDeploymentService) ClearInvalidApps() (err error) {
	// ...
	invalidAppIDs := []uint64{1, 2, 3}
	s.eventLoop.Trigger("clear_app", invalidAppIDs)
	msg := &models.AppIDsMsg{
		AppIDs: invalidAppIDs,
	}
	err = s.pipeline.SendClearAppMessage(msg)
	return
}

func (s *appDeploymentService) ListApps() (appList []string, err error) {
	// ...
	return
}
