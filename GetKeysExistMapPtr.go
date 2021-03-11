package rediswrapper

import "gitlab.com/evatix-go/core/coredata/corestr"

func (rw *Wrapper) GetKeysExistMapPtr(keys *[]string) (
	individualsMap *map[string]bool,
	allExists bool,
) {
	length := corestr.LengthOfStrings(keys)
	itemsMap := make(map[string]bool, length)
	allExists = false

	if length == 0 {
		return &itemsMap, allExists
	}

	allExists = true
	for _, key := range *keys {
		has := rw.HasKey(key)
		itemsMap[key] = has

		if !has {
			allExists = false
		}
	}

	return &itemsMap, allExists
}
