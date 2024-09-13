package infra

import (
	"context"
	"sync"
	"time"

	"github.com/biocrosscoder/flex/typed/collections/dict"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
)

type DLM interface {
	TryLock(key string) (success bool, lockID string)
	UnLock(key string, lockID string) (success bool, err error)
}

var (
	dlm     DLM
	dlmOnce sync.Once
)

type dlmImpl struct {
	store      *redsync.Redsync
	expiration time.Duration
	m          sync.Locker
	locks      dict.Dict[string, *redsync.Mutex]
	contexts   dict.Dict[string, context.CancelFunc]
}

func NewDLM() DLM {
	dlmOnce.Do(func() {
		client := NewRedisClient()
		pool := goredis.NewPool(client)
		rs := redsync.New(pool)
		dlm = &dlmImpl{
			store:      rs,
			expiration: 10 * time.Second,
			m:          &sync.Mutex{},
			locks:      make(dict.Dict[string, *redsync.Mutex]),
			contexts:   make(dict.Dict[string, context.CancelFunc]),
		}
	})
	return dlm
}

func (d *dlmImpl) TryLock(key string) (success bool, lockID string) {
	d.m.Lock()
	defer d.m.Unlock()
	ctx, cancel := context.WithCancel(context.Background())
	lock := d.store.NewMutex(key, redsync.WithExpiry(d.expiration))
	err := lock.TryLock()
	success = err == nil
	if !success {
		cancel()
		return
	}
	d.locks.Set(key, lock)
	d.contexts.Set(key, cancel)
	go d.keepLock(ctx, key)
	lockID = lock.Value()
	return
}

func (d *dlmImpl) UnLock(key string, lockID string) (success bool, err error) {
	d.m.Lock()
	defer d.m.Unlock()
	lock := d.locks.Get(key, d.store.NewMutex(key))
	if lock.Value() == lockID {
		success, err = lock.Unlock()
	}
	if err == nil && success {
		d.locks.Delete(key)
		if d.contexts.Has(key) {
			d.contexts.Get(key)()
			d.contexts.Delete(key)
		}
	}
	return
}

func (d *dlmImpl) keepLock(ctx context.Context, key string) {
	lock := d.locks.Get(key)
	ticker := time.NewTicker(d.expiration / 2)
	defer ticker.Stop()
	var err error
	for err == nil {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			_, err = lock.Extend()
		}
	}
}
