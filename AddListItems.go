package rediswrapper

func (rw *Wrapper) AddListItems(key string, values ...[]byte) {
	if values == nil {
		return
	}

	// https://redis.io/commands/LSET
	for _, value := range values {
		rw.client.RPush(rw.ctx, key, value)
	}
}
