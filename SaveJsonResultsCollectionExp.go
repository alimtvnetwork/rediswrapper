package rediswrapper

import (
	"encoding/json"
	"time"

	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"

	"github.com/evatix-go/rediswrapper/internal/messages"
)

func (rw *Wrapper) SaveJsonResultsCollectionExp(
	key string,
	jsonResultsCollection *corejson.ResultsCollection,
	exp time.Duration,
) *errorwrapper.Wrapper {
	if jsonResultsCollection == nil {
		return errnew.MessagesPtr(
			errtype.NullOrEmptyReference,
			messages.JsonResultNullPointer,
			keyValidationMsg(key))
	}

	marshalBytes, err := json.Marshal(jsonResultsCollection)

	if err != nil {
		return marshallingError(key, err, jsonResultsCollection)
	}

	return rw.SaveBytesExp(key, marshalBytes, exp)
}
