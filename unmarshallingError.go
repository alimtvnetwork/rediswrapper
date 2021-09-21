package rediswrapper

import (
	"gitlab.com/evatix-go/core/coredata/coredynamic"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func unmarshallingError(
	key string,
	err error,
	unmarshallingObject interface{},
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errnew.MessagesPtr(
		errtype.Unmarshalling,
		err.Error(),
		coredynamic.TypeName(unmarshallingObject),
		errorwrapper.SimpleReferencesCompileOptimized(errtype.KeyValidationFailed, key))
}
