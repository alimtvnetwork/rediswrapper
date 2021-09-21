package rediswrapper

import (
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
)

// ListAsCollection (redis actual list retrieval)
//
// Added items to key using (AddListRawItems or AddListStringItems or RPUSH)
// will be retrieved using this method.
// Usages LRange inside to get list items and can limit the data.
// (to use limit use ListUsingLimit)
//
// Here, we will retrieve all the items from the redis list.
//
// Reference :
//  - https://redis.io/commands/LSET
//  - https://redis.io/commands/RPUSH
//  - https://redis.io/commands/LRANGE
func (rw *Wrapper) ListAsCollection(key string) *errstr.Collection {
	errStrResult := rw.List(key)

	return &errstr.Collection{
		Collection: corestr.
			NewCollectionUsingStrings(
				errStrResult.ValueMust(),
				false),
		ErrorWrapper: errStrResult.ErrorWrapper,
	}
}
