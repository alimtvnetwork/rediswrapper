package rediswrapper

import (
	"encoding/json"

	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"

	"github.com/evatix-go/rediswrapper/internal/messages"
)

func (rw *Wrapper) SaveJsonResultsCollection(
	key string,
	jsonResultsCollection *corejson.ResultsCollection,
) *errorwrapper.Wrapper {
	if jsonResultsCollection == nil {
		return errorwrapper.NewUsingMessagePtr(
			errtype.NullOrEmptyReference,
			messages.JsonResultNullPointer)
	}

	marshalBytes, err := json.
		Marshal(jsonResultsCollection)

	if err != nil {
		return errnew.ErrPtr(err)
	}

	return rw.SaveBytes(key, marshalBytes)
}
