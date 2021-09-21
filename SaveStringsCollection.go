package rediswrapper

import (
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper"
)

func (rw *Wrapper) SaveStringsCollection(
	key string,
	stringsCollection *corestr.Collection,
) *errorwrapper.Wrapper {
	jsonResult := stringsCollection.Json()

	return rw.SaveJsonResult(
		key,
		jsonResult)
}
