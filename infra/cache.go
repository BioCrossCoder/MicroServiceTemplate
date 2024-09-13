package infra

import (
	"context"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	cacheClient *redis.Client
	cacheOnce   sync.Once
)

func NewRedisClient() *redis.Client {
	cacheOnce.Do(func() {
		cacheClient = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
		_, err := cacheClient.Ping(context.TODO()).Result()
		if err != nil {
			panic(err)
		}
	})
	return cacheClient
}
