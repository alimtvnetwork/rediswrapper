package rediswrapper

import (
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

func (rw *Wrapper) GetStrHashmapPtr(
	key string,
) *errstr.Hashmap {
	jsonResult := rw.GetAsErrorJsonResult(key)
	hashmap := corestr.EmptyHashmap()

	if jsonResult.HasError() {
		return &errstr.Hashmap{
			Hashmap:      hashmap,
			ErrorWrapper: errnew.ErrPtr(jsonResult.Error),
		}
	}

	_, err := hashmap.ParseInjectUsingJson(jsonResult.Result)

	return &errstr.Hashmap{
		Hashmap:      hashmap,
		ErrorWrapper: errnew.ErrPtr(err),
	}
}
