package rediswrapper

import "github.com/go-redis/redis/v8"

func (rw *Wrapper) GetStringCmd(
	key string,
) *redis.StringCmd {
	return rw.client.Get(rw.ctx, key)
}
