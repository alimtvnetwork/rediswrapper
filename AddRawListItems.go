package rediswrapper

import "github.com/go-redis/redis/v8"

// AddRawListItems (RPUSH - adds items to Redis List)
//
// Usages RPUSH to push items.
//
// Reference : https://redis.io/commands/LSET
//
// Retrieve items using
// ListUsingLimit or List or ListAsCollection
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
