package rediswrapper

import (
	"encoding/json"

	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

func (rw *Wrapper) GetStrings(
	key string,
) *errstr.Results {
	statusCmd := rw.client.Get(
		rw.ctx,
		key)

	if statusCmd == nil {
		return &errstr.Results{
			Values:       nil,
			ErrorWrapper: rw.statusCmdNullError(),
		}
	}

	allBytes, err := statusCmd.Bytes()
	if err != nil {
		return &errstr.Results{
			Values:       nil,
			ErrorWrapper: errnew.ErrPtr(err),
		}
	}

	var stringsResults []string
	err2 := json.Unmarshal(allBytes, &stringsResults)

	return &errstr.Results{
		Values:       &stringsResults,
		ErrorWrapper: errnew.ErrPtr(err2),
	}
}
