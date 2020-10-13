package rediswrapper

import (
	"encoding/json"
)

func (rw *RedisWrapper) GetFromJson(key string, jsonObj interface{}) error {
	data, err := rw.rdb.Get(rw.ctx, key).Result()
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(data), jsonObj); err != nil {
		return err
	}

	return nil
}