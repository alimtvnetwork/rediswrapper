package rediswrapper

import (
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"

	"github.com/evatix-go/rediswrapper/internal/messages"
)

// AddUniqueAnysToKey (using redis SAdd - Redis Hashset)
//
// Usages SAdd to add (adds unique items to the key),
// whereas, RPUSH (AddListStringItems) pushes items
// to the list regardless.
//
// Reference : http://t.ly/31MR
//
// Get items UniqueListItems(SMembers) or
// UniqueListItemsAsHashset(SMembers) or
// UniqueListAsCollection(SMembers)
func (rw *Wrapper) AddUniqueAnysToKey(
	key string,
	uniqueItems ...interface{},
) *errwrappers.Collection {
	errsCollection := errwrappers.Empty()

	if len(uniqueItems) == 0 {
		return errsCollection
	}

	for _, item := range uniqueItems {
		intCmd := rw.client.SAdd(
			rw.ctx,
			key,
			item)

		if intCmd == nil {
			errsCollection.AddUsingMessages(
				errtype.NullOrEmptyReference,
				messages.IntCmdNull)

			continue
		}

		err := intCmd.Err()

		if err != nil {
			errsCollection.AddTypeError(
				errtype.RequestFailed,
				err)
		}
	}

	return errsCollection
}
