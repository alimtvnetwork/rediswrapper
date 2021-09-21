package rediswrapper

import (
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"

	"github.com/evatix-go/rediswrapper/internal/messages"
)

// AddUniqueStringsToKey (using redis SAdd - Redis Hashset)
//
// Usages SAdd to add (SAdd adds unique values, like of hashset in redis)
// whereas RPUSH (AddListStringItems) add any item to the list.
//
// Reference : http://t.ly/31MR
//
// Get items UniqueListItems(SMembers) or
// UniqueListItemsAsHashset(SMembers) or
// UniqueListAsCollection(SMembers)
func (rw *Wrapper) AddUniqueStringsToKey(
	key string,
	uniqueItems ...string,
) *errwrappers.Collection {
	errsCollection := errwrappers.Empty()

	if uniqueItems == nil {
		return errsCollection.Add(
			errtype.NullOrEmptyReference)
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
