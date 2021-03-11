package rediswrapper

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errbool"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
)

func (rw *Wrapper) ExistMany(keys ...string) *errbool.ResultsWithErrorCollection {
	length := len(keys)
	if length == 0 {
		return errbool.EmptyResultsWithErrorCollectionPtr()
	}

	list := make([]bool, 0, length)
	errsCollection := errwrappers.Empty()
	for _, key := range keys {
		errBool := rw.Delete(key)
		list = append(
			list,
			errBool.Value)

		errsCollection.
			AddWrapperPtr(errBool.ErrorWrapper)
	}

	return &errbool.ResultsWithErrorCollection{
		Values:        &list,
		ErrorWrappers: errsCollection,
	}
}
