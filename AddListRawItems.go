package rediswrapper

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

// AddListRawItems
//
// Usages RPUSH to push items.
// (Redis unordered list, not unique, just list)
//
// Reference : https://redis.io/commands/LSET | https://redis.io/commands/RPUSH
//
// Retrieve items using
// ListUsingLimit or List or ListAsCollection
func (rw *Wrapper) AddListRawItems(
	key string,
	values ...[]byte,
) *errorwrapper.Wrapper {
	if values == nil {
		return nil
	}

	// https://redis.io/commands/LSET
	for _, value := range values {
		intCmd := rw.client.RPush(rw.ctx, key, value)

		if intCmd.Err() != nil {
			return errnew.MessagesPtr(
				errtype.RedisCreateFailed,
				string(value),
				keyValidationMsg(key),
			)
		}
	}

	return nil
}
