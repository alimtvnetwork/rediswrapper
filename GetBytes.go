package rediswrapper

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errbyte"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

func (rw *Wrapper) GetBytes(key string) *errbyte.Results {
	statusCmd := rw.client.Get(rw.ctx, key)

	allBytes, err := statusCmd.Bytes()

	if err != nil || allBytes == nil {
		return &errbyte.Results{
			Values:       &[]byte{},
			ErrorWrapper: errnew.ErrPtr(err),
		}
	}

	return &errbyte.Results{
		Values:       &allBytes,
		ErrorWrapper: errnew.EmptyPtr,
	}
}
