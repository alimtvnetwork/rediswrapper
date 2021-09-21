package rediswrapper

import "github.com/go-redis/redis/v8"

func (rw *Wrapper) GetListRangeCmd(key string, start, end int64) *redis.StringCmd {
	return rw.client.GetRange(rw.ctx, key, start, end)
}
