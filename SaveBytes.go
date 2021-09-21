package rediswrapper

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

// SaveBytes will store bytes in the redis server.
// If Get encounters any errors, it will return the error message.
func (rw *Wrapper) SaveBytes(
	key string,
	byteData []byte,
) *errorwrapper.Wrapper {
	statusCmd := rw.client.Set(
		rw.ctx,
		key,
		byteData,
		0)

	return rw.statusCmdErrWrapper(
		key,
		errtype.RedisUpdateFailed,
		statusCmd)
}
