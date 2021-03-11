package rediswrapper

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"

	"github.com/evatix-go/rediswrapper/internal/messages"
)

func (rw *Wrapper) SaveBytesPtr(
	key string,
	byteData *[]byte,
) *errorwrapper.Wrapper {
	statusCmd := rw.client.Set(
		rw.ctx,
		key,
		*byteData,
		0)

	if statusCmd == nil {
		return errorwrapper.NewUsingMessagePtr(
			errtype.NullOrEmptyReference,
			messages.StringCmdNull)
	}

	return errnew.ErrPtr(statusCmd.Err())
}
