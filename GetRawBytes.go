package rediswrapper

import "gitlab.com/evatix-go/core/defaulterr"

// GetAsUnmarshal will bind the json data with the jsonObj.
// If Get encounters any errors, it will return the error message.
func (rw *Wrapper) GetRawBytes(key string) ([]byte, error) {
	stringCmd := rw.client.Get(rw.ctx, key)

	if stringCmd == nil {
		return []byte{}, defaulterr.CannotProcessNilOrEmpty
	}

	return stringCmd.Bytes()
}
