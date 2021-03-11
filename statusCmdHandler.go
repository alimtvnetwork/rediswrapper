package rediswrapper

import (
	"github.com/go-redis/redis/v8"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"

	"github.com/evatix-go/rediswrapper/internal/messages"
)

func (rw *Wrapper) statusCmdErrWrapper(statusCmd *redis.StatusCmd) *errorwrapper.Wrapper {
	if statusCmd == nil {
		return rw.statusCmdNullError()
	}

	return errnew.ErrPtr(statusCmd.Err())
}

func (rw *Wrapper) statusCmdNullError() *errorwrapper.Wrapper {
	return errorwrapper.NewUsingMessagePtr(
		errtype.NullOrEmptyReference,
		messages.StringCmdNull)
}
