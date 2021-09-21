package rediswrapper

import (
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

func (rw *Wrapper) GetStrHashset(
	key string,
) *errstr.Hashset {
	jsonResult := rw.GetAsErrorJsonResult(key)
	hashset := corestr.EmptyHashset()

	if jsonResult.HasError() {
		return &errstr.Hashset{
			Hashset:      hashset,
			ErrorWrapper: errnew.ErrPtr(jsonResult.Error),
		}
	}

	_, err := hashset.ParseInjectUsingJson(
		jsonResult.Result)

	return &errstr.Hashset{
		Hashset:      hashset,
		ErrorWrapper: errnew.ErrPtr(err),
	}
}
