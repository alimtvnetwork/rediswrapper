package rediswrapper

import (
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/defaulterr"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func (rw *Wrapper) GetJsonResultsCollectionUnmarshallInto(
	key string,
	unmarshallingItems ...interface{},
) (*corejson.ResultsCollection, *errorwrapper.Wrapper) {
	if unmarshallingItems == nil {
		return nil, errnew.NewPtr(
			errtype.Unmarshalling,
			defaulterr.UnMarshallingFailedDueToNilOrEmpty)
	}

	jsonCollection, errW := rw.GetJsonResultsCollection(key)

	if errW.IsEmpty() {
		jsonCollection.
			UnmarshalIntoSameIndex(unmarshallingItems...)
	}

	return jsonCollection, errW
}
