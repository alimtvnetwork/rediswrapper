package rediswrapper

import (
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

func (rw *Wrapper) GetCharStrCollection(
	key string,
) *errstr.CharCollectionMap {
	statusCmd := rw.
		client.
		Get(rw.ctx, key)

	if statusCmd == nil {
		return &errstr.CharCollectionMap{
			CharCollectionMap: nil,
			ErrorWrapper:      rw.statusCmdNullError(),
		}
	}

	allBytes, err := statusCmd.Bytes()
	jsonResult := corejson.NewPtr(allBytes, err)

	if jsonResult.HasError() {
		return &errstr.CharCollectionMap{
			CharCollectionMap: nil,
			ErrorWrapper:      errnew.ErrPtr(err),
		}
	}

	charCollectionMap := corestr.EmptyCharCollectionMap()
	err2 := charCollectionMap.JsonParseSelfInject(jsonResult)

	return &errstr.CharCollectionMap{
		CharCollectionMap: charCollectionMap,
		ErrorWrapper:      errnew.ErrPtr(err2),
	}
}
