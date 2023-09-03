package cache

import (
	"sync"
	"time"
)

type Cache struct {
	m       sync.Map
	timer   *time.Timer
	timerMu sync.Mutex
}

type cacheItem struct {
	value     interface{}
	createdAt time.Time
}

func NewCache() *Cache {
	return &Cache{}
}

func (c *Cache) Get(key string) interface{} {
	value, ok := c.m.Load(key)

	if !ok {
		return nil
	}

	return value.(cacheItem).value
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.m.Store(key, cacheItem{value, time.Now()})

	if ttl > 0 {
		c.timerMu.Lock()

		if c.timer == nil {
			c.timer = time.AfterFunc(ttl, func() {
				expiredKeys := make([]string, 0)
				c.m.Range(func(key, value interface{}) bool {
					if time.Since(value.(cacheItem).createdAt) > ttl {
						expiredKeys = append(expiredKeys, key.(string))
					}

					return true
				})

				c.timerMu.Lock()
				defer c.timerMu.Unlock()
				if c.timer != nil {
					c.timer.Stop()
					c.timer = nil
				}

				for _, key := range expiredKeys {
					c.m.Delete(key)
				}
			})
		} else {
			c.timer.Reset(ttl)
		}

		c.timerMu.Unlock()
	}
}
