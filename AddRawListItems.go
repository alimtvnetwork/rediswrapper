package rediswrapper

import "github.com/go-redis/redis/v8"

// Reference : https://redis.io/commands/LSET
func (rw *Wrapper) AddRawListItems(
	key string,
	values ...interface{},
) *redis.IntCmd {
	if values == nil {
		return nil
	}

	return rw.client.RPush(
		rw.ctx,
		key,
		values...)
}
