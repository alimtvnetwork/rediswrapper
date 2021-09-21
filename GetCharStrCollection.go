package rediswrapper

import (
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"

	"github.com/evatix-go/rediswrapper/internal/rediserrwrapper"
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
			ErrorWrapper:      rediserrwrapper.NullEmptyStringCmd,
		}
	}

	allBytes, err := statusCmd.Bytes()
	jsonResult := corejson.NewPtr(allBytes, err)

	if jsonResult.HasError() {
		return &errstr.CharCollectionMap{
			CharCollectionMap: nil,
			ErrorWrapper:      errnew.ErrPtr(jsonResult.MeaningfulError()),
		}
	}

	charCollectionMap := corestr.EmptyCharCollectionMap()
	err2 := charCollectionMap.JsonParseSelfInject(jsonResult)
	finalErr := errnew.EmptyPtr

	if err2 != nil {
		finalErr = errnew.MessagesPtr(
			errtype.Unmarshalling,
			err2.Error(),
			errorwrapper.SimpleReferencesCompileOptimized(errtype.KeyValidationFailed, key))
	}

	return &errstr.CharCollectionMap{
		CharCollectionMap: charCollectionMap,
		ErrorWrapper:      finalErr,
	}
}
