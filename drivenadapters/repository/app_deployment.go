package repository

import (
	"main/logics/dependency"
	"main/logics/proxy"
	"os"
)

func init() {
	if os.Getenv("MODE") != "test" {
		dependency.SetAppDeploymentRepo(newAppDeploymentRepo())
	}
}

type appDeploymentRepo struct {
	amProxy func() proxy.AppManagementProxy
}

func newAppDeploymentRepo() dependency.AppDeploymentRepo {
	return &appDeploymentRepo{
		amProxy: proxy.ConnectAppManagementProxy,
	}
}

func (r *appDeploymentRepo) RegisterApp(appName string) (err error) {
	// amSvc := r.amProxy()
	//...
	return
}

func (r *appDeploymentRepo) CancelApp(appName string) (err error) {
	// amSvc := r.amProxy()
	//...
	return
}

func (r *appDeploymentRepo) GetAppList() (appList []string, err error) {
	// amSvc := r.amProxy()
	//...
	return
}
