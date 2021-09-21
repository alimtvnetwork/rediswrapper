package rediswrapper

import (
	"encoding/json"
	"time"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

// SaveAsMarshalExp marshall data and save as bytes with expiration
//
// exp 0 means never removes
func (rw *Wrapper) SaveAsMarshalExp(
	key string,
	jsonData interface{},
	exp time.Duration,
) *errorwrapper.Wrapper {
	jsonDataBytes, err := json.Marshal(jsonData)

	if err != nil {
		return errnew.ErrPtr(err)
	}

	return rw.SaveBytesExp(
		key,
		jsonDataBytes,
		exp)
}
