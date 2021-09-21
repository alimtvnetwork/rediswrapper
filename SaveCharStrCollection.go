package rediswrapper

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func (rw *Wrapper) SaveCharStrCollection(
	key string,
	charCollectionMap *corestr.CharCollectionMap,
) *errorwrapper.Wrapper {
	jsonResult := charCollectionMap.Json()

	if jsonResult.HasError() {
		return errnew.MessagesPtr(
			errtype.Unmarshalling,
			jsonResult.Error.Error(),
			keyValidationMsg(key))
	}

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
