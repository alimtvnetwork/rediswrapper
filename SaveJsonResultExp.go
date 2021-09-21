package rediswrapper

import (
	"time"

	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"

	"github.com/evatix-go/rediswrapper/internal/messages"
)

func (rw *Wrapper) SaveJsonResultExp(
	key string,
	jsonResult *corejson.Result,
	time time.Duration,
) *errorwrapper.Wrapper {
	if jsonResult == nil {
		return errorwrapper.NewUsingMessagePtr(
			errtype.NullOrEmptyReference,
			messages.JsonResultNullPointer)
	}

	if jsonResult.HasError() {
		return errnew.ErrPtr(jsonResult.Error)
	}

	return rw.SaveBytesExp(
		key,
		jsonResult.Bytes,
		time)
}
