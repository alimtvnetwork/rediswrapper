package rediswrapper

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

// UniqueListItems
//
// Retrieve items which set by SetItems (actual array of Redis)
func (rw *Wrapper) UniqueListItems(key string) *errstr.Results {
	stringsSliceCmd := rw.client.SMembers(
		rw.ctx,
		key)

	if stringsSliceCmd == nil {
		return rw.stringsSliceCmdNilResult(key, errtype.ReadFailed)
	}

	stringItems, err := stringsSliceCmd.Result()

	if err != nil {
		return rw.emptyErrorStringResultsFor(
			key,
			errtype.ReadFailed,
			err)
	}

	return errstr.EmptyErrorResults(stringItems...)
}
