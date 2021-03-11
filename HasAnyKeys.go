package rediswrapper

func (rw *Wrapper) HasAnyKeys(keys ...string) bool {
	if keys == nil {
		return false
	}

	for _, key := range keys {
		if rw.HasKey(key) {
			return true
		}
	}

	return false
}
