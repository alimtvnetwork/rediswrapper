package rediswrapper

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

// UniqueListSlice
//
// Retrieves SetItems key list,
// it uses SMember for redis to get the list items.
//
// It is actual redis list.
func (rw *Wrapper) UniqueListSlice(key string) *errstr.Results {
	stringsSliceCmd := rw.client.SMembers(rw.ctx, key)

	if stringsSliceCmd == nil {
		return rw.stringsSliceCmdNilResult(
			key,
			errtype.ReadFailed)
	}

	values, err := stringsSliceCmd.Result()

	if err != nil {
		return errstr.EmptyResultsWithError(
			errnew.MessagesPtr(
				errtype.RedisCrudFailed,
				errtype.CommunicationFailed.String(),
				err.Error(),
				keyValidationMsg(key)))
	}

	return errstr.EmptyErrorResults(values...)
}
