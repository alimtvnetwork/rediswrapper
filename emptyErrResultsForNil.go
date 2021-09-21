package rediswrapper

import (
	"github.com/evatix-go/rediswrapper/internal/messages"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func (rw *Wrapper) emptyErrorStringResultsForNil(key string) *errstr.Results {
	return &errstr.Results{
		Values: []string{},
		ErrorWrapper: errnew.MessagesPtr(
			errtype.EmptyPointerOrNullPointer,
			messages.StringCmdNull,
			errorwrapper.SimpleReferencesCompileOptimized(errtype.KeyNotFound, key)),
	}
}
