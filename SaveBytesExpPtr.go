package rediswrapper

import (
	"time"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func (rw *Wrapper) SaveBytesExpPtr(
	key string,
	byteData *[]byte,
	time time.Duration,
) *errorwrapper.Wrapper {
	statusCmd := rw.client.Set(
		rw.ctx,
		key,
		*byteData,
		time)

	return rw.statusCmdErrWrapper(
		key,
		errtype.RedisUpdateFailed,
		statusCmd)
}
