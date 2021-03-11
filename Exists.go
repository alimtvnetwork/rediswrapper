package rediswrapper

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errbool"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"

	"github.com/evatix-go/rediswrapper/internal/messages"
)

func (rw *Wrapper) Exists(key string) *errbool.Result {
	intCmd := rw.client.Exists(rw.ctx, key)

	if intCmd == nil {
		return &errbool.Result{
			Value: false,
			ErrorWrapper: errnew.MessagesPtr(
				errtype.CommunicationFailed,
				messages.IntCmdNull),
		}
	}

	result, err :=
		intCmd.Result()

	if err != nil {
		return &errbool.Result{
			Value: false,
			ErrorWrapper: errnew.NewPtr(
				errtype.ReadFailed,
				err),
		}
	}

	return &errbool.Result{
		Value:        result > 0,
		ErrorWrapper: errnew.EmptyPtr,
	}
}
