package rediswrapper

import (
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
)

// Reference : http://t.ly/31MR
func (rw *Wrapper) AddSetStringItems(key string, uniqueItems ...string) *errwrappers.Collection {
	errsCollection := errwrappers.Empty()

	if uniqueItems == nil {
		return errsCollection.Add(errtype.NullOrEmptyReference)
	}

	for _, item := range uniqueItems {
		intCmd := rw.client.SAdd(rw.ctx, key, item)

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
