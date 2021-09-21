package rediswrapper

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errjson"
)

func (rw *Wrapper) GetAsErrorJsonResult(key string) *errjson.Result {
	allBytes, err := rw.GetRawBytes(key)

	return errjson.New(allBytes, err)
}
