package rediswrapper

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"

	"github.com/evatix-go/rediswrapper/internal/messages"
)

// GetUnmarshalMany unmarshal serialized binary into multiple objects
func (rw *Wrapper) GetUnmarshalMany(
	key string,
	anyItems ...interface{},
) *errorwrapper.Wrapper {
	if anyItems == nil {
		return errnew.MessagesPtr(
			errtype.NullOrEmptyReference,
			messages.CannotUnmarshallNilItems)
	}

	jsonCollection, errW := rw.GetJsonResultsCollection(key)

	if errW.IsEmpty() {
		jsonCollection.
			UnmarshalIntoSameIndex(anyItems...)
	}

	return errW
}
