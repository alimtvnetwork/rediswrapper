package rediswrapper

func (rw *Wrapper) HasKey(key string) bool {
	result := rw.Exists(key)

	return result.Value
}
