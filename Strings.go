package rediswrapper

import (
	"encoding/json"

	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

// Strings
//
// Retrieves strings array / slice from redis,
// (it is not actual redis list but marshal / unmarshalled json)
//
// To use actual redis list (use AddListRawItems or AddListStringItems)
func (rw *Wrapper) Strings(
	key string,
) *errstr.Results {
	statusCmd := rw.client.Get(
		rw.ctx,
		key)

	if statusCmd == nil {
		return rw.emptyErrorStringResultsForNil(key)
	}

	allBytes, err := statusCmd.Bytes()

	if err != nil {
		return rw.emptyErrorStringResultsFor(
			key,
			errtype.ReadFailed,
			err)
	}

	var stringsResults []string
	err2 := json.Unmarshal(
		allBytes,
		&stringsResults)

	if err2 != nil {
		errWp := unmarshallingError(key, err2, stringsResults)

		return errstr.EmptyResultsWithError(errWp)
	}

	return errstr.EmptyErrorResults(
		stringsResults...)
}
