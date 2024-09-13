package dependency

type AppDeploymentRepo interface {
	RegisterApp(appName string) (err error)
	CancelApp(appName string) (err error)
	GetAppList() (appList []string, err error)
}

var appDeploymentRepo AppDeploymentRepo

func SetAppDeploymentRepo(repo AppDeploymentRepo) {
	appDeploymentRepo = repo
}

func GetAppDeploymentRepo() AppDeploymentRepo {
	if appDeploymentRepo == nil {
		panic("AppDeploymentRepo not set")
	}
	return appDeploymentRepo
}
