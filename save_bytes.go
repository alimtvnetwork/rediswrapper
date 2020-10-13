package rediswrapper

// SaveAsJson will store bytes in the redis server.
// If Get encounters any errors, it will return the error message.

func (rw *RedisWrapper) SaveBytes(key string, byteData []byte) error {
	return rw.rdb.Set(rw.ctx, key, byteData, 0).Err()
}