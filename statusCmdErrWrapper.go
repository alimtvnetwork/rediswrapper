package rediswrapper

import (
	"github.com/go-redis/redis/v8"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func (rw *Wrapper) statusCmdErrWrapper(
	key string,
	additionalErrorType errtype.Variation,
	statusCmd *redis.StatusCmd,
) *errorwrapper.Wrapper {
	if statusCmd == nil {
		return errnew.MessagesPtr(
			errtype.EmptyPointerOrNullPointer,
			"Redis status-cmd null.",
			additionalErrorType.String(),
			keyValidationMsg(key))
	}

	err := statusCmd.Err()
	if err != nil {
		return errnew.MessagesPtr(
			errtype.RedisCrudFailed,
			err.Error(),
			additionalErrorType.String(),
			keyValidationMsg(key))
	}

	return nil
}
