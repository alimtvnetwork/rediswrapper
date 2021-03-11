package rediswrapper

import (
	"gitlab.com/evatix-go/core"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"

	"github.com/evatix-go/rediswrapper/internal/messages"
)

func (rw *Wrapper) GetListAsSlicePtr(key string) *errstr.Results {
	stringSliceCmd := rw.client.LRange(rw.ctx, key, 0, -1)

	if stringSliceCmd == nil {
		return &errstr.Results{
			Values: core.EmptyStringsPtr(),
			ErrorWrapper: errnew.MessagesPtr(
				errtype.NullOrEmptyReference,
				messages.StringCmdNull),
		}
	}

	stringItems, err := stringSliceCmd.Result()

	if err != nil || stringItems == nil {
		return &errstr.Results{
			Values: core.EmptyStringsPtr(),
			ErrorWrapper: errnew.NewPtr(
				errtype.ReadFailed,
				err),
		}
	}

	return &errstr.Results{
		Values:       &stringItems,
		ErrorWrapper: errnew.EmptyPtr,
	}
}
