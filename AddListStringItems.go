package rediswrapper

// AddListStringItems
//
// Usages RPUSH to push items.
// (Redis unordered list, not unique, just list)
//
// Reference : https://redis.io/commands/LSET | https://redis.io/commands/RPUSH
//
// Retrieve items using
// ListUsingLimit or List or ListAsCollection
func (rw *Wrapper) AddListStringItems(
	key string,
	values ...string,
) {
	if values == nil {
		return
	}

	for _, value := range values {
		rw.client.RPush(rw.ctx, key, value)
	}
}
