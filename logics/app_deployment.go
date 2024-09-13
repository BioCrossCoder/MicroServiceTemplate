package logics

import (
	"main/common"
	"main/logics/dependency"
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
}

func NewAppDeploymentService() AppDeploymentService {
	adSvcOnce.Do(func() {
		adSvc = &appDeploymentService{
			repo:      dependency.GetAppDeploymentRepo(),
			eventLoop: common.GetEventLoop(common.Channel),
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
	// s.eventLoop.Trigger("clear_app", []uint64{1, 2, 3})
	return
}

func (s *appDeploymentService) ListApps() (appList []string, err error) {
	// ...
	return
}
