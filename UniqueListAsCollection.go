package rediswrapper

import (
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
)

// UniqueListAsCollection
//
// Retrieve items set by SetItems method.
func (rw *Wrapper) UniqueListAsCollection(key string) *errstr.Collection {
	errStrResult := rw.UniqueListSlice(key)

	return &errstr.Collection{
		Collection: corestr.NewCollectionUsingStrings(
			errStrResult.ValueMust(),
			false),
		ErrorWrapper: errStrResult.ErrorWrapper,
	}
}
