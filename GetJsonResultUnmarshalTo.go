package rediswrapper

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errdata/errjson"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func (rw *Wrapper) GetJsonResultUnmarshalTo(
	key string,
	unmarshallingAny interface{},
) *errorwrapper.Wrapper {
	allBytes, err := rw.GetRawBytes(key)

	if err != nil {
		return errnew.NewPtr(errtype.ReadFailed, err)
	}

	jsonResultWithError := errjson.New(
		allBytes,
		err)

	err2 := jsonResultWithError.
		Unmarshal(unmarshallingAny)

	if err2.HasError() {
		return unmarshallingError(
			key,
			err2.Error(),
			unmarshallingAny)
	}

	return errnew.EmptyPtr
}
