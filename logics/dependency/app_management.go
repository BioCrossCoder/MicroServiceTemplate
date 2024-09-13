package dependency

import "main/models"

type AppManagementRepo interface {
	CreateApp(appInfo *CreateAppDTO) (appID int, err error)
	DeleteApp(appID int) (err error)
	GetAllAppInfos() (appInfos []*AppInfoDTO, err error)
	UpdateApp(appID int, appInfo *UpdateAppDTO) (err error)
	ToggleAppSwitch(appID int, enabled bool) (err error)
}

var appManagementRepo AppManagementRepo

func SetAppManagementRepo(repo AppManagementRepo) {
	appManagementRepo = repo
}

func GetAppManagementRepo() AppManagementRepo {
	if appManagementRepo == nil {
		panic("AppManagementRepo is not set")
	}
	return appManagementRepo
}

type CreateAppDTO struct {
	Name  string
	Entry string
}

type AppInfoDTO struct {
	ID          int
	Name        string
	Entry       string
	Enabled     bool
	LimitAccess bool
	Users       []*models.UserObject
}

type UpdateAppDTO struct {
	Name        string
	Entry       string
	LimitAccess bool
	Users       []*models.UserObject
}
