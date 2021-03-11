package rediswrapper

import (
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

func (rw *Wrapper) SaveCharStrCollection(
	key string,
	charCollectionMap *corestr.CharCollectionMap,
) *errorwrapper.Wrapper {
	jsonResult := charCollectionMap.Json()

	if jsonResult.HasError() {
		return errnew.ErrPtr(jsonResult.Error)
	}

	statusCmd := rw.client.Set(
		rw.ctx,
		key,
		*jsonResult.Bytes,
		0)

	return rw.statusCmdErrWrapper(statusCmd)
}
