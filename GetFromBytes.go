package rediswrapper

// GetFromJson will bind the json data with the jsonObj.
// If Get encounters any errors, it will return the error message.
func (rw *RedisWrapper) GetFromBytes(key string) ([]byte, error) {
	data, err := rw.rdb.Get(rw.ctx, key).Result()
	if err != nil {
		return nil, err
	}

	return []byte(data), nil
}