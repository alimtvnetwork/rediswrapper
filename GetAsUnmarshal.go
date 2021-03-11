package rediswrapper

import (
	"encoding/json"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

// GetAsUnmarshal will unmarshall and get the data for the object
func (rw *Wrapper) GetAsUnmarshal(
	key string,
	jsonObj interface{},
) *errorwrapper.Wrapper {
	data, err := rw.GetRawBytes(key)

	if err != nil {
		return errnew.ErrPtr(err)
	}

	err2 := json.Unmarshal(data, jsonObj)

	if err2 != nil {
		return errnew.ErrPtr(err2)
	}

	return errnew.EmptyPtr
}
