package rediswrapper

import (
	"gitlab.com/evatix-go/core"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func (rw *Wrapper) GetSetItemsAsSlice(key string) *errstr.Results {
	stringsSliceCmd := rw.client.SMembers(rw.ctx, key)

	if stringsSliceCmd == nil {
		return &errstr.Results{
			Values: core.EmptyStringsPtr(),
			ErrorWrapper: errnew.MessagesPtr(
				errtype.NullOrEmptyReference,
				"stringsSliceCmd is nil, cannot communicate."),
		}
	}

	values, err := stringsSliceCmd.Result()

	if err != nil || values == nil {
		return &errstr.Results{
			Values: core.EmptyStringsPtr(),
			ErrorWrapper: errnew.NewPtr(
				errtype.CommunicationFailed,
				err),
		}
	}

	return &errstr.Results{
		Values:       &values,
		ErrorWrapper: errnew.EmptyPtr,
	}
}
