package rediswrapper

import (
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper"
)

func (rw *Wrapper) SaveStrHashset(
	key string,
	hashset *corestr.Hashset,
) *errorwrapper.Wrapper {
	jsonResult := hashset.Json()

	return rw.SaveJsonResult(key, jsonResult)
}
