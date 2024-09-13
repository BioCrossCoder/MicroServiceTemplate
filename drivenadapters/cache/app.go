package cache

import (
	"main/infra"
	"main/models"
	"sync"

	"github.com/redis/go-redis/v9"
)

type AppCache interface {
	SetAppInfo(appID string, appInfo *models.AppCache) (err error)
	GetAppInfo(appID string) (appInfo *models.AppCache, err error)
	GetAppInfos() (appInfos map[string]*models.AppCache, err error)
	SetAppInfos(appInfos map[string]*models.AppCache) (err error)
}

var (
	ac     AppCache
	acOnce sync.Once
)

type appCache struct {
	client redis.Cmdable
}

func NewAppCache() AppCache {
	acOnce.Do(func() {
		ac = &appCache{
			client: infra.NewRedisClient(),
		}
	})
	return ac
}

func (c *appCache) SetAppInfo(appID string, appInfo *models.AppCache) (err error) {
	// ...
	return
}

func (c *appCache) GetAppInfo(appID string) (appInfo *models.AppCache, err error) {
	// ...
	return
}

func (c *appCache) GetAppInfos() (appInfos map[string]*models.AppCache, err error) {
	// ...
	return
}

func (c *appCache) SetAppInfos(appInfos map[string]*models.AppCache) (err error) {
	// ...
	return
}
