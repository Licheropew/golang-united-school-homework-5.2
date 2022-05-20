package cache

import "time"

type Cache struct {
	cache map[string]customCache
}

type customCache struct {
	value    string
	deadline time.Time
}

func NewCache() *Cache {
	return &Cache{cache: make(map[string]customCache)}
}

func (c Cache) Get(key string) (string, bool) {
	now := time.Now()
	for k, v := range c.cache {
		if k == key && (v.deadline.Before(now) || v.deadline == time.Unix(1<<63-1, 0)) {
			return v.value, true
		}
	}
	return "", false
}

func (c *Cache) Put(key, value string) {
	c.cache[key] = customCache{value: value, deadline: time.Unix(1<<63-1, 0)}
}

func (c Cache) Keys() []string {
	now := time.Now()
	result := []string{}
	for k, v := range c.cache {
		if v.deadline.Before(now) || v.deadline == time.Unix(1<<63-1, 0) {
			result = append(result, k)
		}
	}
	return result
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.cache[key] = customCache{value: value, deadline: deadline}
}
