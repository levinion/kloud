package cache

import "time"

type Cache interface {
	Set(key []byte, value []byte, expiration time.Duration)
	Get(key []byte) ([]byte, error)
}

var defaultCache Cache = nil
