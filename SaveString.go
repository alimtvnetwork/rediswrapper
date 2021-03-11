package rediswrapper

import "gitlab.com/evatix-go/errorwrapper"

func (rw *Wrapper) SaveString(
	key string,
	stringData string,
) *errorwrapper.Wrapper {
	statusCmd := rw.client.Set(
		rw.ctx,
		key,
		stringData,
		0)

	return rw.statusCmdErrWrapper(statusCmd)
}
