package rediswrapper

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func (rw *Wrapper) emptyErrorStringResultsFor(
	key string,
	additionalErrType errtype.Variation,
	err error,
) *errstr.Results {
	if err == nil {
		return errstr.EmptyResults()
	}

	return errstr.EmptyResultsWithError(
		errnew.MessagesPtr(
			errtype.EmptyResult,
			err.Error(),
			additionalErrType.String(),
			keyValidationMsg(key)))
}
