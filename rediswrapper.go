package rediswrapper

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var(
	ctx = context.Background()
)

type RedisWrapper struct {
	*redis.Client
}

// NewClient returns a client wrapper to the Redis Server specified by Options.
func NewClient(options *redis.Options) *RedisWrapper {
	rdb := redis.NewClient(options)
	return &RedisWrapper{rdb}
}