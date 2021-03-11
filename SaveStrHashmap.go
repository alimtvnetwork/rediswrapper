package rediswrapper

import (
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper"
)

func (rw *Wrapper) SaveStrHashmap(
	key string,
	hashmap *corestr.Hashmap,
) *errorwrapper.Wrapper {
	jsonResult := hashmap.Json()

	return rw.SaveJsonResult(key, jsonResult)
}
