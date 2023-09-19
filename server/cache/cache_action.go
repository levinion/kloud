package cache

import "time"

func setDefaultCache(cache Cache) {
	defaultCache = cache
}

func Set(key []byte, value []byte, expiration time.Duration) {
	defaultCache.Set(key, value, expiration)
}

func Get(key []byte) ([]byte, error) {
	return defaultCache.Get(key)
}
