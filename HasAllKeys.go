package rediswrapper

func (rw *Wrapper) HasAllKeys(keys ...string) bool {
	if keys == nil {
		return false
	}

	for _, key := range keys {
		if !rw.HasKey(key) {
			return false
		}
	}

	return true
}
