package rediswrapper

import (
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper"
)

func (rw *Wrapper) SaveStrCollectionPtr(
	key string,
	hashmap *corestr.CollectionPtr,
) *errorwrapper.Wrapper {
	jsonResult := hashmap.Json()

	return rw.
		SaveJsonResult(key, jsonResult)
}
