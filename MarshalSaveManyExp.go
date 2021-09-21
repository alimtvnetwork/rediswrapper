package rediswrapper

import (
	"time"

	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"

	"github.com/evatix-go/rediswrapper/internal/messages"
)

func (rw *Wrapper) MarshalSaveManyExp(
	key string,
	exp time.Duration,
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

	return rw.SaveJsonResultsCollectionExp(
		key,
		jsonResultCollection,
		exp)
}
