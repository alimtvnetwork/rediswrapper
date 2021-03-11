package rediswrapper

import (
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/errorwrapper"
)

func (rw *Wrapper) SaveJsoner(
	key string,
	jsoner corejson.Jsoner,
) *errorwrapper.Wrapper {
	return rw.SaveJsonResult(key, jsoner.Json())
}
