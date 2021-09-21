package rediswrapper

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
)

// GetAllKeys Ref : https://redis.io/commands/KEYS
func (rw *Wrapper) GetAllKeys() *errstr.Results {
	return rw.GetSearchedKeys(constants.AstrekSymbol)
}
