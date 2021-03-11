package rediswrapper

import (
	"encoding/json"

	"gitlab.com/evatix-go/errorwrapper/errdata/errbool"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

func (rw *Wrapper) GetBools(
	key string,
) *errbool.Results {
	statusCmd := rw.client.Get(
		rw.ctx,
		key)

	if statusCmd == nil {
		return &errbool.Results{
			Values:       nil,
			ErrorWrapper: rw.statusCmdNullError(),
		}
	}

	allBytes, err := statusCmd.Bytes()
	if err != nil {
		return &errbool.Results{
			Values:       nil,
			ErrorWrapper: errnew.ErrPtr(err),
		}
	}

	var stringsResults []bool
	err2 := json.Unmarshal(allBytes, &stringsResults)

	return &errbool.Results{
		Values:       &stringsResults,
		ErrorWrapper: errnew.ErrPtr(err2),
	}
}
