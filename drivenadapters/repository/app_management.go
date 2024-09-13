package repository

import (
	"main/drivenadapters/cache"
	"main/drivenadapters/db"
	"main/drivenadapters/http"
	"main/drivenadapters/mq"
	"main/logics/dependency"
	"os"
)

func init() {
	if os.Getenv("MODE") != "test" {
		dependency.SetAppManagementRepo(newAppManagementRepo())
	}
}

type appManagementRepo struct {
	appDB          db.AppStore
	appCache       cache.AppCache
	userMgntCache  cache.UserManagementCache
	userMgnt       http.UserManagementClient
	auditLogBroker mq.AuditLogBroker
	userMgntBroker mq.UserManagementBroker
}

func newAppManagementRepo() dependency.AppManagementRepo {
	return &appManagementRepo{
		appDB:          db.NewAppStore(),
		appCache:       cache.NewAppCache(),
		userMgntCache:  cache.NewUserManagementCache(),
		userMgnt:       http.NewUserManagementClient(),
		auditLogBroker: mq.NewAuditLogBroker(),
		userMgntBroker: mq.NewUserManagementBroker(),
	}
}

func (r *appManagementRepo) CreateApp(appInfo *dependency.CreateAppDTO) (appID int, err error) {
	// ...
	return
}

func (r *appManagementRepo) DeleteApp(appID int) (err error) {
	// ...
	return
}

func (r *appManagementRepo) GetAllAppInfos() (appInfos []*dependency.AppInfoDTO, err error) {
	// ...
	return
}

func (r *appManagementRepo) UpdateApp(appID int, appInfo *dependency.UpdateAppDTO) (err error) {
	// ...
	return
}

func (r *appManagementRepo) ToggleAppSwitch(appID int, enabled bool) (err error) {
	// ...
	return
}
