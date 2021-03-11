package rediswrapper

import (
	"encoding/json"

	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/errorwrapper"
)

func (rw *Wrapper) SaveStrings(
	key string,
	stringsData *[]string,
) *errorwrapper.Wrapper {
	jsonBytes, err := json.Marshal(*stringsData)
	jsonResult := corejson.NewPtr(jsonBytes, err)
	statusCmd := rw.client.Set(
		rw.ctx,
		key,
		*jsonResult.Bytes,
		0)

	return rw.statusCmdErrWrapper(statusCmd)
}
