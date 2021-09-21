package rediswrapper

import (
	"time"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

// SaveBytesExp will store bytes in the redis server with expiration.
//
// exp 0 means never.
func (rw *Wrapper) SaveBytesExp(
	key string,
	byteData []byte,
	exp time.Duration,
) *errorwrapper.Wrapper {
	statusCmd := rw.client.Set(
		rw.ctx,
		key,
		byteData,
		exp)

	return rw.statusCmdErrWrapper(
		key,
		errtype.RedisUpdateFailed,
		statusCmd)
}
