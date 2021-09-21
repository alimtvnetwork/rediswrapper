package rediswrapper

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

// GetSearchedKeys
//
// Ref : https://redis.io/commands/KEYS
func (rw *Wrapper) GetSearchedKeys(
	keySearchingPattern string,
) *errstr.Results {
	stringsSliceCmd := rw.
		client.
		Keys(
			rw.ctx,
			keySearchingPattern)

	if stringsSliceCmd == nil {
		return rw.stringsSliceCmdNilResult(
			keySearchingPattern,
			errtype.SearchFailed)
	}

	stringItems, err := stringsSliceCmd.Result()

	if err != nil {
		return rw.emptyErrorStringResultsFor(
			keySearchingPattern,
			errtype.SearchFailed,
			err)
	}

	return errstr.EmptyErrorResults(stringItems...)
}
