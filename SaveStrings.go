package rediswrapper

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func (rw *Wrapper) SaveStrings(
	key string,
	stringsData []string,
) *errorwrapper.Wrapper {
	jsonResult := corejson.NewFromAny(stringsData)
	statusCmd := rw.client.Set(
		rw.ctx,
		key,
		jsonResult.Bytes,
		constants.Zero)

	return rw.statusCmdErrWrapper(
		key,
		errtype.RedisUpdateFailed,
		statusCmd)
}
