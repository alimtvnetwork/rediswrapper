package rediswrapper

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errjson"
)

func (rw *Wrapper) GetJsonResult(key string) *errjson.Result {
	allBytes, err := rw.GetRawBytes(key)

	if err != nil {
		return errjson.EmptyWithError(err)
	}

	return errjson.New(allBytes, err)
}
