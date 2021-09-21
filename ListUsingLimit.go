package rediswrapper

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

// ListUsingLimit (redis actual list retrieval)
//
// Added items to key using (AddListRawItems or AddListStringItems or RPUSH)
// will be retrieved using this method.
// Usages LRange inside to get list items and can limit the data.
//
// Reference :
//  - https://redis.io/commands/LSET
//  - https://redis.io/commands/RPUSH
//  - https://redis.io/commands/LRANGE
//
// limit -1 represents all
func (rw *Wrapper) ListUsingLimit(key string, limit int64) *errstr.Results {
	stringsSliceCmd := rw.client.LRange(
		rw.ctx,
		key,
		constants.Zero,
		limit)

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
