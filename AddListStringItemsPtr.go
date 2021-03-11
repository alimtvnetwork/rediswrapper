package rediswrapper

func (rw *Wrapper) AddListStringItemsPtr(key string, values *[]string) {
	if values == nil {
		return
	}

	// https://redis.io/commands/LSET
	for _, value := range *values {
		rw.client.RPush(rw.ctx, key, value)
	}
}
