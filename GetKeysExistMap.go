package rediswrapper

func (rw *Wrapper) GetKeysExistMap(keys ...string) (
	individualsMap *map[string]bool,
	allExists bool,
) {
	if keys == nil {
		return rw.GetKeysExistMapPtr(nil)
	}

	return rw.GetKeysExistMapPtr(&keys)
}
