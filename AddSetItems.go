package rediswrapper

import (
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
)

// Reference : http://t.ly/31MR
func (rw *Wrapper) AddSetItems(key string, uniqueItems ...interface{}) *errwrappers.Collection {
	errsCollection := errwrappers.Empty()

	if uniqueItems == nil {
		return errsCollection.Add(errtype.Empty)
	}

	for _, item := range uniqueItems {
		intCmd := rw.client.SAdd(
			rw.ctx, key, item)

		if intCmd == nil {
			errsCollection.AddUsingMessages(
				errtype.NullOrEmptyReference,
				"intCmd is nil")

			continue
		}

		err := intCmd.Err()

		if err != nil {
			errsCollection.AddTypeError(
				errtype.RequestFailed,
				err)
		}
	}

	return errsCollection
}
