package cache

import (
	"sync"
	"time"
)

type Cache interface {
	Get(key string) ([]byte, bool)
	Set(key string, value []byte, ttl time.Duration)
}

type Data struct {
	Value []byte
	Exp   int64
}

type cache struct {
	mutex sync.RWMutex
	items map[string]Data
}

func NewCache() *cache {
	return &cache{
		items: make(map[string]Data),
	}
}

func (d Data) expired() bool {
	if d.Exp == 0 {
		return false
	}

	return time.Now().UnixNano() > d.Exp
}

func (c *cache) Get(key string) ([]byte, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	item, found := c.items[key]
	if !found || item.expired() {
		return nil, false
	}

	return item.Value, true
}

func (c *cache) Set(key string, value []byte, ttl time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	var expiration int64
	if ttl > 0 {
		expiration = time.Now().Add(ttl).UnixNano()
	}

	c.items[key] = Data{
		Value: value,
		Exp:   expiration,
	}
}
