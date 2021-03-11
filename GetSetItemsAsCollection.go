package rediswrapper

import (
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
)

func (rw *Wrapper) GetSetItemsAsCollection(key string) *errstr.Collection {
	errStrResult := rw.GetSetItemsAsSlice(key)

	return &errstr.Collection{
		Collection: corestr.NewCollectionUsingStrings(
			errStrResult.Values,
			false),
		ErrorWrapper: errStrResult.ErrorWrapper,
	}
}
