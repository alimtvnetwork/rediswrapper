package rediswrapper

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/errorwrapper"
)

func (rw *Wrapper) SaveJsonResultsCollection(
	key string,
	jsonResultsCollection *corejson.ResultsCollection,
) *errorwrapper.Wrapper {
	return rw.SaveJsonResultsCollectionExp(
		key,
		jsonResultsCollection,
		constants.Zero)
}
