package rediswrapper

import (
	"encoding/json"
)

// SaveAsJson will store an interface in the redis server.
// If Get encounters any errors, it will return the error message.
func (rw *RedisWrapper) SaveAsJson(key string, jsonData interface{}) error {
	jsonDataBytes, err := json.Marshal(jsonData)
	if err != nil {
		return err
	}

	return rw.SaveBytes(key, jsonDataBytes)
}