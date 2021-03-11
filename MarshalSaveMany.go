package rediswrapper

import (
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"

	"github.com/evatix-go/rediswrapper/internal/messages"
)

// MarshalSaveMany marshall multiple objects into corejson.ResultCollection
func (rw *Wrapper) MarshalSaveMany(
	key string,
	anyItems ...interface{},
) *errorwrapper.Wrapper {
	if anyItems == nil {
		return errnew.MessagesPtr(
			errtype.NullOrEmptyReference,
			messages.CannotMarshallNilOrEmpty)
	}

	jsonResultCollection := corejson.NewResultsCollection(len(anyItems))
	jsonResultCollection.
		AddsAnysPtr(&anyItems)

	return rw.SaveJsonResultsCollection(
		key,
		jsonResultCollection)
}
