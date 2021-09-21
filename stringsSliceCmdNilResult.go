package rediswrapper

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"

	"github.com/evatix-go/rediswrapper/internal/messages"
)

func (rw *Wrapper) stringsSliceCmdNilResult(
	key string,
	additionalErrType errtype.Variation,
) *errstr.Results {
	return errstr.EmptyResultsWithError(
		errnew.MessagesPtr(
			errtype.EmptyPointerOrNullPointer,
			additionalErrType.String(),
			messages.StringCmdNull,
			keyValidationMsg(key)))
}
