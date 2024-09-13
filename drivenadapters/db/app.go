package db

import (
	"main/infra"
	"main/models"
	"sync"

	"gorm.io/gorm"
)

type AppStore interface {
	CreateApp(appInfo *models.AppPO) (err error)
	RemoveApp(appID uint64) (err error)
	FindAppBy(appID uint64) (app *models.AppPO, err error)
	UpdateApp(appInfo *models.AppPO, users []*models.AppUserPO) (err error)
	UpdateAppFields(appID uint64, fields map[string]any) (err error)
	RemoveUsers(userID string) (err error)
	GetAllUsers() (users []*models.AppUserPO, err error)
	GetAllApps() (appInfos []*models.AppPO, err error)
}

var (
	asdb     AppStore
	asdbOnce sync.Once
)

type appStore struct {
	db *gorm.DB
}

func NewAppStore() AppStore {
	asdbOnce.Do(func() {
		asdb = &appStore{
			db: infra.NewMySQLClient(),
		}
	})
	return asdb
}

func (d *appStore) CreateApp(appInfo *models.AppPO) (err error) {
	// ...
	return
}

func (d *appStore) RemoveApp(appID uint64) (err error) {
	// ...
	return
}

func (d *appStore) FindAppBy(appID uint64) (app *models.AppPO, err error) {
	// ...
	return
}

func (d *appStore) UpdateApp(appInfo *models.AppPO, users []*models.AppUserPO) (err error) {
	// ...
	return
}

func (d *appStore) UpdateAppFields(appID uint64, fields map[string]any) (err error) {
	// ...
	return
}

func (d *appStore) RemoveUsers(userID string) (err error) {
	// ...
	return
}

func (d *appStore) GetAllUsers() (users []*models.AppUserPO, err error) {
	// ...
	return
}

func (d *appStore) GetAllApps() (appInfos []*models.AppPO, err error) {
	// ...
	return
}
