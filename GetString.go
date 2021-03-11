package rediswrapper

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func (rw *Wrapper) GetString(
	key string,
) *errstr.Result {
	strCmd := rw.client.Get(rw.ctx, key)
	currentStr, err := strCmd.Result()

	if err != nil {
		return &errstr.Result{
			Value: constants.EmptyString,
			ErrorWrapper: errnew.
				NewPtr(errtype.ReadFailed, err),
		}
	}

	return &errstr.Result{
		Value: currentStr,
		ErrorWrapper: errnew.
			EmptyPtr,
	}
}
