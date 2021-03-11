package rediswrapper

import (
	"encoding/json"

	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func (rw *Wrapper) GetJsonResultsCollection(
	key string,
) (*corejson.ResultsCollection, *errorwrapper.Wrapper) {
	allBytes, err := rw.GetRawBytes(key)

	if err != nil {
		return nil, errnew.NewPtr(errtype.ReadFailed, err)
	}

	var jsonResultsCollection *corejson.ResultsCollection

	err2 := json.Unmarshal(
		allBytes,
		&jsonResultsCollection)

	if err2 != nil {
		return nil, errnew.NewPtr(
			errtype.Unmarshalling, err2)
	}

	return jsonResultsCollection, errnew.EmptyPtr
}
