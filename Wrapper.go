package rediswrapper

import (
	"context"
	"strconv"

	"github.com/go-redis/redis/v8"
)

type Wrapper struct {
	ctx     context.Context
	client  *redis.Client
	options *redis.Options
	toStr   *string
}

func (rw *Wrapper) Client() *redis.Client {
	return rw.client
}

func (rw *Wrapper) String() string {
	if rw.toStr != nil {
		return *rw.toStr
	}

	if rw.toStr == nil {
		options := rw.options
		toStr := "Connection at : " + options.Addr + "\n" +
			"Database : " + strconv.Itoa(options.DB) + "\n" +
			"Username : " + options.Username + "\n" +
			"Network : " + options.Network + "\n" +
			"PoolSize : " + strconv.Itoa(options.PoolSize) + "\n"

		rw.toStr = &toStr
	}

	return *rw.toStr
}
