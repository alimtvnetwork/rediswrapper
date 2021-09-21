package rediswrapper

// AddListStringItemsPtr
//
// Usages RPUSH to push items.
// (Redis unordered list, not unique, just list)
//
// Reference : https://redis.io/commands/LSET | https://redis.io/commands/RPUSH
//
// Retrieve items using
// ListUsingLimit or List or ListAsCollection
func (rw *Wrapper) AddListStringItemsPtr(
	key string,
	values *[]string,
) {
	if values == nil || len(*values) == 0 {
		return
	}

	for _, value := range *values {
		rw.client.RPush(
			rw.ctx,
			key,
			value)
	}
}
