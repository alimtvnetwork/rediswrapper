package rediswrapper

import (
	"encoding/json"
	"errors"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errdata/errbool"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func (rw *Wrapper) GetBooleans(
	key string,
) *errbool.Results {
	stringCmd := rw.client.Get(
		rw.ctx,
		key)

	if stringCmd == nil {
		return errbool.ResultsError(
			errtype.EmptyPointerOrNullPointer,
			errors.New("redis stringCmd for key - "+key))
	}

	allBytes, err := stringCmd.Bytes()
	if err != nil {
		return errbool.ResultsError(
			errtype.RedisCrudFailed,
			err)
	}

	var stringsResults []bool
	err2 := json.Unmarshal(
		allBytes,
		&stringsResults)

	if err2 != nil {
		return errbool.EmptyResultsWithError(
			errnew.MessagesPtr(
				errtype.Unmarshalling,
				err2.Error(),
				errorwrapper.SimpleReferencesCompileOptimized(errtype.KeyValidationFailed, key)))
	}

	return errbool.EmptyErrorResults(stringsResults...)
}
