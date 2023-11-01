package cache

import "time"

type CombinedCache struct {
	TTL   time.Time
	Cache map[string]interface{}
}

func NewCombinedCache(cap int, ttl time.Time) CombinedCache {
	return CombinedCache{
		ttl,
		make(map[string]interface{}, cap),
	}
}

func (c *CombinedCache) Add(key string, value interface{}) error {
	return nil
}

func (c *CombinedCache) Get(key string) (interface{}, bool) {
	return nil, false
}
