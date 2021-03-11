package rediswrapper

import (
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

func (rw *Wrapper) GetStringsCollection(
	key string,
) *errstr.Collection {
	statusCmd := rw.
		client.
		Get(rw.ctx, key)

	if statusCmd == nil {
		return &errstr.Collection{
			Collection:   nil,
			ErrorWrapper: rw.statusCmdNullError(),
		}
	}

	allBytes, err := statusCmd.Bytes()
	jsonResult := corejson.NewPtr(allBytes, err)

	if jsonResult.HasError() {
		return &errstr.Collection{
			Collection:   nil,
			ErrorWrapper: errnew.ErrPtr(err),
		}
	}

	collection := corestr.EmptyCollection()
	_, err2 := collection.ParseInjectUsingJson(jsonResult)

	return &errstr.Collection{
		Collection:   collection,
		ErrorWrapper: errnew.ErrPtr(err2),
	}
}
