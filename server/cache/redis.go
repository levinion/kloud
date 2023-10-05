package cache

import (
	"context"
	"kloud/constant"
	"time"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	client *redis.Client
}

func UseRedisCache() *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     constant.RedisAddr,
		Password: "",
		DB:       0,
	})
	cache := &Redis{client: client}
	setDefaultCache(cache)
	return cache
}

func (r *Redis) Set(key []byte, value []byte, expiration time.Duration) {
	r.client.Set(context.Background(), string(key), value, expiration)
}

func (r *Redis) Get(key []byte) ([]byte, error) {
	str, err := r.client.Get(context.Background(), string(key)).Result()
	return []byte(str), err
}
