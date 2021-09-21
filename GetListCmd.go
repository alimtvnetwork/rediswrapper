package rediswrapper

import (
	"github.com/go-redis/redis/v8"
)

func (rw *Wrapper) GetListCmd(key string) *redis.StringSliceCmd {
	return rw.client.LRange(rw.ctx, key, 0, -1)
}
