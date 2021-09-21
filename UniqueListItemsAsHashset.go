package rediswrapper

import (
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
)

func (rw *Wrapper) UniqueListItemsAsHashset(key string) *errstr.Hashset {
	errStrResult := rw.UniqueListSlice(key)

	return &errstr.Hashset{
		Hashset: corestr.NewHashsetUsingStrings(
			errStrResult.ValueMust()),
		ErrorWrapper: errStrResult.ErrorWrapper,
	}
}
