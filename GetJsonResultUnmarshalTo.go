package rediswrapper

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errdata/errjson"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"

	"github.com/evatix-go/rediswrapper/internal/messages"
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
		allBytes, err)

	if jsonResultWithError.HasError() {
		return jsonResultWithError.ErrorWrapper
	}

	if jsonResultWithError.IsEmptyJsonBytes() {
		return errnew.MessagesPtr(
			errtype.Unmarshalling,
			messages.CannotUnmarshallEmpty)
	}

	err2 := jsonResultWithError.
		Unmarshal(unmarshallingAny)

	return errnew.ErrPtr(err2)
}
