package rediswrapper

import (
	"encoding/json"

	"gitlab.com/evatix-go/core/coredata/coredynamic"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

// GetAsUnmarshal will unmarshall and get the data for the object
func (rw *Wrapper) GetAsUnmarshal(
	key string,
	jsonObj interface{},
) *errorwrapper.Wrapper {
	data, err := rw.GetRawBytes(key)

	if err != nil {
		return errnew.MessagesPtr(
			errtype.Unmarshalling,
			err.Error(),
			coredynamic.TypeName(jsonObj),
			errorwrapper.SimpleReferencesCompileOptimized(errtype.KeyNotFound, key))
	}

	err2 := json.Unmarshal(data, jsonObj)

	if err2 != nil {
		return errnew.MessagesPtr(
			errtype.Unmarshalling,
			err2.Error(),
			coredynamic.TypeName(jsonObj),
			errorwrapper.SimpleReferencesCompileOptimized(errtype.KeyValidationFailed, key))
	}

	return errnew.EmptyPtr
}
