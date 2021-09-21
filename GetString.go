package rediswrapper

import (
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
		return errstr.ErrorWrapper(
			errnew.MessagesPtr(
				errtype.RedisCrudFailed,
				errtype.ReadFailed.String(),
				err.Error(),
				keyValidationMsg(key)))
	}

	return errstr.EmptyErrorResult(currentStr)
}
