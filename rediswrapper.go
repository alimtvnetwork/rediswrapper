package rediswrapper

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type RedisWrapper struct {
	ctx context.Context
	rdb *redis.Client
}

// NewClient returns a client wrapper to the Redis Server specified by Options.
func NewClient(ctx context.Context, options *redis.Options) *RedisWrapper {
	rdb := redis.NewClient(options)
	return &RedisWrapper{ctx, rdb }
}