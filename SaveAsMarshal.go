package rediswrapper

import (
	"encoding/json"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

// SaveAsMarshal marshall data and save as bytes
func (rw *Wrapper) SaveAsMarshal(
	key string,
	jsonData interface{},
) *errorwrapper.Wrapper {
	jsonDataBytes, err := json.Marshal(jsonData)

	if err != nil {
		return errnew.ErrPtr(err)
	}

	return rw.SaveBytes(key, jsonDataBytes)
}
