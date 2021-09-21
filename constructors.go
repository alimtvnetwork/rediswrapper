package rediswrapper

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// New returns a client wrapper to the Redis Server specified by Options.
func New(ctx context.Context, options *redis.Options) *Wrapper {
	rdb := redis.NewClient(options)

	return &Wrapper{
		ctx:     ctx,
		client:  rdb,
		options: options,
	}
}
