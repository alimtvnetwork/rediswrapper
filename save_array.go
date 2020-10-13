package rediswrapper

import "fmt"

// SaveBytes will store array elements in the redis server.
// If Get encounters any errors, it will return the error message.

func (rw *RedisWrapper) SaveArray(key string, items ...interface{}) error {
	fmt.Print("items ", items) // Also a variadic function.
	rw.rdb.Del(rw.ctx, key)
	return rw.rdb.LPush(rw.ctx, key, items...).Err()
}