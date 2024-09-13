package cache

import (
	"main/infra"
	"main/models"
	"sync"

	"github.com/redis/go-redis/v9"
)

type UserManagementCache interface {
	SetUserAccessors(accessors *models.UserAccessorsCache) (err error)
	GetUserAccessors(userID string) (accessors *models.UserAccessorsCache, err error)
}

var (
	um     UserManagementCache
	umOnce sync.Once
)

type userManagementCache struct {
	client redis.Cmdable
}

func NewUserManagementCache() UserManagementCache {
	umOnce.Do(func() {
		um = &userManagementCache{
			client: infra.NewRedisClient(),
		}
	})
	return um
}

func (c *userManagementCache) SetUserAccessors(accessors *models.UserAccessorsCache) (err error) {
	// ...
	return
}

func (c *userManagementCache) GetUserAccessors(userID string) (accessors *models.UserAccessorsCache, err error) {
	// ...
	return
}
