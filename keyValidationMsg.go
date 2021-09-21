package rediswrapper

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func keyValidationMsg(key string) string {
	return errorwrapper.SimpleReferencesCompileOptimized(errtype.KeyValidationFailed, key)
}
