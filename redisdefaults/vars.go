package redisdefaults

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var (
	Options = &redis.Options{
		Addr:     Address,
		Password: EmptyPassword, // no password set
		DB:       FirstDatabase, // use default DB
	}

	Context = context.Background()
)
